package github

import (
	"github.com/Benbentwo/bens-binary/pkg/cmd/common"
	"github.com/Benbentwo/utils/util"
	"github.com/spf13/cobra"
)

// options for the command
type CreateGitServerOptions struct {
	CreateOptions
	Name string
	Kind string
	URL  string
}

var (
	githubCreateGitserverLong = `
Adds a git server object to your ~/.bb/gitAuth.yaml configuration. You can add different users to each git server. A git server is like github.com or github.com, there is also bitbucket support
`

	githubCreateGitserverExample = `
bb gh create git-server
`
)

func NewCmdCreateGitServerConfig(commonOpts *common.CommonOptions) *cobra.Command {
	options := &CreateGitServerOptions{
		CreateOptions: CreateOptions{
			CommonOptions: commonOpts,
		},
	}

	cmd := &cobra.Command{
		Use:     "gitserver",
		Short:   "create a git server in your configuration",
		Long:    githubCreateGitserverLong,
		Example: githubCreateGitserverExample,
		Run: func(cmd *cobra.Command, args []string) {
			options.Cmd = cmd
			options.Args = args
			err := options.Run()
			common.CheckErr(err)
		},
	}

	cmd.Flags().StringVarP(&options.Name, "name", "n", "", "The name for the Git server being created")
	cmd.Flags().StringVarP(&options.Kind, "kind", "k", "", "The kind of Git server being created")
	cmd.Flags().StringVarP(&options.URL, "url", "u", "", "The git server URL")

	return cmd
}

// Run implements this command
func (o *CreateGitServerOptions) Run() error {
	args := o.Args
	kind := o.Kind
	if kind == "" {
		if len(args) < 1 {
			return util.MissingOption("kind")
		}
		kind = args[0]
	}
	name := o.Name
	if name == "" {
		name = kind
	}
	gitUrl := o.URL
	if gitUrl == "" {
		if len(args) > 1 {
			gitUrl = args[1]
		} else {
			// lets try find the git URL based on the provider
			gitUrl = GetDefaultUrlFromGitServer(kind)
			// if serviceName != "" {
			// 	url, err := o.FindService(serviceName)
			// 	if err != nil {
			// 		return errors.Wrapf(err, "Failed to find %s Git service %s", kind, serviceName)
			// 	}
			// 	gitUrl = url
			// }
		}
	}
	if gitUrl == "" {
		return util.MissingOption("url")
	}
	// configService, err := auth.NewFileAuthConfigService(GitAuthConfigFile)
	// TODO reimplement this, its gross right now, do it right when I have time
	return common.ErrorUnimplemented()
	// authConfigSvc, err := clients.Factory.CreateAuthConfigService(factory, GitAuthConfigFile, "")
	// if err != nil {
	// 	return errors.Wrap(err, "failed to create CreateGitAuthConfigService")
	// }
	//
	// util.Logger().Info("authconfigsvc: %s", authConfigSvc)
	// config := authConfigSvc.Config()
	// server := config.GetOrCreateServerName(gitUrl, name, kind)
	// util.Logger().Info("server: %s", server)
	// config.CurrentServer = gitUrl
	// err = authConfigSvc.SaveConfig()
	// if err != nil {
	// 	return errors.Wrap(err, "failed to save GitAuthConfigService")
	// }
	// util.Logger().Infof("Added Git server %s for URL %s", util.ColorInfo(name), util.ColorInfo(gitUrl))
	//
	// err = authConfigSvc.SaveConfig()
	// common.CheckErr(err)
	return nil
}
