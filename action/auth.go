package action

import (
  "fmt"
  "github.com/urfave/cli/v2"
  "github.com/AlecAivazis/survey/v2"
  "scaf/cli/scafio"
  "scaf/cli/auth"
)

func SignInAction(c *cli.Context) error {
  if c.Bool("forget-password") {
    return ForgetPasswordAction(c)
  }
  var err error
  questions := []*survey.Question{}
  answers := struct {
    Email string
    Password string
  }{}
  answers.Email, err = scafio.GetArg(c, 0)
  if err != nil {
    questions = append(questions, emailQuestion)
  }
  questions = append(questions, passwordQuestion)
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  message, err := auth.SignIn(answers.Email, answers.Password)
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func ForgetPasswordAction(c *cli.Context) error {
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
  message, err := auth.ForgetPassword(answers.Email)
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func SignUpAction(c *cli.Context) error {
  var err error
  questions := []*survey.Question{}
  answers := struct {
    Email string
    Password string
    PasswordConfirm string
  }{}
  answers.Email, err = scafio.GetArg(c, 0)
  if err != nil {
    err = survey.Ask([]*survey.Question{emailQuestion}, &answers)
    if err != nil {
      return err
    }
  }
  questions = append(questions, passwordQuestion)
  questions = append(questions, passwordConfirmQuestion)
  for i := 0; i < 3; i++ {
    if i > 0 {
      fmt.Println("Passwords do not match. Please try again.")
    }
    err = survey.Ask(questions, &answers)
    if err != nil {
      return err
    }
    if answers.Password == answers.PasswordConfirm {
      break
    }
  }
  if answers.Password != answers.PasswordConfirm {
    return fmt.Errorf("Passwords do not match.")
  }
  message, err := auth.SignUp(answers.Email, answers.Password)
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func SignOutAction(c *cli.Context) error {
  message, err := auth.SignOut()
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func WhoamiAction(c *cli.Context) error {
  message, err := auth.Whoami()
  if err != nil {
  }
  fmt.Println(message)
  return nil
}
