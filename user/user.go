package user

import (
  "log"
  "os"
  "net/http"
  "encoding/json"
  "bytes"
)

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

