package auth
// TODO: rename this file to appropriate name

import (
  "log"
  "encoding/json"
  "scaf/cli/scafio"
  "scaf/cli/scafreq"
)

func signIn(email, password string) (string, error) {
  log.Println("signIn:", email, password) // TODO: remove logging password on production

  signInRequest := map[string]string{
    "email":    email,
    "password": password,
  }
  signInRequestJSON, err := json.Marshal(signInRequest)
  if err != nil {
    return "", err
  }
  req, err := scafreq.NewRequest("POST", "/signin", signInRequestJSON)
  if err != nil {
    return "", err
  }
  resp, err := scafreq.DoRequest(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }

  return body["message"].(string), nil
}

func signOut() (string, error) {
  err := scafreq.DeleteCookies()
  if err != nil {
    return "", err
  }

  return "Signed out success", nil
}

func forgetPassword(email string) (string, error) {
  log.Println("forgetPassword:", email)

  forgetPasswordRequest := map[string]string{
    "email": email,
  }
  forgetPasswordRequestJSON, err := json.Marshal(forgetPasswordRequest)
  if err != nil {
    return "", err
  }
  req, err := scafreq.NewRequest("POST", "/forgot", forgetPasswordRequestJSON)
  if err != nil {
    return "", err
  }
  resp, err := scafreq.DoRequest(req)
  if err != nil {
    return "", err
  }
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }

  return body["message"].(string), nil
}

func signUp(email, password string) (string, error) {
  log.Println("signup", email, password) // TODO: remove logging password on production

  signUpRequest := map[string]string{
    "email":    email,
    "password": password,
  }
  signUpRequestJSON, err := json.Marshal(signUpRequest)
  if err != nil {
    return "", err
  }
  req, err := scafreq.NewRequest("POST", "/signup", signUpRequestJSON)
  if err != nil {
    return "", err
  }
  resp, err := scafreq.DoRequest(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }
  return body["message"].(string), nil
}

func whoami() (string, error) {
  // TODO: fix api to /user/:email
  req, err := scafreq.NewRequest("GET", "/hello", nil)
  if err != nil {
    return "", err
  }
  resp, err := scafreq.DoRequest(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()
  body, err := scafio.ReadBody(resp)

  if val, ok := body["uesrEmail_claims"]; ok { // TODO: fix backend typo
    return "You are logged in as " + val.(string), nil
  } else {
    return "You are not logged in", nil
  }
}
