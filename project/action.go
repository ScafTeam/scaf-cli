package project

import (
  "github.com/urfave/cli/v2"
  "scaf/cli/scafio"
)

func ListProjectsAction(c *cli.Context) error {
  email, err := scafio.GetEmail(c)
  if err != nil {
    return err
  }
  projects, err := getProjects(email)
  if err != nil {
    return err
  }
  for _, project := range projects {
    projectMap := project.(map[string]interface{})
    scafio.PrintProject(projectMap)
  }

  return nil
}
