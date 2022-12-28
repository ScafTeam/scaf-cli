package auth
// TODO: rename this file to appropriate name

import (
  "log"
  "net/http"
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
  emailCookie := http.Cookie{
    Name:  "email",
    Value: email,
  }
  err = scafreq.SaveCookies([]*http.Cookie{&emailCookie})
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
  err := scafreq.DeleteCookies([]string{"email", "jwt"})
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
  defer resp.Body.Close()
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
  // TODO: check if signin is expired
  emailCookie, err := scafreq.LoadCookie("email")
  if err != nil {
    return "You are not signed in", nil
  }
  return "You are signed in as " + emailCookie.Value, nil
}

