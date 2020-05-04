package uninstall

import (
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
)

// GetAddonOptions the command line options
type UninstallBinaryOptions struct {
	UninstallOptions
}

var (
	uninstall_binary_long = templates.LongDesc(`
		Uninstalls the BB binary from your machine

`)

	uninstall_binary_example = templates.Examples(`
		# Uninstall the binary
		bb uninstall binary
	`)
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

	err := UninstallBinary()
	if err != nil {
		return errors.Wrapf(err, "Uninstall binary command failed.")
	}
	return nil
}
