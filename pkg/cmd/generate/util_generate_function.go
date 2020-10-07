package generate

import (
	"bufio"
	"context"
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	github_helpers "github.com/Benbentwo/bens-binary/pkg/github"
	"github.com/Benbentwo/utils/util"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"text/template"
)

// This should be found in the template_base.txt
const BASE_COMMAND_INSERT_LINE = `Section to add commands to:`
const BASE_COMMAND_TEMPLATE = `template_base_v0.go`

// https://gist.github.com/Benbentwo/7f0d31820b4228864ba4dc00fb17767b
const DefaultGistForTemplates = "7f0d31820b4228864ba4dc00fb17767b"

var (
	TemplateFUNctionMap = template.FuncMap{
		"title":   strings.Title,
		"toLower": strings.ToLower,
	}
	logs = util.Logger()
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type SupportedOptions struct {
	GenerateOptions
	common.CommonOptions
}

// GetAddonOptions the command line options
type GenerateFunctionOptions struct {
	GenerateOptions
	isBaseCommand bool

	Folder               string
	Filename             string
	TitledFolderFilename string

	LongDescription     string
	ExampleString       string
	CommandUse          string
	ShortDescription    string
	SupportedOptions    SupportedOptions
	ChosenOption        string
	NoExtensionFilename string
	TemplateFile        string
}

var (
	util_generate_function_long = `
		Attempts to generate a go file to help the development of this application
`

	util_generate_function_example = `
		# Utility to search a file for a string
		bb util generate function util util_generate_function

		# Don't ask questions - run in batch mode
	`
)

func NewCmdGenerateFunction(commonOpts *common.CommonOptions) *cobra.Command {
	options := &GenerateFunctionOptions{
		GenerateOptions: GenerateOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:     "go-code",
		Short:   "Generates a go file for adding a new command",
		Long:    util_generate_function_long,
		Example: util_generate_function_example,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}
	// this command is not intended to be run in batch mode...
	cmd.Flags().StringVarP(&options.Folder, "Folder-name", "d", "", "Folder to create the file in")
	cmd.Flags().StringVarP(&options.Filename, "file-name", "f", "", "File to create in a Folder")

	return cmd
}

// Run implements this command
func (o *GenerateFunctionOptions) Run() error {

	client, _, _, err := github_helpers.Preamble()
	if err != nil {
		return errors.Errorf("getting github client: %s", err)
	}

	var isGist bool
	var thisGist *github.Gist
	localTemplates := "./templates"
	ex, err := util.DirExists(localTemplates)
	availableTemplates := make([]string, 0)
	if err != nil {
		return errors.Errorf("local template dir issue: %s", err)
	} else if ex {
		availableTemplates = util.ListFilesInDirFilter(localTemplates, `(.*\.txt)`)
	} else {
		// TODO get Files from GIST
		util.Logger().Info("Using " + util.ColorInfo("Gist") + " as " + util.ColorInfo("Template Source"))
		gist, resp, err := client.Gists.Get(context.Background(), DefaultGistForTemplates)
		if err != nil {
			return errors.Errorf("getting gist failed: %s", err)
		}
		if resp.StatusCode >= 300 {
			return errors.Errorf("getting gist failed, response code %d: %s", resp.StatusCode, err)
		}

		for filename := range gist.Files {
			if strings.HasSuffix(string(filename), ".go") {
				availableTemplates = append(availableTemplates, *gist.Files[filename].Filename)
				isGist = true
			}
		}
		thisGist = gist
	}

	o.TemplateFile, err = util.Pick("What template would you like to use?", availableTemplates, "")
	check(err)
	// if isGist {
	// 	o.TemplateFile
	// }

	var isBase = o.TemplateFile == BASE_COMMAND_TEMPLATE
	if isBase {
		logs.Info("Base Commands should be one word then the .go extension")
		logs.Info("There should not be more than one base in a folder")
		logs.Info("	this allows us to run `bb <some base> <some other base>` and see a list of commands added to that other base.")
		logs.Info("	which will create command structure similar to the commands.", nil)
	}
	if o.Folder == "" {
		// o.Folder, err = util.PickValue("What Folder would you like to put this in", "",true, "Folder inside of cmd of this project root", o.In, o.Out, o.Err)
		logs.Info("What Folder would you like this in? starts in ./pkg/cmd/<your-answer>", nil)
		logs.Info("You can create new ones, and subdirectories ./pkg/cmd/a/b", nil)
		// o.Folder, err = util.Pick(o.CommonOptions, "What Folder would you like this in (starting from pkg/cmd/...? you can create new directories", util.ListSubDirectories("./pkg/cmd/"), "dev")

		o.Folder, err = util.Pick("What Folder would you like this in?", util.ListSubDirectoriesRecusively("./pkg/cmd/"), "dev")
		if err != nil {
			return err
		}
		logs.Info("Folders: %s", o.Folder)
	}
	var path = o.Folder
	splitPath := strings.Split(o.Folder, "/")
	o.Folder = splitPath[len(splitPath)-1]
	var originalFilename = ""
	if o.Filename == "" {
		if isBase {
			o.Filename, err = util.PickValueFromPath("What would you like to call the file", o.Folder, true, "File name should follow the structure of foldername_filename", o.In, o.Out, o.Err)
		} else {
			o.Filename, err = util.PickValue("What would you like to call the file? (this results in *folder*_*thisFileName*)", "", "File name should follow the structure of foldername_filename", false)
			originalFilename = o.Filename
			o.TitledFolderFilename = strings.Title(o.Folder) + strings.Title(RemoveGoExtension(o.Filename))
			o.Filename = strings.ToLower(o.Folder) + "_" + o.Filename
		}
		check(err)
		matched, _ := regexp.MatchString(`(.*\.go)`, o.Filename)
		if !matched {
			logs.Debugln("Adding .go extension")
			o.NoExtensionFilename = o.Filename
			o.Filename = o.Filename + ".go"
			originalFilename += ".go"
		} else {
			logs.Debugln("Not adding .go extension")

			var extension = filepath.Ext(o.Filename)
			o.NoExtensionFilename = o.NoExtensionFilename[0 : len(o.Filename)-len(extension)]
		}
	}
	var fullFilePath = util.StripTrailingSlash(path) + "/" + o.Filename
	b, err := util.FileExists(fullFilePath)
	if b {
		response, err := util.Confirm("Are you Sure you want to override the file that already exists? This is NOT recommended", false, "that file name already exists, confirming this will override it")
		if err != nil {
			return errors.Errorf("getting response from user")
		}
		if !response { // answered no I don't
			return nil
		}
	}

	{ // section for command stuff - braces help you collapse it in your IDE
		fileNameStripped := RemoveGoExtension(originalFilename)
		if isBase {
			o.CommandUse, err = util.PickValue("Command:", fileNameStripped, "What would you like for the command to use, this should be a single word, or hyphenated", false)
		} else {
			o.CommandUse, err = util.PickValue("Command:", fileNameStripped, "What would you like for the command to use, this should be a single word, or hyphenated", false)
		}
		check(err)

		if !isBase {
			o.ShortDescription, err = util.PickValue("Short Description: ", "", "What would you like for the short description, this should be a single word, or hyphenated", false)
			check(err)
			o.LongDescription, err = util.PickValue("Long Description: ", "", "What would you like for the long description", false)
			check(err)
			o.ExampleString, err = util.PickValue("Example: ", "", " What would you like for the example command", false)
			check(err)
		}
	}

	// create an empty array
	var bases = make([]string, 0)
	bases, err = FindBaseCommands("./")

	// File Gen - put it down here so we don't create the file till they answer all the questions
	// if they ctrl-c we don't want empty files cluttering our project.
	if exists, _ := util.DirExists(path); !exists {
		err := os.MkdirAll(path, 0760)
		if err != nil {
			return errors.Wrap(err, "couldn't make dir for folder")
		}
	}
	//
	var codeTemplate []byte
	if isGist {
		gistFile := github.GistFilename(o.TemplateFile)
		util.Logger().Debugf("File: %s", o.TemplateFile)
		util.Logger().Debugf("GistFile: %s", gistFile)
		codeTemplate = []byte(*thisGist.Files[gistFile].Content)
	} else {
		codeTemplate, err = ioutil.ReadFile("templates/" + o.TemplateFile)
		check(err)
	}

	t := template.Must(template.New("template").Funcs(TemplateFUNctionMap).Parse(string(codeTemplate)))

	if t == nil {
		return errors.Errorf("Unable to parse template %s", t)
	}

	// create the file they want
	f, err := os.Create(fullFilePath)
	check(err)

	err = t.Execute(f, o)
	if err != nil {
		return errors.Wrapf(err, "Error executing template %s", t.Name())
	}

	logs.Debug("BASES: %s", bases)
	var pickedBase = ""
	pickedBase, err = util.Pick("What base file would you like to use?", bases, "")
	check(err)

	// Time to determine what type of line we're adding, as its different in cmd.go (the main cmd)
	//   which is found by running just `bb` - it shows groups
	//   vs anything else in which case we just add the New Cmd string.
	// Must match Template of generated function
	if isBase {
		if pickedBase == "pkg/cmd/cmd.go" {

		} else {
			// NewCmdDev(commonOpts *common.CommonOptions)
			err = addNewCmdToBaseFile(pickedBase, "cmd.AddCommand(NewCmd"+strings.Title(o.CommandUse)+"(commonOpts))\n")
		}
	} else { // adding a generationCommand to an existing base.
		err = addNewCmdToBaseFile(pickedBase, "\tcmd.AddCommand(NewCmd"+o.TitledFolderFilename+"(commonOpts))\n")
	}
	check(err)

	return nil
}

func FindBaseCommands(path string) ([]string, error) {
	libRegEx, e := regexp.Compile("^[^_]*go$")
	if e != nil {
		return nil, errors.Errorf("Error: %s", e, e)
	}
	var splice = make([]string, 0) //create an empty array

	e = filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && libRegEx.MatchString(info.Name()) && !info.IsDir() {
			splice = append(splice, path)
		}
		return nil
	})
	if e != nil {
		return nil, errors.Errorf("Error: %s", e, e)
	}
	logs.Info("%", splice)
	return splice, nil
}

func FindLineToInsertCommandTo(path string, search string) (int, error) {
	arr, err := util.FindMatchesInFile(search, path)
	check(err)
	if len(arr) == 0 {
		return -1, errors.Errorf("No String `%s` found in file `%s`, Error: %s", err, search, path, err)
	}
	if len(arr) > 1 {
		logs.Warn("Found multiple lines to insert to on lines %s", arr)
		// TODO add support for multiple finds incase there are multiple declarations in one file that we want to support
		//   e.g. Theres NewCmdUtil and NewCmdUtility or something.
	}
	return arr[0], nil
}

// Common practice should be make an exportable (Titled func) generic, then make a local named the same call with your default value.
func findLineToInsertCommandTo(path string) (int, error) {
	val, err := FindLineToInsertCommandTo(path, BASE_COMMAND_INSERT_LINE)
	if err != nil {
		util.Logger().Warn("Couldn't find line to insert add command onto.")
		util.Logger().Warn("This command will not be immediately available and code changes are required.")
		// 	TODO add info line here for the exact command to add somewhere to get it to work.
	}
	check(err)
	return val, nil
}

func AddNewCmdToBaseFile(path string, insertString string, lineNumber int) error {
	err := InsertStringToFile(path, insertString, lineNumber)
	check(err)
	return nil
}
func addNewCmdToBaseFile(path string, functionName string) error {
	line, err := findLineToInsertCommandTo(path)
	check(err)
	err = AddNewCmdToBaseFile(path, functionName, line)
	check(err)
	return nil
}

func File2lines(filePath string) ([]string, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return LinesFromReader(f)
}

func LinesFromReader(r io.Reader) ([]string, error) {
	var lines []string
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return lines, nil
}

/**
 * Insert sting to n-th line of file.
 * If you want to insert a line, append newline '\n' to the end of the string.
 */
func InsertStringToFile(path, str string, index int) error {
	lines, err := File2lines(path)
	if err != nil {
		return err
	}

	fileContent := ""
	for i, line := range lines {
		if i == index {
			fileContent += str
		}
		fileContent += line
		fileContent += "\n"
	}

	return ioutil.WriteFile(path, []byte(fileContent), 0644)
}

func RemoveGoExtension(fileName string) string {
	matched, _ := regexp.MatchString(`(.*\.go)`, fileName)
	if matched {
		return fileName[0 : len(fileName)-len(".go")]
	}
	return fileName
}
