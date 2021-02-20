package main

import (
    "github.com/sirupsen/logrus"
    "github.com/urfave/cli"
    "os"
)

const usage = `docker_go`

func main() {
    app := cli.NewApp()
    app.Name = "docker_go"
    app.Usage = usage

    app.Commands = []cli.Command{
        runCommand,
        initCommand,
        execCommand,
        logCommand,
        listCommand,
        stopCommand,
        removeCommand,
    }
    app.Before = func(context *cli.Context) error {
        logrus.SetFormatter(&logrus.JSONFormatter{})
        logrus.SetOutput(os.Stdout)
        return nil
    }
    if err := app.Run(os.Args); err != nil {
        logrus.Fatal(err)
    }
}
