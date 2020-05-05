package init

import (
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/Benbentwo/bens-binary/pkg/cmd/github"
	"github.com/Benbentwo/utils/log"
	"github.com/Benbentwo/utils/util"
	"github.com/go-errors/errors"
	"github.com/spf13/cobra"
	_ "gopkg.in/yaml.v2"
	_ "k8s.io/apimachinery/pkg/util/yaml"
	"os"
	_ "path/filepath"
	"runtime"
	"strings"
)

const (
	BB_HOME_VAR   = "BB_HOME"
	BB_CONFIG_DIR = "~/.bb"
	CONFIG        = "config.yaml"
)

type InitOptions struct {
	*common.CommonOptions
	Flags InitFlags
}

type InitFlags struct {
	ConfigDir   string
	ProjectsDir string
}

var logs = util.Logger()

var (
	initLong = `
		This Command will setup the configuration for later use.
		This will create configuration files in ~/.bb/ that later bb commands will use
`
	initExample = `
		bb init	
`
)

func NewCmdInit(commonOpts *common.CommonOptions) *cobra.Command {
	options := &InitOptions{
		CommonOptions: commonOpts,
	}
	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Initializes the " + BB_CONFIG_DIR + " configuration directory",
		Long:    initLong,
		Example: initExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}
	options.AddInitFlags(cmd)
	// Section to add commands to:
	return cmd
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}

func (o *InitOptions) Run() error {

	replacer := strings.NewReplacer("~", os.Getenv("HOME"))
	path := replacer.Replace(BB_CONFIG_DIR)

	log.Blank()

	// if it doesn't already exist create it
	exists, err := util.DirExists(path)
	if err != nil {
		return err
	}
	if !exists {
		logs.Debugf("Directory `~/.bb` not found... creating")
		err = os.MkdirAll(path, util.DefaultWritePermissions)
		if err != nil {
			return err
		}
	}

	if o.Experimental {
		// Add to the bash profile
		if os.Getenv(BB_HOME_VAR) != path {
			logs.Debugf("BB HOME Set to %s", os.Getenv(BB_HOME_VAR))
			logs.Debugf("Path Set to %s", path)
			// set current shell
			err = os.Setenv(BB_HOME_VAR, path)
			if err != nil {
				return err
			}

			response, err := util.Confirm("Would you like to update your bash profile?", true, "Set a variable on your bash profile?")
			if err != nil {
				return errors.Errorf("getting response: %s", err)
			}
			if response {

				stringExists, line, err := util.DoesFileContainString("export BB_HOME=~/.bb", "~/.bash_profile")
				if err != nil {
					return err
				}
				if line != -1 {
					logs.Debugf("Found String at line %s", line)
				}
				if !stringExists {
					f, err := os.OpenFile(replacer.Replace("~/.bash_profile"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
					if err != nil {
						logs.Errorf("Couldn't open or find ~/.bash_profile")
						panic(err)
					}

					defer f.Close()
					var carriage = GetCarriageReturn()

					if _, err = f.WriteString("export BB_HOME=" + path + carriage); err != nil {
						panic(err)
					}
					logs.Debugf("Updated Bash Profile to include BB_HOME")
				}
			}

		}
	}

	err = CreateFileIfNotFound(path + "/" + CONFIG)
	check(err)

	configPath := path + "/" + github.GitAuthConfigFile
	err = CreateFileIfNotFound(configPath)
	check(err)

	// Lets setup a Git Profile
	logs.Info("Looks like you do not have any git servers configured")
	response, err := util.Confirm("Would you like to set one up now?", true, "Would you like to create a connection configuraiton to a git server?")
	if err != nil {
		return errors.Errorf("getting response from confirmation: %s", err)
	}
	if response {
		SetupGitConfigFile(configPath, *o.CommonOptions)
	}
	logs.Println()
	logs.Infof("SUCCESS: BB Directory configured to %s", path)
	return nil
}

func CreateFileIfNotFound(configPath string) error {
	exists, err := util.FileExists(configPath)
	check(err)
	if exists {
		logs.Infof(configPath + " file found.")
	} else {
		logs.Infof(configPath + " file NOT found, creating...")
		_, err = os.Create(configPath)
		if err != nil {
			return err
		}
	}
	return nil
}
func (o *InitOptions) AddInitFlags(cmd *cobra.Command) {
	// add flags
	cmd.Flags().StringVarP(&o.Flags.ProjectsDir, "project-dir", "p", "~/dev", "The Directory you would like to store your Projects in")

}

// Writes a string to a file and returns whether or not it did exist
func WriteStringIfDoesntExist(writeString string, filePath string) bool {
	if exists, _, _ := util.DoesFileContainString(writeString, filePath); !exists {
		f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		check(err)
		_, err = f.WriteString(writeString)
		check(err)
		return false
	}
	return true
}

func GetCarriageReturn() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	} else {
		return "\n"
	}
}
func SetupGitConfigFile(configPath string, o common.CommonOptions) {
	carriage := GetCarriageReturn()
	WriteStringIfDoesntExist("currentserver: "+carriage, configPath)
	WriteStringIfDoesntExist("defaultusername: "+carriage, configPath)
	hasServers := WriteStringIfDoesntExist("servers: "+carriage, configPath)

	if !hasServers {

		serverName, err := util.PickValue("Git Server Name:", "", "What would you like to name this gitServer?", false)
		if err != nil {
			panic(err)
		}
		kind, err := util.Pick("Which Git Server would you like to add", github.ServerTypes, "What is your remote repository kind?")
		if err != nil {
			panic(err)
		}
		defaultUrl := github.GetDefaultUrlFromGitServer(kind)
		url, err := util.PickValue("Git Server URL:", defaultUrl, "What would you like to name this gitServer?", false)
		if err != nil {
			panic(err)
		}
		gitServerOptions := github.CreateGitServerOptions{
			CreateOptions: github.CreateOptions{
				CommonOptions: &o,
				DisableImport: true,
				OutDir:        configPath,
			},
			Name: serverName,
			Kind: kind,
			URL:  url,
		}
		err = gitServerOptions.Run()
		check(err)
	}
}
