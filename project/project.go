package project

import (
  "log"
  "scaf/cli/scafio"
  "scaf/cli/scafreq"
)

func getProjects(email string) ([]interface{}, error) {
  log.Println("getProjects:", email)

  req, err := scafreq.NewRequest("GET", "/user/" + email + "/project", nil)
  if err != nil {
    return nil, err
  }
  resp, err := scafreq.DoRequest(req)
  if err != nil {
    return nil, err
  }
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return nil, err
  }
  // bodyIndent, err := json.MarshalIndent(body, "", "  ")
  // log.Println(string(bodyIndent))

  return body["projects"].([]interface{}), nil
}
