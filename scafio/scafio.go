package scafio

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "golang.org/x/term"
  "github.com/urfave/cli/v2"
)

func InputEmail() (string, error) {
  var email string
  fmt.Print("Please enter your email: ")
  fmt.Scanln(&email)
  return email, nil
}

func InputPassword() (string, error) {
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

func InputComfirmedPassword(retry_times int) (string, error) {
  var password string
  var err error
  for i := 0; i < retry_times; i++ {
    password, err = InputPassword()
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

// get email from first argument, or prompt user to input
func GetEmail(c *cli.Context) (string, error) {
  var email string
  var err error
  if c.NArg() > 0 {
    email = c.Args().Get(0)
  } else {
    email, err = InputEmail()
    if err != nil {
      return "", err
    }
  }
  return email, nil
}


func PrintProject(projectMap map[string]interface{}) {
  fmt.Printf("* [%s] %s (%s)\n", projectMap["Id"], projectMap["Name"], projectMap["Author"])
}

// read json format response body and return a map
func ReadBody(resp *http.Response) (map[string]interface{}, error) {
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  bodyMap := make(map[string]interface{})
  err = json.Unmarshal(body, &bodyMap)
  if err != nil {
    bodyMap["message"] = string(body)
  }

  return bodyMap, nil
}
