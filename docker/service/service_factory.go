package service

import (
	"github.com/zhgqiang/libcompose/config"
	"github.com/zhgqiang/libcompose/docker/ctx"
	"github.com/zhgqiang/libcompose/project"
)

// Factory is an implementation of project.ServiceFactory.
type Factory struct {
	context *ctx.Context
}

// NewFactory creates a new service factory for the given context
func NewFactory(context *ctx.Context) *Factory {
	return &Factory{
		context: context,
	}
}

// Create creates a Service based on the specified project, name and service configuration.
func (s *Factory) Create(project *project.Project, name string, serviceConfig *config.ServiceConfig) (project.Service, error) {
	return NewService(name, serviceConfig, s.context), nil
}
