package action

import (
  "fmt"
  "github.com/urfave/cli/v2"
  "github.com/AlecAivazis/survey/v2"
  "scaf/cli/user"
  "scaf/cli/config"
  "scaf/cli/scafio"
  "scaf/cli/scafreq"
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
    questions = append(questions, configCategoryQuestion)
  }
  answers.Field, err = scafio.GetArg(c, 1)
  if err != nil {
    questions = append(questions, configFieldQuestion)
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
  case config.User:
    message, err = user.UpdateUser(map[string]interface{}{
      answers.Field: answers.Value,
    })
  // case config.Project:
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
  // get config input
  var err error
  questions := []*survey.Question{}
  answers := struct {
    Category string
    Field    string
  }{}
  answers.Category, err = scafio.GetArg(c, 0)
  if err != nil {
    questions = append(questions, configCategoryQuestion)
  }
  answers.Field, err = scafio.GetArg(c, 1)
  if err != nil {
    questions = append(questions, configFieldQuestion)
  }
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  // get config
  var value interface{}
  switch answers.Category {
  case config.User:
    email, err := scafreq.LoadCookieValue("email")
    if err != nil {
      return err
    }
    userData, err := user.GetUser(email)
    if err != nil {
      return err
    }
    var ok bool
    value, ok = userData[answers.Field]
    if !ok {
      return fmt.Errorf("Invalid field: %s", answers.Field)
    }
  // case config.Project:
  //   value, err = project.GetProject(answers.Field)
  default:
    err = fmt.Errorf("Invalid category: %s", answers.Category)
  }
  if err != nil {
    return err
  }
  fmt.Printf("%s.%s = %v\n", answers.Category, answers.Field, value)
  return nil
}

func ChangePasswordAction(c *cli.Context) error {
  // get config input
  var err error
  questions := []*survey.Question{
    oldPasswordQuestion,
    newPasswordQuestion,
  }
  answers := struct {
    OldPassword string
    NewPassword string
  }{}
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  // change password
  message, err := user.ChangePassword(answers.OldPassword, answers.NewPassword)
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}
