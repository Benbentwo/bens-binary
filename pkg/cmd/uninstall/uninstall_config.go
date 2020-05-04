package uninstall

import (
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"

	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
)

// GetAddonOptions the command line options
type UninstallConfigOptions struct {
	UninstallOptions
}

var (
	uninstall_config_long = templates.LongDesc(`
		Wipes the BB Binary's configuration from your machine,
		this allows you to reconfigure it with new settings.

		essentially runs "rm -rf ~/.bb"

`)

	uninstall_config_example = templates.Examples(`
		# Uninstall the configuration
		bb uninstall config
	`)
)

func NewCmdUninstallConfig(commonOpts *common.CommonOptions) *cobra.Command {
	options := &UninstallConfigOptions{
		UninstallOptions: UninstallOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:     "config [flags]",
		Short:   "Uninstalls the binary",
		Long:    uninstall_config_long,
		Example: uninstall_config_example,
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
func (o *UninstallConfigOptions) Run() error {

	err := UninstallConfig()
	if err != nil {
		return errors.Wrapf(err, "Uninstall Config command failed.")
	}
	return nil
}
