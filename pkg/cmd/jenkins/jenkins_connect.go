package jenkins

import (
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/Benbentwo/utils/util"
	"github.com/spf13/cobra"
)

// options for the command
type JenkinsConnectOptions struct {
	*common.CommonOptions
	batch bool
}

var (
	jenkinsConnectLong = `
Sets up a jenkins job for a repository and sets up webhooks for that repository
`

	jenkinsConnectExample = `
bb jenkins connect
`
)

func NewCmdJenkinsConnect(commonOpts *common.CommonOptions) *cobra.Command {
	options := &JenkinsConnectOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "connect",
		Short:   "connect a repo to jenkins",
		Long:    jenkinsConnectLong,
		Example: jenkinsConnectExample,
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
func (o *JenkinsConnectOptions) Run() error {
	util.Logger().Infof("Congratulations generating %s", o.Cmd.Use)
	return nil
}
