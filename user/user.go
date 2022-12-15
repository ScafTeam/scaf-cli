package user

import (
  "fmt"
  "log"
  "os"
  "net/http"
  "encoding/json"
  "bytes"

  "github.com/urfave/cli/v2"
  "golang.org/x/term"
)

func SignInAction(c *cli.Context) error {
  var err error = nil
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
  } else {
    // if email provided as argument, use it
    if c.NArg() > 0 {
      email = c.Args().Get(0)
    } else {
      email = inputEmail()
    }
    password = inputPassword()
    resp, err = signIn(email, password)
  }
  log.Println(resp.StatusCode)
  return err
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

func signIn(email, password string) (*http.Response, error) {
  log.Println("signIn:", email, password) // TODO: remove logging password on production

  signInRequest := SignInRequest{
    Email:    email,
    Password: password,
  }
  signInRequestJSON, err := json.Marshal(signInRequest)
  if err != nil {
    return nil, err
  }

  resp, err := http.Post(
    os.Getenv("SCAF_BACKEND_URL") + "/signin",
    "application/json",
    bytes.NewBuffer(signInRequestJSON),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  return resp, nil
}

func forgetPassword(email string) (*http.Response, error) {
  log.Println("forgetPassword:", email)

  forgetPasswordRequest := ForgetPasswordRequest{
    Email: email,
  }
  forgetPasswordRequestJSON, err := json.Marshal(forgetPasswordRequest)
  if err != nil {
    return nil, err
  }

  resp, err := http.Post(
    os.Getenv("SCAF_BACKEND_URL") + "/forget",
    "application/json",
    bytes.NewBuffer(forgetPasswordRequestJSON),
  )
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()

  return resp, nil
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
