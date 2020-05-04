package generate

import (
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/Benbentwo/utils/util"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// GetAddonOptions the command line options
type GenerateSearchFileOptions struct {
	GenerateOptions

	SearchString string
	SearchFile   string
}

var (
	utilSearchFileLong = `
		Searches a file for a string and prints all line numbers

`

	utilSearchFileExample = `
		# Utility to search a file for a string
		bb util search "hello world" HelloWorld.java
	`
)

func NewCmdUtilitySearchFile(commonOpts *common.CommonOptions) *cobra.Command {
	options := &GenerateSearchFileOptions{
		GenerateOptions: GenerateOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:     "searchfile String File [flags]",
		Short:   "Searches a file for a string",
		Long:    utilSearchFileLong,
		Example: utilSearchFileExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}
	cmd.Flags().StringVarP(&options.SearchString, "search-string", "s", "", "Search for this string")
	cmd.Flags().StringVarP(&options.SearchFile, "search-file", "f", "", "Search this file for the string")

	return cmd
}

// Run implements this command
func (o *GenerateSearchFileOptions) Run() error {
	if o.SearchFile == "" {
		return util.MissingOption("search-file")
	}
	if o.SearchString == "" {
		return util.MissingOption("search-string")
	}
	o.SearchFile = util.HomeReplace(o.SearchFile)
	count, err := util.FindMatchesInFile(o.SearchString, o.SearchFile)
	if err != nil {
		return errors.Wrapf(err, "Could not search the file %s for the string %s.", o.SearchFile, o.SearchString)
	}
	util.Logger().Infof("Found %d instances of `%s` in the file `%s`", len(count), o.SearchString, o.SearchFile)
	if len(count) > 0 {
		util.Logger().Infof("Found on lines %d", count)
	}
	return nil
}
