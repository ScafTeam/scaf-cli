package scafio

import (
  "fmt"
  "net/http"
  "encoding/json"
  "io/ioutil"
  "golang.org/x/term"
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
