package auth

import (
  "log"
  "fmt"
  "github.com/urfave/cli/v2"

  "scaf/cli/scafio"
)

func SignInAction(c *cli.Context) error {
  if c.Bool("forget-password") {
    return ForgetPasswordAction(c)
  } else {
    var err error
    var email, password string
    email, err = scafio.GetEmail(c)
    if err != nil {
      return err
    }

    password, err = scafio.InputPassword()
    if err != nil {
      return err
    }

    message, err := signIn(email, password)
    if err != nil {
      return err
    }

    fmt.Println(message)
    return nil
  }
}

func ForgetPasswordAction(c *cli.Context) error {
  var err error
  var email string

  email, err = scafio.GetEmail(c)
  if err != nil {
    return err
  }

  message, err := forgetPassword(email)
  if err != nil {
    return err
  }

  fmt.Println(message)
  return nil
}

func SignUpAction(c *cli.Context) error {
  var err error
  var email, password string

  email, err = scafio.GetEmail(c)
  if err != nil {
    return err
  }

  password, err = scafio.InputComfirmedPassword(3)
  if err != nil {
    return err
  }

  message, err := signUp(email, password)
  if err != nil {
    return err
  }

  fmt.Println(message)
  return nil
}

func SignOutAction(c *cli.Context) error {
  var err error

  message, err := signOut()
  if err != nil {
    return err
  }

  fmt.Println(message)
  return nil
}

func WhoamiAction(c *cli.Context) error {
  var err error

  message, err := whoami()
  if err != nil {
    log.Println(err)
  }

  fmt.Println(message)
  return nil
}
