package auth
// TODO: rename this file to appropriate name

import (
  "log"
  "os"
  "net/http"
  "encoding/json"
  "bytes"
)

func signIn(email, password string) (*http.Response, error) {
  log.Println("signIn:", email, password) // TODO: remove logging password on production
  var err error

  signInRequest := AuthRequest{
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

  err = saveCookies(resp)
  if err != nil {
    return nil, err
  }

  return resp, nil
}

func signOut() error {
  err := deleteCookies()
  if err != nil {
    return err
  }

  return nil
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

  err = saveCookies(resp)
  if err != nil {
    return nil, err
  }

  return resp, nil
}

func signUp(email, password string) (*http.Response, error) {
  log.Println("signup", email, password) // TODO: remove logging password on production

  signUpRequest := AuthRequest{
    Email:    email,
    Password: password,
  }
  signUpRequestJSON, err := json.Marshal(signUpRequest)
  if err != nil {
    return nil, err
  }

  resp, err := http.Post(
    os.Getenv("SCAF_BACKEND_URL") + "/signup",
    "application/json",
    bytes.NewBuffer(signUpRequestJSON),
  )
  if err != nil {
    return nil, err
  }

  err = saveCookies(resp)
  if err != nil {
    return nil, err
  }

  return resp, nil
}

func whoami() (*http.Response, error) {
  jwt, err := readJWT()
  if err != nil {
    return nil, err
  }

  client := &http.Client{}
  req, err := http.NewRequest("GET", os.Getenv("SCAF_BACKEND_URL") + "/hello", nil)
  if err != nil {
    return nil, err
  }

  req.Header.Add("Authorization", "Bearer " + jwt)
  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }

  return resp, nil
}
