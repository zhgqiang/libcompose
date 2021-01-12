package project

import (
	"golang.org/x/net/context"

	"github.com/zhgqiang/libcompose/project/events"
	"github.com/zhgqiang/libcompose/project/options"
)

// Build builds the specified services (like docker build).
func (p *Project) Build(ctx context.Context, buildOptions options.Build, services ...string) error {
	return p.perform(events.ProjectBuildStart, events.ProjectBuildDone, services, wrapperAction(func(wrapper *serviceWrapper, wrappers map[string]*serviceWrapper) {
		wrapper.Do(wrappers, events.ServiceBuildStart, events.ServiceBuild, func(service Service) error {
			return service.Build(ctx, buildOptions)
		})
	}), nil)
}
