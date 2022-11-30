package main

import (
  "log"
  "os"

  "github.com/urfave/cli/v2"
  "github.com/ScafTeam/scaf-cli/user"
)

func main() {
  app := &cli.App{
    Name:  "scaf",
    Usage: "SCAF - Software Co-working Assistance Framework",
    Commands: []*cli.Command{
      {
        Name:    "login",
        Usage:   "login to SCAF",
        Action:  user.Login,
        Flags: []cli.Flag{
          &cli.BoolFlag{
            Name:  "forget-password",
            Usage: "forget password",
          },
        },
      }, {
        Name:    "register",
        Usage:   "register to SCAF",
        Action:  user.Register,
      },
    },
  }

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
