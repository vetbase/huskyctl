package main

import (
    "fmt"
    "os"
    "os/exec"

    "github.com/urfave/cli"
)

func main() {
    // create CLI application
    app := cli.NewApp()

    // set fields
    app.Name        = "Husky"
    app.Usage       = "Microframework for developing services"
    app.Version     = "0.1.0"
    app.Commands    = registerCommands()

    // return CLI app
    app.Run(os.Args)
}

func registerCommands() []cli.Command {
    return []cli.Command{
        {
            Name:       "run",
            Aliases:    []string{"t"},
            Usage:      "Run Husky service",
            Action:     runService,
        },
    }
}

func runService(c *cli.Context) error {
    output,err := exec.Command("./" + c.Args().Get(0))

    if err == nil {
        fmt.Println(output)
    } else {
        fmt.Println(err)
    }

    return nil
}
