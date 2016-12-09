package services

import (
  "fmt"
    "golang.org/x/net/context"

    "github.com/docker/libcompose/docker"
    "github.com/docker/libcompose/docker/ctx"
    "github.com/docker/libcompose/project"
    "github.com/docker/libcompose/project/options"
    "github.com/umweltdk/teamcity-steps/docker/utils"
    "github.com/urfave/cli"
)

func Action() cli.Command {
  return cli.Command{
    Name:    "services",
    Usage:   "start services configured in .umfig.yml",
    Action:  ServicesAction,
    Flags: utils.DockerBuildFlags([]cli.Flag{}),
  }
}

func ServicesAction(c *cli.Context) error {
  fmt.Println("added task: ", c.Args().First())
  projectName, label, err := utils.PopulateBuild(c)
  if err != nil {
    return nil
  }

  dockerContext := &ctx.Context{
    Context: project.Context{
      ComposeFiles: []string{".umfig.yml"},
      ProjectName:  projectName,
    },
  }
  utils.Populate(dockerContext, c)
  dockerContext.ServiceFactory = NewFactory(dockerContext, label, projectName)
  project, err := docker.NewProject(dockerContext, nil)

  if err != nil {
      return err
  }

  err = project.Up(context.Background(), options.Up{})
  return err
}