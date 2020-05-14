package setup

import (
	"errors"
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/Benbentwo/utils/util"
	"github.com/spf13/cobra"
)

// options for the command
type SetupGoOptions struct {
	*common.CommonOptions
	batch   bool
	AllMode bool
}

var (
	setup_go_long    = `setup gvm and gopath`
	setup_go_example = `bb setup go`
)

func NewCmdSetupGo(commonOpts *common.CommonOptions) *cobra.Command {
	options := &SetupGoOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "go",
		Short:   "setup go",
		Long:    setup_go_long,
		Example: setup_go_example,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}
	// cmd.Flags().BoolVarP(&options.AllMode, "all", "a", false, "Being setup from everything") // the line below (Section to...) is for the generate-function command to add a template_command to.

	options.AddSetupFlags(cmd)
	return cmd
}

// Run implements this command
func (o *SetupGoOptions) Run() error {
	util.Logger().Infof("Congratulations generating %s", o.Cmd.Use)
	return errors.New("ABC")
}

func (o *SetupGoOptions) AddSetupFlags(cmd *cobra.Command) {
	o.Cmd = cmd
}
