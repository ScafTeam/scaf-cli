package auth

import (
  "fmt"
  "log"
  "net/http"
  "golang.org/x/term"
  "github.com/urfave/cli/v2"
)

func SignInAction(c *cli.Context) error {
  var err error
  var resp *http.Response = nil
  var email, password string

  if c.Bool("forget-password") {
    // if email provided as argument, use it
    if c.NArg() > 0 {
      email = c.Args().Get(0)
    } else {
      email = inputEmail()
    }
    resp, err = forgetPassword(email)
    if err != nil {
      return err
    }
  } else {
    // if email provided as argument, use it
    if c.NArg() > 0 {
      email = c.Args().Get(0)
    } else {
      email = inputEmail()
    }
    password = inputPassword()
    resp, err = signIn(email, password)
    if err != nil {
      return err
    }
  }
  log.Println(resp.StatusCode)
  return nil
}

func RegisterAction(c *cli.Context) error {
  // get user email
  var email string
  if c.NArg() > 0 {
    email = c.Args().Get(0)
  } else {
    fmt.Print("Please enter your email: ")
    fmt.Scanln(&email)
  }

  for i := 0; i < 3; i++ {
    // get user password
    var password string
    fmt.Print("Please enter your password: ")
    bytePassword, err := term.ReadPassword(0)
    if err != nil {
      log.Fatal(err)
    }
    password = string(bytePassword)
    fmt.Println()

    // get user password confirmation
    var passwordConfirmation string
    fmt.Print("Please confirm your password: ")
    bytePasswordConfirmation, err := term.ReadPassword(0)
    if err != nil {
      log.Fatal(err)
    }
    passwordConfirmation = string(bytePasswordConfirmation)
    fmt.Println()

    if password == passwordConfirmation {
      log.Println("register:", email, password)
      return nil
    } else {
      fmt.Println("Confirmation failed, please try again")
    }
  }

  log.Println("register failed:", email, "failed")
  return nil
}

func inputEmail() string {
  var email string
  fmt.Print("Please enter your email: ")
  fmt.Scanln(&email)
  return email
}

func inputPassword() string {
  var password string
  fmt.Print("Please enter your password: ")
  bytePassword, err := term.ReadPassword(0)
  if err != nil {
    log.Fatal(err)
  }
  password = string(bytePassword)
  fmt.Println()
  return password
}

