package uninstall

import (
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/spf13/cobra"
)

type UninstallOptions struct {
	*common.CommonOptions
	Output string
	All    bool
}

const (
	uninstallResources = `uninstall options include:

    * config
    * binary
    * all (e.g. 'all above')
    `
)

var (
	getLong = `
		Uninstalls one or more resources.

		` + uninstallResources + `

`

	getExample = `
		# uninstall the binary
			bb uninstall binary

		# Uninstall your current config **Cannot be undone**
			bb uninstall config

		# Uninstall EVERYTHING
			bb uninstall -a
			# or
			bb uninstall --all
	`
)

// NewCmdGet creates a command object for the generic "get" action, which
// retrieves one or more resources from a server.
func NewCmdUninstall(commonOpts *common.CommonOptions) *cobra.Command {
	options := &UninstallOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "uninstall TYPE [flags]",
		Short:   "Uninstall one or more resources",
		Long:    getLong,
		Example: getExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}
	options.AddUninstallFlags(cmd)
	// Section to add commands to:
	cmd.AddCommand(NewCmdUninstallBinary(commonOpts))
	cmd.AddCommand(NewCmdUninstallConfig(commonOpts))
	return cmd
}

// Run implements this command
func (o *UninstallOptions) Run() error {
	if o.All {
		err := UninstallAll()
		if err != nil {
			return err
		}
		return nil
	}
	return o.Cmd.Help()
}

func (o *UninstallOptions) AddUninstallFlags(cmd *cobra.Command) {
	o.Cmd = cmd
	cmd.Flags().BoolVarP(&o.All, "all", "a", false, "Uninstall everything")
}
