package auth

import (
  "log"
  "fmt"
  "sort"
  "github.com/urfave/cli/v2"
  "github.com/AlecAivazis/survey/v2"

  "scaf/cli/scafio"
)

func SignInAction(c *cli.Context) error {
  if c.Bool("forget-password") {
    return ForgetPasswordAction(c)
  } else {
    var err error
    questions := []*survey.Question{}
    answers := struct {
      Email string
      Password string
    }{}

    answers.Email, err = scafio.GetArg(c, 0)
    if err != nil {
      questions = append(questions, scafio.EmailQuestion)
    }
    questions = append(questions, scafio.PasswordQuestion)

    err = survey.Ask(questions, &answers)
    if err != nil {
      return err
    }

    message, err := signIn(answers.Email, answers.Password)
    if err != nil {
      return err
    }

    fmt.Println(message)
    return nil
  }
}

func ForgetPasswordAction(c *cli.Context) error {
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

  message, err := forgetPassword(answers.Email)
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
    err = survey.Ask([]*survey.Question{scafio.EmailQuestion}, &answers)
    if err != nil {
      return err
    }
  }
  questions = append(questions, scafio.PasswordQuestion)
  questions = append(questions, scafio.PasswordConfirmQuestion)

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

  message, err := signUp(answers.Email, answers.Password)
  if err != nil {
    return err
  }

  fmt.Println(message)
  return nil
}

func SignOutAction(c *cli.Context) error {
  message, err := signOut()
  if err != nil {
    return err
  }

  fmt.Println(message)
  return nil
}

func WhoamiAction(c *cli.Context) error {
  message, err := whoami()
  if err != nil {
    log.Println(err)
  }

  fmt.Println(message)
  return nil
}

func GetUserAction(c *cli.Context) error {
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
  userData, err := getUser(answers.Email)
  if err != nil {
    return err
  }

  keys := make([]string, 0, len(userData))
  for k := range userData {
    keys = append(keys, k)
  }
  sort.Strings(keys)
  for _, key := range keys {
    fmt.Printf("%s: %s\n", key, userData[key])
  }

  return nil
}
