package main

import (
  "fmt"
  "log"
  "os"

  "github.com/urfave/cli/v2"
)

func main() {
  app := &cli.App{
    Name:  "scaf",
    Usage: "SCAF - Software Co-working Assistance Framework",
    Commands: []*cli.Command{
      {
        Name:    "test",
        Aliases: []string{"t"},
        Usage:   "testing",
        Flags: []cli.Flag{
          &cli.StringFlag{
            Name:    "flag-test",
            Aliases: []string{"f"},
            Usage:   "testing flag",
          },
        },
        Action: func(c *cli.Context) error {
          fmt.Println("test command")
          fmt.Println(c.String("flag-test"))
          return nil
        },
        Subcommands: []*cli.Command{
          {
            Name:    "subtest",
            Aliases: []string{"st"},
            Usage:   "testing subcommand",
            Flags: []cli.Flag{
              &cli.BoolFlag{
                Name:    "flag-subtest",
                Aliases: []string{"f"},
                Usage:   "testing subcommand flag",
              },
            },
            Action: func(c *cli.Context) error {
              fmt.Println("subtest command")
              fmt.Println(c.String("flag-test"))
              fmt.Println(c.Bool("flag-subtest"))
              return nil
            },
          },
        },
      },
    },
  }

  if err := app.Run(os.Args); err != nil {
    log.Fatal(err)
  }
}
