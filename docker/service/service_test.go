package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zhgqiang/libcompose/config"
)

func TestSpecifiesHostPort(t *testing.T) {
	servicesWithHostPort := []Service{
		{serviceConfig: &config.ServiceConfig{Ports: []string{"8000:8000"}}},
		{serviceConfig: &config.ServiceConfig{Ports: []string{"127.0.0.1:8000:8000"}}},
	}

	for _, service := range servicesWithHostPort {
		assert.True(t, service.specificiesHostPort())
	}

	servicesWithoutHostPort := []Service{
		{serviceConfig: &config.ServiceConfig{Ports: []string{"8000"}}},
		{serviceConfig: &config.ServiceConfig{Ports: []string{"127.0.0.1::8000"}}},
	}

	for _, service := range servicesWithoutHostPort {
		assert.False(t, service.specificiesHostPort())
	}
}
