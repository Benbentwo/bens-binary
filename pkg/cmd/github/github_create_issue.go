package github

import (
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/Benbentwo/utils/util"
	"github.com/spf13/cobra"
)

// options for the command
type GithubCreate_issueOptions struct {
	*common.CommonOptions
	batch bool
}

var (
	githubCreateIssueLong = `
Creates a github issue
`

	githubCreateIssueExample = `
bb github create issue
`)


func NewCmdGithubCreate_issue(commonOpts *common.CommonOptions) *cobra.Command {
	options := &GithubCreate_issueOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "issue",
		Short:   "create github issue",
		Long:    githubCreateIssueLong,
		Example: githubCreateIssueExample,
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
func (o *GithubCreate_issueOptions) Run() error {
	util.Logger().Infof("Congratulations generating %s", o.Cmd.Use)
	return nil
}
