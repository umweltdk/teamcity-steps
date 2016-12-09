package cleanup // github.com/umweltdk/teamcity-steps/docker/cleanup"

import (
/*
  "fmt"

    "github.com/docker/libcompose/docker"
    "github.com/docker/libcompose/project"
    "github.com/docker/libcompose/project/options"
    "github.com/rickar/props"
    */
    "golang.org/x/net/context"
    "github.com/docker/docker/api/types"
    "github.com/docker/docker/api/types/filters"
    "github.com/docker/docker/client"
    "github.com/docker/libcompose/docker/ctx"
    "github.com/Sirupsen/logrus"
    "github.com/umweltdk/teamcity-steps/docker/utils"
    "github.com/urfave/cli"
    "time"
)

func Action() cli.Command {
  return cli.Command{
    Name:    "cleanup",
    Usage:   "cleanup all containers started by the build",
    Action:  CleanupAction,
    Flags: utils.DockerBuildFlags([]cli.Flag{
      cli.StringFlag{
        Name: "tag",
        Usage: "Docker tag to clean up containers for",
        EnvVar: "DOCKER_TAG",
      },
      cli.DurationFlag{
        Name: "timeout",
        Usage: "Timeout to wait for stop",
        Value: time.Duration(10)*time.Second,
      },
    }),
  }
}

type Client struct {
  client client.APIClient
}

func (cli *Client) stopContainer(context context.Context, containerID string, timeout *time.Duration) error {
  logrus.Infof("Stopping %s", containerID)
  err := cli.client.ContainerStop(context, containerID, timeout)
  if err != nil {
    return err
  }
  logrus.Infof("Removing %s", containerID)
  err = cli.client.ContainerRemove(context, containerID, types.ContainerRemoveOptions{
    RemoveVolumes: true,
    Force: true,
  })
  return err
}

func (cli *Client) stopContainers(context context.Context, filters filters.Args, timeout *time.Duration) error {
  list, err := cli.client.ContainerList(context, types.ContainerListOptions{
    All: true,
    Filter: filters,
  })
  if err != nil {
    return err
  }
  for _, container := range list {
    err = cli.stopContainer(context, container.ID, timeout)
    if err != nil {
      return err
    }
  }
  return nil
}

func CleanupAction(c *cli.Context) error {
  dockerContext := ctx.Context{}
  utils.Populate(&dockerContext, c)
  client := Client{
    client: dockerContext.ClientFactory.Create(nil),
  }

  timeout := c.Duration("timeout")
  if c.IsSet("tag") {
    filter := filters.NewArgs()
    filter.Add("ancestor", c.String("tag"))
    err := client.stopContainers(context.Background(), filter, &timeout)
    if err != nil {
      return nil
    }
  } else {
    logrus.Info("Nothing to do")
  }

  projectName, label, err := utils.PopulateBuild(c)
  if err != nil {
    return err
  }

  filter := filters.NewArgs()
  logrus.Infof("Nothing to do %s=%s", label, projectName)
  filter.Add("label", label + "=" + projectName)
  err = client.stopContainers(context.Background(), filter, &timeout)
  if err != nil {
    return nil
  }

  return nil
}