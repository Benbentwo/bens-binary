package jenkins

import (
	"github.com/Benbentwo/bb/pkg/log"
	"github.com/jenkins-x/jx/pkg/cmd/helper"
	"github.com/jenkins-x/jx/pkg/cmd/opts"
	"github.com/jenkins-x/jx/pkg/cmd/templates"
	"github.com/spf13/cobra"
)

// options for the command
type JenkinsJenkins_connectOptions struct {
	*common.CommonOptions
	batch bool
}

var (
	jenkins_jenkins_jenkins_connect_long = templates.LongDesc(`
Sets up a jenkins job for a repository and sets up webhooks for that repository
`)

	jenkins_jenkins_jenkins_connect_example = templates.Examples(`
bb jenkins connect
`)
)

func NewCmdJenkins_connect(commonOpts *common.CommonOptions) *cobra.Command {
	options := &JenkinsJenkins_connectOptions{
		CommonOptions: commonOpts,
	}

	cmd := &cobra.Command{
		Use:     "connect",
		Short:   "connect a repo to jenkins",
		Long:    jenkins_jenkins_jenkins_connect_long,
		Example: jenkins_jenkins_jenkins_connect_example,
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
func (o *JenkinsJenkins_connectOptions) Run() error {
	log.Logger().Infof("Congratulations generating %s", o.Cmd.Use)
	return nil
}
