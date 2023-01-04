package action

import (
  "fmt"
  "github.com/urfave/cli/v2"
)

func ShowQAAction(c *cli.Context) error {
  fmt.Println("Please refer to the following link for the QA process:")
  fmt.Println("https://tang-shao-xiansorganization.gitbook.io/ruan-ti-gong-cheng/")
  return nil
}
