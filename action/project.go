package action

import (
  "fmt"
  "github.com/urfave/cli/v2"
  "github.com/AlecAivazis/survey/v2"
  "scaf/cli/scafio"
  "scaf/cli/project"
)

func CloneProjectAction(c *cli.Context) error {
  var err error
  questions := []*survey.Question{}
  answers := struct {
    Email string
    ProjectName string
  }{}
  answers.Email, err = scafio.GetArg(c, 0)
  if err != nil {
    questions = append(questions, scafio.EmailQuestion)
  }
  answers.ProjectName, err = scafio.GetArg(c, 1)
  if err != nil {
    questions = append(questions, scafio.ProjectNameQuestion)
  }
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  message, err := project.CloneProjectIntoLocal(answers.Email, answers.ProjectName)
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}
