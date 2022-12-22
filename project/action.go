package project

import (
  "fmt"
  "github.com/urfave/cli/v2"
  "github.com/AlecAivazis/survey/v2"
  "scaf/cli/scafio"
)

var (
  DevModeQuestion = &survey.Question{
    Name: "DevMode",
    Prompt: &survey.Select{
      Message: "Please select your development mode:",
      Options: DevModes,
    },
    Validate: survey.Required,
  }
  DevToolsQuestion = &survey.Question{
    Name: "DevTools",
    Prompt: &survey.MultiSelect{
      Message: "Please select your dev tools:",
      Options: DevTools,
    },
  }
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
    questions = append(questions, scafio.ProjectNameQuestion)
  }
  questions = append(questions, DevModeQuestion)
  questions = append(questions, DevToolsQuestion)

  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }

  message, err := createProject(answers.ProjectName, answers.DevMode, answers.DevTools)
  if err != nil {
    return err
  }

  fmt.Println(message)
  return nil
}
