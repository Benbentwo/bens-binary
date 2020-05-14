package setup

import (
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/Benbentwo/utils/util"
	"github.com/spf13/cobra"
)

// options for the command
type SetupOptions struct {
	*common.CommonOptions
	// All bool
}

func NewCmdSetup(commonOpts *common.CommonOptions) *cobra.Command {
	options := &SetupOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use: "setup",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}

	// cmd.Flags().BoolVarP(&options.All, "all", "a", false, "Setup everything - runs all sub commands in the right order") // the line below (Section to...) is for the generate-function command to add a template_command to.

	// Section to add commands to:
	cmd.AddCommand(NewCmdSetupGo(commonOpts))

	options.AddSetupFlags(cmd)
	return cmd
}

// Run implements this command
func (o *SetupOptions) Run() error {
	// if !o.All {
	// 	return o.Cmd.Help()
	// }
	// o.All = false
	if !o.BatchMode {
		choice, err := util.Confirm("Print Help and exit?", true, "Do you want to learn about this command or run it? Yes to run help")
		if choice || err != nil {
			util.Logger().Errorf("Confirmation failed: %s", err)
			return o.Cmd.Help()
		}
	}

	failedCommands := make([]*cobra.Command, 0)

	// child := NewCmdSetupGo(o.CommonOptions)
	// o.Cmd.TraverseChildren = true
	// child.ResetFlags()
	// c, err := child.ExecuteC()
	// // util.Logger().Debugf("c: %s", c)
	// err := child.RunE(child, o.Args)
	// // err := o.Cmd.Run(child, o.Args)
	// if err != nil {
	// 	util.Logger().Errorf(util.ColorInfo("Setup "+child.Use) + " failed with: %s", err)
	// 	failedCommands = append(failedCommands, child)
	// }
	children := o.Cmd.Commands()
	for _, child := range children {
		// err := child.Execute()
		err := child.RunE(child, make([]string, 0))
		// err :=
		if err != nil {
			util.Logger().Errorf(util.ColorInfo("Setup "+child.Use)+" failed with: %s", err)
			failedCommands = append(failedCommands, child)
		}
	}

	return nil
}

func (o *SetupOptions) AddSetupFlags(cmd *cobra.Command) {
	o.Cmd = cmd
}
