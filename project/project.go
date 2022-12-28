package project

import (
  "log"
  "encoding/json"
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

func createProject(name string, devMode string, devTools []string) (string, error) {
  log.Println("createProject:", name, devMode, devTools)

  createProjectRequest := map[string]interface{}{
    "name": name,
    "devMode": devMode,
    "devTools": devTools,
  }
  createProjectRequestJSON, err := json.Marshal(createProjectRequest)
  if err != nil {
    return "", err
  }
  emailCookie, err := scafreq.LoadCookie("email")
  req, err := scafreq.NewRequest(
    "POST",
    "/user/" + emailCookie.Value + "/project",
    createProjectRequestJSON,
  )
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
