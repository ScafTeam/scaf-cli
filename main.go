package main

import (
  "log"
  "os"

  "github.com/urfave/cli/v2"

  "scaf/cli/user"
)

func run(args []string) {
  app := &cli.App{
    Name:  "scaf",
    Usage: "SCAF - Software Co-working Assistance Framework",
    Commands: []*cli.Command{
      {
        Name:    "signin",
        Usage:   "signin to SCAF",
        Action:  user.SignInAction,
        Flags: []cli.Flag{
          &cli.BoolFlag{
            Name:  "forget-password",
            Usage: "forget password",
          },
        },
      },
      {
        Name:    "register",
        Usage:   "register to SCAF",
        Action:  user.Register,
      },
      {
        Name:    "config",
        Usage:   "configure SCAF",
        Action:  notImplemented,
      },
      {
        Name:    "project",
        Usage:   "manage projects",
        Action:  notImplemented,
      },
      {
        Name:    "repo",
        Usage:   "manage repositories",
        Action:  notImplemented,
      },
      {
        Name:    "doc",
        Usage:   "manage documents",
        Action:  notImplemented,
      },
      {
        Name:    "kanban",
        Usage:   "manage kanban boards",
        Action:  notImplemented,
      },
      {
        Name:    "qa",
        Usage:   "show Q&A",
        Action:  notImplemented,
      },
    },
  }

  if err := app.Run(args); err != nil {
    log.Fatal(err)
  }
}

func main() {
  run(os.Args)
}

func notImplemented(c *cli.Context) error {
  return cli.Exit("not implemented", 1)
}
