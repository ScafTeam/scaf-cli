package config

import (
  "fmt"
  "github.com/urfave/cli/v2"
  "github.com/AlecAivazis/survey/v2"
  "scaf/cli/user"
  "scaf/cli/scafio"
)

var (
  categoryQuestion = &survey.Question{
    Name: "Category",
    Prompt: &survey.Select{
      Message: "Please select your config category:",
      Options: Categories,
    },
    Validate: survey.Required,
  }
  fieldQuestion = &survey.Question{
    Name: "Field",
    Prompt: &survey.Input{
      Message: "Please input your config field:",
    },
    Validate: survey.Required,
  }
  valueQuestion = &survey.Question{
    Name: "Value",
    Prompt: &survey.Input{
      Message: "Please input your config value:",
    },
    Validate: survey.Required,
  }
  // asking user to set config to true or false
  boolQuestion = &survey.Question{
    Name: "Value",
    Prompt: &survey.Select{
      Message: "Please select your config value:",
      Options: []string{"true", "false"},
    },
    Validate: survey.Required,
  }
)

func SetConfigAction(c *cli.Context) error {
  // get config input
  var err error
  questions := []*survey.Question{}
  answers := struct {
    Category string
    Field    string
    Value    string
  }{}

  answers.Category, err = scafio.GetArg(c, 0)
  if err != nil {
    questions = append(questions, categoryQuestion)
  }
  answers.Field, err = scafio.GetArg(c, 1)
  if err != nil {
    questions = append(questions, fieldQuestion)
  }
  answers.Value, err = scafio.GetArg(c, 2)
  if err != nil {
    questions = append(questions, valueQuestion)
  }
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }

  // set config
  var message string
  switch answers.Category {
  case User:
    message, err = user.UpdateUser(map[string]interface{}{
      answers.Field: answers.Value,
    })
  // case Project:
  //   message, err = project.UpdateProject(map[string]interface{}{
  //     answers.Field: answers.Value,
  //   })
  default:
    err = fmt.Errorf("Invalid category: %s", answers.Category)
  }
  if err != nil {
    return err
  }

  fmt.Println(message)
  return nil
}

func GetConfigAction(c *cli.Context) error {
  return nil
  // // get config input
  // var err error
  // questions := []*survey.Question{}
  // answers := struct {
  //   Category string
  //   Field    string
  // }{}

  // answers.Category, err = scafio.GetArg(c, 0)
  // if err != nil {
  //   questions = append(questions, categoryQuestion)
  // }
  // answers.Field, err = scafio.GetArg(c, 1)
  // if err != nil {
  //   questions = append(questions, fieldQuestion)
  // }
  // err = survey.Ask(questions, &answers)
  // if err != nil {
  //   return err
  // }

  // // get config
  // var value interface{}
  // switch answers.Category {
  // case User:
  //   value, err = user.GetUser(answers.Field)
  // case Project:
  //   value, err = project.GetProject(answers.Field)
  // default:
  //   err = fmt.Errorf("Invalid category: %s", answers.Category)
  // }
  // if err != nil {
  //   return err
  // }

  // fmt.Printf("Config %s.%s is %v\n", answers.Category, answers.Field, value)
  // return nil
}