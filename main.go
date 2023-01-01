package main

import (
  "log"
  "os"
  "github.com/urfave/cli/v2"
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
        Action:  action.GetUserAction,
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
            Action:  action.SetConfigAction,
          },
          {
            Name:    "get",
            Usage:   "get config",
            Action:  action.GetConfigAction,
          },
          {
            Name:    "password",
            Usage:   "change password",
            Action:  action.ChangePasswordAction,
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
            Action:  action.ListProjectAction,
          },
          {
            Name:    "create",
            Usage:   "create a new project, and clone it",
            Action:  action.CreateProjectAction,
          },
          {
            Name:    "delete",
            Usage:   "delete a project",
            Action:  action.DeleteProjectAction,
          },
          {
            Name:    "clone",
            Usage:   "clone a project",
            Action:  action.CloneProjectAction,
          },
          {
            Name:    "pull",
            Usage:   "pull current project, need to be in a project folder",
            Action:  action.PullProjectAction,
          },
        },
      },
      {
        Name:    "repo",
        Usage:   "manage code repositories, need to be in a project folder",
        Subcommands: []*cli.Command{
          {
            Name:    "list",
            Usage:   "list repositories",
            Action:  action.ListRepoAction,
          },
          {
            Name:    "add",
            Usage:   "add a repository",
            Action:  action.AddRepoAction,
          },
          {
            Name:    "update",
            Usage:   "update a repository",
            Action:  action.UpdateRepoAction,
          },
          {
            Name:    "delete",
            Usage:   "delete a repository",
            Action:  action.DeleteRepoAction,
          },
        },
      },
      {
        Name:    "doc",
        Usage:   "manage documents",
        Action:  notImplemented,
      },
      {
        Name:    "kanban",
        Usage:   "manage kanban boards",
        Subcommands: []*cli.Command{
          {
            Name:    "list",
            Usage:   "list kanban boards",
            Action:  action.ListWorkflowAction,
            Flags: []cli.Flag{
              &cli.BoolFlag{
                Name:  "oneline",
                Usage: "print oneline",
              },
            },
          },
          {
            Name:    "add",
            Usage:   "add a kanban board",
            Action:  action.AddWorkflowAction,
          },
        },
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
