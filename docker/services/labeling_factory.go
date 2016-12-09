package services

import (
  "errors"
  "github.com/docker/libcompose/config"
  "github.com/docker/libcompose/docker/ctx"
  "github.com/docker/libcompose/project"
  "github.com/docker/libcompose/docker/service"
  "github.com/docker/libcompose/yaml"
  "reflect"
)

// Factory is an implementation of project.ServiceFactory.
type Factory struct {
  context *ctx.Context
  label string
  labelValue string
}

// NewFactory creates a new service factory for the given context
func NewFactory(context *ctx.Context, label, labelValue string) *Factory {
  return &Factory{
    context: context,
    label: label,
    labelValue: labelValue,
  }
}

// Create creates a Service based on the specified project, name and service configuration.
func (s *Factory) Create(project *project.Project, name string, serviceConfig *config.ServiceConfig) (project.Service, error) {
  if serviceConfig.Labels == nil {
    serviceConfig.Labels = make(yaml.SliceorMap)
  }
  serviceConfig.Labels[s.label] = s.labelValue
  if !reflect.DeepEqual(serviceConfig.Build, yaml.Build{}) {
    return nil, errors.New("services do not support docker build")
  }
  
  return service.NewService(name, serviceConfig, s.context), nil
}