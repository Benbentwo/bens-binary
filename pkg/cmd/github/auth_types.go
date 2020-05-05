package github

import (
	"fmt"
	"github.com/Benbentwo/utils/util"
	"io/ioutil"
	"sigs.k8s.io/yaml"
)

const (
	GitAuthConfigFile = "gitAuth.yaml"
)

type AuthServer struct {
	URL   string      `json:"url"`
	Users []*UserAuth `json:"users"`
	Name  string      `json:"name"`
	Kind  string      `json:"kind"`

	CurrentUser string `json:"currentuser"`
}

type UserAuth struct {
	Username    string `json:"username"`
	ApiToken    string `json:"apitoken"`
	BearerToken string `json:"bearertoken"`
	Password    string `json:"password,omitempty"`
}

type AuthConfig struct {
	Servers []*AuthServer `json:"servers"`

	DefaultUsername string `json:"defaultusername"`
	CurrentServer   string `json:"currentserver"`
}

// AuthConfigService implements the generic features of the ConfigService because we don't have superclasses
type AuthConfigService struct {
	config *AuthConfig
	saver  ConfigSaver
}

// FileAuthConfigSaver is a ConfigSaver that saves its config to the local filesystem
type FileAuthConfigSaver struct {
	FileName string
}

// LoadConfig loads the configuration from the users JX config directory
func (s *FileAuthConfigSaver) LoadConfig() (*AuthConfig, error) {
	config := &AuthConfig{}
	fileName := s.FileName
	if fileName != "" {
		exists, err := util.FileExists(fileName)
		if err != nil {
			return config, fmt.Errorf("Could not check if file exists %s due to %s", fileName, err)
		}
		if exists {
			data, err := ioutil.ReadFile(fileName)
			if err != nil {
				return config, fmt.Errorf("Failed to load file %s due to %s", fileName, err)
			}
			err = yaml.Unmarshal(data, config)
			if err != nil {
				return config, fmt.Errorf("Failed to unmarshal YAML file %s due to %s", fileName, err)
			}
		}
	}
	return config, nil
}

// SaveConfig saves the configuration to disk
func (s *FileAuthConfigSaver) SaveConfig(config *AuthConfig) error {
	fileName := s.FileName
	if fileName == "" {
		return fmt.Errorf("no filename defined")
	}
	data, err := yaml.Marshal(config)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, data, util.DefaultWritePermissions)
}

// MemoryAuthConfigSaver uses memory
type MemoryAuthConfigSaver struct {
	config AuthConfig
}
