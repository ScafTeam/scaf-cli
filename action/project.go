package action

import (
  "fmt"
  "github.com/urfave/cli/v2"
  "github.com/AlecAivazis/survey/v2"
  "scaf/cli/scafio"
  "scaf/cli/project"
)

func ListProjectsAction(c *cli.Context) error {
  var err error
  questions := []*survey.Question{}
  answers := struct {
    Email string
  }{}
  answers.Email, err = scafio.GetArg(c, 0)
  if err != nil {
    questions = append(questions, emailQuestion)
  }
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  projects, err := project.GetProjects(answers.Email)
  if err != nil {
    return err
  }
  if len(projects) == 0 {
    fmt.Println("No project found")
  } else {
    for _, project := range projects {
      projectMap := project.(map[string]interface{})
      scafio.PrintProject(projectMap, c.Bool("oneline"))
    }
  }
  return nil
}

func CreateProjectAction(c *cli.Context) error {
  var err error
  questions := []*survey.Question{}
  answers := struct {
    ProjectName string
    DevMode string
    DevTools []string
  }{}
  answers.ProjectName, err = scafio.GetArg(c, 0)
  if err != nil {
    questions = append(questions, projectNameQuestion)
  }
  questions = append(questions, devModeQuestion)
  questions = append(questions, devToolsQuestion)
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  message, err := project.CreateProject(answers.ProjectName, answers.DevMode, answers.DevTools)
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func CloneProjectAction(c *cli.Context) error {
  var err error
  questions := []*survey.Question{}
  answers := struct {
    Email string
    ProjectName string
  }{}
  answers.Email, err = scafio.GetArg(c, 0)
  if err != nil {
    questions = append(questions, emailQuestion)
  }
  answers.ProjectName, err = scafio.GetArg(c, 1)
  if err != nil {
    questions = append(questions, projectNameQuestion)
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
