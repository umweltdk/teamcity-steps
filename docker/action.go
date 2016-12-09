package docker

import (
/*
  "fmt"
    "golang.org/x/net/context"

    "github.com/docker/docker/client"
    "github.com/docker/libcompose/docker"
    "github.com/docker/libcompose/docker/ctx"
    "github.com/docker/libcompose/project"
    "github.com/docker/libcompose/project/options"
    "github.com/rickar/props"
    */
    "github.com/urfave/cli"
    //"os"

    "github.com/umweltdk/teamcity-steps/docker/services"
    "github.com/umweltdk/teamcity-steps/docker/cleanup"
    "github.com/umweltdk/teamcity-steps/docker/utils"
)

func Action() cli.Command {
  return cli.Command{
    Name:    "docker",
    Usage:   "docker commands",
    Subcommands: []cli.Command{
      services.Action(),
      cleanup.Action(),
    },
    Flags: utils.DockerClientFlags(),
  }
}

