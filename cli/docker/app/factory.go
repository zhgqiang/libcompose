package app

import (
	"github.com/urfave/cli"
	"github.com/zhgqiang/libcompose/cli/logger"
	"github.com/zhgqiang/libcompose/docker"
	"github.com/zhgqiang/libcompose/docker/ctx"
	"github.com/zhgqiang/libcompose/project"
)

// ProjectFactory is a struct that holds the app.ProjectFactory implementation.
type ProjectFactory struct {
}

// Create implements ProjectFactory.Create using docker client.
func (p *ProjectFactory) Create(c *cli.Context) (project.APIProject, error) {
	context := &ctx.Context{}
	context.LoggerFactory = logger.NewColorLoggerFactory()
	Populate(context, c)
	return docker.NewProject(context, nil)
}
