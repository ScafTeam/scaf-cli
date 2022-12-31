package main

import (
  "log"
  "os"

  "github.com/urfave/cli/v2"

  "scaf/cli/user"
  "scaf/cli/config"
  "scaf/cli/action"
)

func run(args []string) {
  app := &cli.App{
    Name:  "scaf",
    Usage: "SCAF - Software Co-working Assistance Framework",
    Commands: []*cli.Command{
      {
        Name:    "user",
        Usage:   "user",
        Action:  user.GetUserAction,
      },
      {
        Name:    "signin",
        Usage:   "signin to SCAF",
        Action:  action.SignInAction,
        Flags: []cli.Flag{
          &cli.BoolFlag{
            Name:  "forget-password",
            Usage: "forget password",
          },
        },
      },
      {
        Name:    "signup",
        Usage:   "signup to SCAF",
        Action:  action.SignUpAction,
      },
      {
        Name:    "signout",
        Usage:   "signout from SCAF",
        Action:  action.SignOutAction,
      },
      {
        Name:    "config",
        Usage:   "configure SCAF",
        Subcommands: []*cli.Command{
          {
            Name:    "set",
            Usage:   "set config",
            Action:  config.SetConfigAction,
          },
          {
            Name:    "get",
            Usage:   "get config",
            Action:  config.GetConfigAction,
          },
          {
            Name:    "password",
            Usage:   "change password",
            Action:  config.ChangePasswordAction,
          },
        },
      },
      {
        Name:    "project",
        Usage:   "manage projects",
        Subcommands: []*cli.Command{
          {
            Name:    "list",
            Usage:   "list projects",
            Flags: []cli.Flag{
              &cli.BoolFlag{
                Name:  "oneline",
                Usage: "print oneline",
              },
            },
            Action:  action.ListProjectsAction,
          },
          {
            Name:    "create",
            Usage:   "create a new project",
            Action:  action.CreateProjectAction,
          },
          {
            Name:    "clone",
            Usage:   "clone a project",
            Action:  action.CloneProjectAction,
          },
        },
      },
      {
        Name:    "repo",
        Usage:   "manage code repositories",
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
      {
        Name:    "whoami",
        Usage:   "show current user",
        Action:  action.WhoamiAction,
      },
    },
  }

  if err := app.Run(args); err != nil {
    log.Fatal(err)
  }
}

func main() {
  os.Setenv("SCAF_BACKEND_URL", "http://localhost:8000")
  home_dir, _ := os.UserHomeDir()
  os.Setenv("SCAF_CONFIG_DIR", home_dir + "/.scaf")
  os.MkdirAll(os.Getenv("SCAF_CONFIG_DIR"), 0777)
  run(os.Args)
}

func notImplemented(c *cli.Context) error {
  return cli.Exit("not implemented", 1)
}
