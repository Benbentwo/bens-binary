package github

import (
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/spf13/cobra"
)

// options for the command
type GhOptions struct {
	*common.CommonOptions
}

func NewCmdGh(commonOpts *common.CommonOptions) *cobra.Command {
	options := &GhOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use: "gh",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}
	options.AddGhFlags(cmd)
	// the line below (Section to...) is for the generate-function command to add a template_command to.
	// the line above this and below can be deleted.
	// DO NOT DELETE THE FOLLOWING LINE:
	// Section to add commands to:
	cmd.AddCommand(NewCmdCreate(commonOpts))

	return cmd
}

// Run implements this command
func (o *GhOptions) Run() error {
	return o.Cmd.Help()
}

func (o *GhOptions) AddGhFlags(cmd *cobra.Command) {
	o.Cmd = cmd
}
