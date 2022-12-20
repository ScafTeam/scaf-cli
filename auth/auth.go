package auth
// TODO: rename this file to appropriate name

import (
  "log"
  "os"
  "net/http"
  "encoding/json"
  "bytes"
  "scaf/cli/scafio"
)

func signIn(email, password string) (string, error) {
  log.Println("signIn:", email, password) // TODO: remove logging password on production
  var err error

  signInRequest := AuthRequest{
    Email:    email,
    Password: password,
  }
  signInRequestJSON, err := json.Marshal(signInRequest)
  if err != nil {
    return "", err
  }

  resp, err := http.Post(
    os.Getenv("SCAF_BACKEND_URL") + "/signin",
    "application/json",
    bytes.NewBuffer(signInRequestJSON),
  )
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  err = saveCookies(resp)
  if err != nil {
    return "", err
  }

  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }

  return body["message"].(string), nil
}

func signOut() (string, error) {
  err := deleteCookies()
  if err != nil {
    return "", err
  }

  return "Signed out success", nil
}

func forgetPassword(email string) (string, error) {
  log.Println("forgetPassword:", email)

  forgetPasswordRequest := ForgetPasswordRequest{
    Email: email,
  }
  forgetPasswordRequestJSON, err := json.Marshal(forgetPasswordRequest)
  if err != nil {
    return "", err
  }

  resp, err := http.Post(
    os.Getenv("SCAF_BACKEND_URL") + "/forgot",
    "application/json",
    bytes.NewBuffer(forgetPasswordRequestJSON),
  )
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  err = saveCookies(resp)
  if err != nil {
    return "", err
  }
  log.Println("forgetPassword: saved cookies")

  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }

  return body["message"].(string), nil
}

func signUp(email, password string) (string, error) {
  log.Println("signup", email, password) // TODO: remove logging password on production

  signUpRequest := AuthRequest{
    Email:    email,
    Password: password,
  }
  signUpRequestJSON, err := json.Marshal(signUpRequest)
  if err != nil {
    return "", err
  }

  resp, err := http.Post(
    os.Getenv("SCAF_BACKEND_URL") + "/signup",
    "application/json",
    bytes.NewBuffer(signUpRequestJSON),
  )
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()

  err = saveCookies(resp)
  if err != nil {
    return "", err
  }

  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }

  return body["message"].(string), nil
}

func whoami() (string, error) {
  jwt, err := readJWT()
  if err != nil {
    jwt = ""
  }

  client := &http.Client{}
  req, err := http.NewRequest("GET", os.Getenv("SCAF_BACKEND_URL") + "/hello", nil)
  if err != nil {
    return "", err
  }

  req.Header.Add("Authorization", "Bearer " + jwt)
  resp, err := client.Do(req)
  if err != nil {
    return "", err
  }

  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }

  if val, ok := body["uesrEmail_claims"]; ok { // TODO: fix backend typo
    return "You are logged in as " + val.(string), nil
  } else {
    return "You are not logged in", nil
  }
}
