package user

import (
  "fmt"
  "log"
  "github.com/urfave/cli/v2"
  "golang.org/x/term"
)

func Login(c *cli.Context) error {
  var err error
  if c.Bool("forget-password") {
    err = forgetPassword(c)
  } else {
    err = login(c)
  }
  return err
}

func login(c *cli.Context) error {
  // get user email
  var email string
  if c.NArg() > 0 {
    email = c.Args().Get(0)
  } else {
    fmt.Print("Please enter your email: ")
    fmt.Scanln(&email)
  }

  // get user password
  var password string
  fmt.Print("Please enter your password: ")
  bytePassword, err := term.ReadPassword(0)
  if err != nil {
    log.Fatal(err)
  }
  password = string(bytePassword)
  fmt.Println()

  log.Println("login:", email, password)
  return nil
}

func forgetPassword(c *cli.Context) error {
  // get user email
  var email string
  if c.NArg() > 0 {
    email = c.Args().Get(0)
  } else {
    fmt.Print("Please enter your email: ")
    fmt.Scanln(&email)
  }

  log.Println("forget password:", email)
  return nil
}

func Register(c *cli.Context) error {
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
