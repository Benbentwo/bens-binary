package github

import (
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/Benbentwo/utils/util"
	"github.com/spf13/cobra"
)

// options for the command
type GithubCreate_profileOptions struct {
	*common.CommonOptions
	batch bool
}

var (
	githubCreateProfileLong = `
Create a github profile for GH or GHE and add to your ~/.bb folder
`

	githubCreateProfileExample = `
bb gh create profile
`)


func NewCmdGithubCreate_profile(commonOpts *common.CommonOptions) *cobra.Command {
	options := &GithubCreate_profileOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "profile",
		Short:   "create a github profile",
		Long:    githubCreateProfileLong,
		Example: githubCreateProfileExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}

	return cmd
}

// Run implements this command
func (o *GithubCreate_profileOptions) Run() error {
	util.Logger().Infof("Congratulations generating %s", o.Cmd.Use)
	return nil
}
