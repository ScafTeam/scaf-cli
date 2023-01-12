package user

import (
  "encoding/json"
  "errors"
  "scaf/cli/scafio"
  "scaf/cli/scafreq"
)

func GetUser(email string) (map[string]interface{}, error) {
  req, err := scafreq.NewRequest("GET", "/user/" + email, nil)
  if err != nil {
    return nil, err
  }
  resp, err := scafreq.DoRequest(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return nil, err
  }
  userData, ok := body["data"].(map[string]interface{})
  if !ok {
    message, ok := body["message"].(string)
    if !ok {
      return nil, errors.New("Invalid response from server")
    }
    return nil, errors.New(message)
  }
  return userData, nil
}

func UpdateUser(data map[string]interface{}) (string, error) {
  updateUserRequestJSON, err := json.Marshal(data)
  if err != nil {
    return "", err
  }
  email, err := scafreq.LoadCookieValue("email")
  if err != nil {
    return "", err
  }
  req, err := scafreq.NewRequest(
    "PUT",
    "/user/" + email,
    updateUserRequestJSON)
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

func ChangePassword(oldPassword string, newPassword string) (string, error) {

  changePasswordRequest := map[string]interface{}{
    "oldPassword": oldPassword,
    "newPassword": newPassword,
  }
  changePasswordRequestJSON, err := json.Marshal(changePasswordRequest)
  if err != nil {
    return "", err
  }
  email, err := scafreq.LoadCookieValue("email")
  if err != nil {
    return "", err
  }
  req, err := scafreq.NewRequest(
    "PUT",
    "/user/" + email + "/reset",
    changePasswordRequestJSON,
  )
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
