package main // import "github.com/umweltdk/teamcity-steps"

import (
    "github.com/urfave/cli"
    "github.com/umweltdk/teamcity-steps/docker"

    "os"
)


func main() {
    app := cli.NewApp()
    app.Commands = []cli.Command{
        docker.Action(),
    }
    app.Run(os.Args)
}