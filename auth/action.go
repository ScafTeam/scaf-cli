package auth

import (
  "log"
  "net/http"
  "github.com/urfave/cli/v2"

  "scaf/cli/input"
)

func SignInAction(c *cli.Context) error {
  if c.Bool("forget-password") {
    return ForgetPasswordAction(c)
  } else {
    var err error
    var resp *http.Response = nil
    var email, password string
    email, err = getEmail(c)
    if err != nil {
      return err
    }

    password, err = input.InputPassword()
    if err != nil {
      return err
    }

    resp, err = signIn(email, password)
    if err != nil {
      return err
    }

    log.Println(resp.StatusCode)
    return nil
  }
}

func ForgetPasswordAction(c *cli.Context) error {
  var err error
  var resp *http.Response = nil
  var email string

  email, err = getEmail(c)
  if err != nil {
    return err
  }

  resp, err = forgetPassword(email)
  if err != nil {
    return err
  }

  log.Println(resp.StatusCode)
  return nil
}


func SignUpAction(c *cli.Context) error {
  var err error
  var resp *http.Response = nil
  var email, password string

  email, err = getEmail(c)
  if err != nil {
    return err
  }

  password, err = input.InputComfirmedPassword(3)
  if err != nil {
    return err
  }

  resp, err = signUp(email, password)
  if err != nil {
    return err
  }

  log.Println(resp.StatusCode)
  return nil
}

// get email from first argument, or prompt user to input
func getEmail(c *cli.Context) (string, error) {
  var email string
  var err error
  if c.NArg() > 0 {
    email = c.Args().Get(0)
  } else {
    email, err = input.InputEmail()
    if err != nil {
      return "", err
    }
  }
  return email, nil
}

