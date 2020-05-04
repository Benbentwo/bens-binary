package generate

import (
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/spf13/cobra"
)

type GenerateOptions struct {
	*common.CommonOptions
	Output string
}

func NewCmdGenerate(commonOpts *common.CommonOptions) *cobra.Command {
	options := &GenerateOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Generate Code or objects",
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}
	options.AddGenerateFlags(cmd)
	// Section to add commands to:
	cmd.AddCommand(NewCmdUtilitySearchFile(commonOpts))
	cmd.AddCommand(NewCmdGenerateFunction(commonOpts))
	return cmd
}

// Run implements this command
func (o *GenerateOptions) Run() error {
	return o.Cmd.Help()
}

func (o *GenerateOptions) AddGenerateFlags(cmd *cobra.Command) {
	o.Cmd = cmd
}
