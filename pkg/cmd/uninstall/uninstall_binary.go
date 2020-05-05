package uninstall

import (
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// GetAddonOptions the command line options
type UninstallBinaryOptions struct {
	UninstallOptions
}

var (
	uninstall_binary_long = `
		Uninstalls the BB binary from your machine

`

	uninstall_binary_example = `
		# Uninstall the binary
		bb uninstall binary
	`
)

func NewCmdUninstallBinary(commonOpts *common.CommonOptions) *cobra.Command {
	options := &UninstallBinaryOptions{
		UninstallOptions: UninstallOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:     "binary [flags]",
		Short:   "Uninstalls the binary",
		Long:    uninstall_binary_long,
		Example: uninstall_binary_example,
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
func (o *UninstallBinaryOptions) Run() error {

	err := Binary()
	if err != nil {
		return errors.Wrapf(err, "Uninstall binary command failed.")
	}
	return nil
}
