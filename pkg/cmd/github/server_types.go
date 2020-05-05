package github

import (
	"github.com/Benbentwo/bb/pkg/log"
)

const (
	// KindBitBucketCloud git kind for BitBucket Cloud
	KindBitBucketCloud = "bitbucketcloud"
	// KindBitBucketServer git kind for BitBucket Server
	KindBitBucketServer = "bitbucketserver"
	// KindGitea git kind for gitea
	KindGitea = "gitea"
	// KindGitlab git kind for gitlab
	KindGitlab = "gitlab"
	// KindGitHub git kind for github
	KindGitHub = "github"
	// KindGitFake git kind for fake git
	KindGitFake = "fakegit"
	// KindUnknown git kind for unknown git
	KindUnknown = "unknown"
)

var ServerTypes = []string{KindBitBucketServer, KindGitHub, KindGitlab, KindGitea, KindBitBucketCloud}

var serverMapUrl = map[string]string{
	KindGitHub:         "https://github.com",
	KindBitBucketCloud: "http://bitbucket.org",
}

func GetDefaultUrlFromGitServer(kind string) string {
	retVal := serverMapUrl[kind]
	if retVal == "" {
		util.Logger().Warnf("Coudn't find a default value for that server type, Setting Default to \"\"")
		return ""
	}
	return retVal
}
