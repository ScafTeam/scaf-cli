package auth

import (
  "fmt"
  "log"
  "net/http"
  "golang.org/x/term"
  "github.com/urfave/cli/v2"
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

    password, err = inputPassword()
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

  password, err = inputComfirmedPassword(3)
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
    email, err = inputEmail()
    if err != nil {
      return "", err
    }
  }
  return email, nil
}

func inputEmail() (string, error) {
  var email string
  fmt.Print("Please enter your email: ")
  fmt.Scanln(&email)
  return email, nil
}

func inputPassword() (string, error) {
  var password string
  var err error
  fmt.Print("Please enter your password: ")
  bytePassword, err := term.ReadPassword(0)
  if err != nil {
    return "", err
  }
  password = string(bytePassword)
  fmt.Println()
  return password, nil

}

func inputComfirmedPassword(retry_times int) (string, error) {
  var password string
  var err error
  for i := 0; i < retry_times; i++ {
    password, err = inputPassword()
    if err != nil {
      return "", err
    }
    fmt.Print("Please enter your password again: ")
    bytePassword, err := term.ReadPassword(0)
    if err != nil {
      return "", err
    }
    confirmedPassword := string(bytePassword)
    fmt.Println()
    if password == confirmedPassword {
      return password, nil
    }
    fmt.Println("Password not match, please try again")
  }
  return "", fmt.Errorf("password confirmation failed")
}
