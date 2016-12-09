package utils

import (
  "github.com/Sirupsen/logrus"
  "github.com/docker/libcompose/docker/client"
  "github.com/docker/libcompose/docker/ctx"
  "github.com/urfave/cli"
)

// DockerClientFlags defines the flags that are specific to the docker client,
// like configdir or tls related flags.
func DockerClientFlags() []cli.Flag {
  return []cli.Flag{
    cli.BoolFlag{
      Name:  "tls",
      Usage: "Use TLS; implied by --tlsverify",
    },
    cli.BoolFlag{
      Name:   "tlsverify",
      Usage:  "Use TLS and verify the remote",
      EnvVar: "DOCKER_TLS_VERIFY",
    },
    cli.StringFlag{
      Name:  "tlscacert",
      Usage: "Trust certs signed only by this CA",
    },
    cli.StringFlag{
      Name:  "tlscert",
      Usage: "Path to TLS certificate file",
    },
    cli.StringFlag{
      Name:  "tlskey",
      Usage: "Path to TLS key file",
    },
    cli.StringFlag{
      Name:  "configdir",
      Usage: "Path to docker config dir, default ${HOME}/.docker",
    },
  }
}

// Populate updates the specified docker context based on command line arguments and subcommands.
func Populate(context *ctx.Context, c *cli.Context) {
  context.ConfigDir = c.String("configdir")

  opts := client.Options{}
  opts.TLS = c.GlobalBool("tls")
  opts.TLSVerify = c.GlobalBool("tlsverify")
  opts.TLSOptions.CAFile = c.GlobalString("tlscacert")
  opts.TLSOptions.CertFile = c.GlobalString("tlscert")
  opts.TLSOptions.KeyFile = c.GlobalString("tlskey")

  clientFactory, err := client.NewDefaultFactory(opts)
  if err != nil {
    logrus.Fatalf("Failed to construct Docker client: %v", err)
  }

  context.ClientFactory = clientFactory
}

func DockerBuildFlags(flags []cli.Flag) []cli.Flag {
  newFlags := []cli.Flag{
    cli.StringFlag{
      Name: "label",
      Usage: "Name of label to use for related containers",
      Value: "dk.umwelt.build",
    },
    cli.StringFlag{
      Name: "p, project-name",
      Usage: "Project name and label to use for related containers",
    },
    /*
    cli.StringFlag{
      Name: "properties",
      Usage: "Build properties file",
      EnvVar: "TEAMCITY_BUILD_PROPERTIES_FILE",
    },
    cli.StringFlag{
      Name: "build-number",
      Usage: "Build number",
      EnvVar: "BUILD_NUMBER",
    },
    cli.StringFlag{
      Name: "build-type",
      Usage: "Build type",
      EnvVar: "BUILD_TYPE",
    },
    */
  }
  return append(flags, newFlags...)
}

func PopulateBuild(c *cli.Context) (string, string, error) {
  /*
  buildType := c.String("build-type")
  buildNumber := c.String("build-number")
  if buildType == "" || buildNumber == "" {
    if !c.IsSet("properties") {
      return "", "", errors.New("properties is required when build-type or build-number is not specified")
    }
    file, err := os.Open(c.String("properties"))
    if err != nil {
      return "", "", err
    }

    properties, err := props.Read(file)
    if err != nil {
      return "", "", err
    }
    if buildType == "" {
      buildType = properties.Get("teamcity.buildType.id")
    }
    if buildNumber == "" {
      buildNumber = properties.Get("build.number")
    }
  }
  */
  return c.String("project-name"), c.String("label"), nil
}
