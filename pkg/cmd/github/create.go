package github

import (
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/spf13/cobra"
)

// options for the command
type CreateOptions struct {
	*common.CommonOptions
	DisableImport bool
	OutDir        string
}

func NewCmdCreate(commonOpts *common.CommonOptions) *cobra.Command {
	options := &CreateOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use: "create",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}
	options.AddCreateFlags(cmd)
	// the line below (Section to...) is for the generate-function command to add a template_command to.
	// the line above this and below can be deleted.
	// DO NOT DELETE THE FOLLOWING LINE:
	// Section to add commands to:
	cmd.AddCommand(NewCmdCreateGitServerConfig(commonOpts))
	cmd.AddCommand(NewCmdGithubCreate_profile(commonOpts))
	cmd.AddCommand(NewCmdGithubCreate_issue(commonOpts))

	return cmd
}

// Run implements this command
func (o *CreateOptions) Run() error {
	return o.Cmd.Help()
}

func (o *CreateOptions) AddCreateFlags(cmd *cobra.Command) {
	o.Cmd = cmd
}
