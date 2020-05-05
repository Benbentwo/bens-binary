package github

import (
	"github.com/Benbentwo/bb/pkg/utilities"
	"path/filepath"
	"strings"
)

// Config gets the AuthConfig from the service
func (s *AuthConfigService) Config() *AuthConfig {
	if s.config == nil {
		s.config = &AuthConfig{}
	}
	return s.config
}

// SetConfig sets the AuthConfig object
func (s *AuthConfigService) SetConfig(c *AuthConfig) {
	s.config = c
}

// SaveUserAuth saves the given user auth for the server url
func (s *AuthConfigService) SaveUserAuth(url string, userAuth *UserAuth) error {
	config := s.config
	config.SetUserAuth(url, userAuth)
	user := userAuth.Username
	if user != "" {
		config.DefaultUsername = user
	}

	config.CurrentServer = url
	return s.saver.SaveConfig(s.config)
}

// DeleteServer removes the given server from the configuration
func (s *AuthConfigService) DeleteServer(url string) error {
	s.config.DeleteServer(url)
	return s.saver.SaveConfig(s.config)
}

// LoadConfig loads the configuration from the users JX config directory
func (s *AuthConfigService) LoadConfig() (*AuthConfig, error) {
	var err error
	s.config, err = s.saver.LoadConfig()
	return s.config, err
}

// SaveConfig saves the configuration to disk
func (s *AuthConfigService) SaveConfig() error {
	return s.saver.SaveConfig(s.Config())
}

// NewAuthConfigService generates a AuthConfigService with a custom saver. This should not be used directly
func NewAuthConfigService(saver ConfigSaver) *AuthConfigService {
	return &AuthConfigService{saver: saver}
}

// NewFileAuthConfigService
func NewFileAuthConfigService(filename string) (ConfigService, error) {
	saver, err := newFileAuthSaver(filename)
	return NewAuthConfigService(saver), err
}

// newFileBasedAuthConfigSaver creates a new FileBasedAuthConfigService that stores its data under the given filename
// If the fileName is an absolute path, it will be used. If it is a simple filename, it will be stored in the default
// Config directory
func newFileAuthSaver(fileName string) (ConfigSaver, error) {
	svc := &FileAuthConfigSaver{}
	// If the fileName is an absolute path, use that. Otherwise treat it as a config filename to be used in
	if fileName == filepath.Base(fileName) {
		dir, err := utilities.ConfigDir()
		if err != nil {
			return svc, err
		}
		svc.FileName = filepath.Join(dir, fileName)
	} else {
		svc.FileName = fileName
	}
	return svc, nil
}

// Base returns the last element of path.
// Trailing slashes are removed before extracting the last element.
// If the path is empty, Base returns ".".
// If the path consists entirely of slashes, Base returns "/".
func Base(path string) string {
	if path == "" {
		return "."
	}
	// Strip trailing slashes.
	for len(path) > 0 && path[len(path)-1] == '/' {
		path = path[0 : len(path)-1]
	}
	// Find the last element
	if i := strings.LastIndex(path, "/"); i >= 0 {
		path = path[i+1:]
	}
	// If empty now, it had only slashes.
	if path == "" {
		return "/"
	}
	return path
}
