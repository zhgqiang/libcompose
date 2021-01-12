package network

import (
	"github.com/zhgqiang/libcompose/config"
	composeclient "github.com/zhgqiang/libcompose/docker/client"
	"github.com/zhgqiang/libcompose/project"
)

// DockerFactory implements project.NetworksFactory
type DockerFactory struct {
	ClientFactory composeclient.Factory
}

// Create implements project.NetworksFactory Create method.
// It creates a Networks (that implements project.Networks) from specified configurations.
func (f *DockerFactory) Create(projectName string, networkConfigs map[string]*config.NetworkConfig, serviceConfigs *config.ServiceConfigs, networkEnabled bool) (project.Networks, error) {
	cli := f.ClientFactory.Create(nil)
	return NetworksFromServices(cli, projectName, networkConfigs, serviceConfigs, networkEnabled)
}
