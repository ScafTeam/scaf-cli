package user

import (
  "errors"
  "log"
  "scaf/cli/scafio"
  "scaf/cli/scafreq"
)

func getUser(email string) (map[string]interface{}, error) {
  log.Println("getUser:", email)

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
  log.Println("getUser body:", body)
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
