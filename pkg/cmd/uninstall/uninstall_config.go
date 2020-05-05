package uninstall

import (
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// GetAddonOptions the command line options
type ConfigOptions struct {
	UninstallOptions
}

var (
	uninstallConfigLong = `
		Wipes the BB Binary's configuration from your machine,
		this allows you to reconfigure it with new settings.

		essentially runs "rm -rf ~/.bb"

`
	uninstallConfigExample = `
		# Uninstall the configuration
		bb uninstall config
	`
)

func NewCmdUninstallConfig(commonOpts *common.CommonOptions) *cobra.Command {
	options := &ConfigOptions{
		UninstallOptions: UninstallOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:     "config [flags]",
		Short:   "Uninstalls the binary",
		Long:    uninstallConfigLong,
		Example: uninstallConfigExample,
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
func (o *ConfigOptions) Run() error {

	err := Config()
	if err != nil {
		return errors.Wrapf(err, "Uninstall Config command failed.")
	}
	return nil
}
