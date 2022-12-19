package auth

import (
  "log"
  "net/http"
  "github.com/urfave/cli/v2"

  "scaf/cli/scafio"
)

func SignInAction(c *cli.Context) error {
  if c.Bool("forget-password") {
    return ForgetPasswordAction(c)
  } else {
    var err error
    var resp *http.Response = nil
    var email, password string
    email, err = scafio.GetEmail(c)
    if err != nil {
      return err
    }

    password, err = scafio.InputPassword()
    if err != nil {
      return err
    }

    resp, err = signIn(email, password)
    if err != nil {
      return err
    }
    defer resp.Body.Close()

    log.Println(resp.StatusCode)
    return nil
  }
}

func ForgetPasswordAction(c *cli.Context) error {
  var err error
  var resp *http.Response = nil
  var email string

  email, err = scafio.GetEmail(c)
  if err != nil {
    return err
  }

  resp, err = forgetPassword(email)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  log.Println(resp.StatusCode)
  return nil
}


func SignUpAction(c *cli.Context) error {
  var err error
  var resp *http.Response = nil
  var email, password string

  email, err = scafio.GetEmail(c)
  if err != nil {
    return err
  }

  password, err = scafio.InputComfirmedPassword(3)
  if err != nil {
    return err
  }

  resp, err = signUp(email, password)
  if err != nil {
    return err
  }
  defer resp.Body.Close()

  log.Println(resp.StatusCode)
  return nil
}

func WhoamiAction(c *cli.Context) error {
  var err error
  var resp *http.Response = nil

  resp, err = whoami()
  if err != nil {
    log.Println(err)
  } else {
    defer resp.Body.Close()
  }

  err = scafio.OutputWhoami(resp)
  if err != nil {
    return err
  }
  return nil
}
