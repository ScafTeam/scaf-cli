package main

import (
  "log"
  "os"

  "github.com/urfave/cli/v2"

  "scaf/cli/auth"
  "scaf/cli/user"
  "scaf/cli/project"
  "scaf/cli/config"
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
        Action:  auth.SignInAction,
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
        Action:  auth.SignUpAction,
      },
      {
        Name:    "signout",
        Usage:   "signout from SCAF",
        Action:  auth.SignOutAction,
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
        },
      },
      {
        Name:    "project",
        Usage:   "manage projects",
        Action:  project.ListProjectsAction,
        Subcommands: []*cli.Command{
          {
            Name:    "list",
            Usage:   "list projects",
            Action:  project.ListProjectsAction,
          },
          {
            Name:    "create",
            Usage:   "create a new project",
            Action:  project.CreateProjectAction,
          },
        },
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
      {
        Name:    "whoami",
        Usage:   "show current user",
        Action:  auth.WhoamiAction,
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
