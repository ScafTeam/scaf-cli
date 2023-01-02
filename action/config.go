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
  answers := struct {
    Category  string
    Field     string
    Value     string
    ValueList []string
  }{}
  answers.Category, err = scafio.GetArg(c, 0)
  if err != nil {
    if prompt, err := config.GetCategoryPrompt(); err != nil {
      return err
    } else {
      err = survey.AskOne(prompt, &answers.Category)
      if err != nil {
        return err
      }
    }
  }
  answers.Field, err = scafio.GetArg(c, 1)
  if err != nil {
    if prompt, err := config.GetFieldPrompt(answers.Category); err != nil {
      return err
    } else {
      err = survey.AskOne(prompt, &answers.Field)
      if err != nil {
        return err
      }
    }
  }
  answers.Value, err = scafio.GetArg(c, 2)
  var message string
  if err != nil {
    if prompt, err := config.GetValuePrompt(answers.Category, answers.Field); err != nil {
      return err
    } else {
      switch v := prompt.(type) {
      case *survey.MultiSelect:
        err = survey.AskOne(v, &answers.ValueList)
        if err != nil {
          return err
        }
        message, err = config.SetConfig(answers.Category, answers.Field, answers.ValueList)
        if err != nil {
          return err
        }
      default:
        err = survey.AskOne(v, &answers.Value)
        if err != nil {
          return err
        }
        message, err = config.SetConfig(answers.Category, answers.Field, answers.Value)
        if err != nil {
          return err
        }
      }
    }
  }
  // set config
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
