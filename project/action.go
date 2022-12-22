package project

import (
  "github.com/urfave/cli/v2"
  "github.com/AlecAivazis/survey/v2"
  "scaf/cli/scafio"
)

func ListProjectsAction(c *cli.Context) error {
  var err error
  questions := []*survey.Question{}
  answers := struct {
    Email string
  }{}

  answers.Email, err = scafio.GetArg(c, 0)
  if err != nil {
    questions = append(questions, scafio.EmailQuestion)
  }

  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  projects, err := getProjects(answers.Email)
  if err != nil {
    return err
  }
  for _, project := range projects {
    projectMap := project.(map[string]interface{})
    scafio.PrintProject(projectMap)
  }

  return nil
}
