package project

import (
  "errors"
  "os"
  "log"
  "encoding/json"
  "scaf/cli/scafio"
  "scaf/cli/scafreq"
)

func GetLocalProject() (map[string]interface{}, error) {
  log.Println("getLocalProject")

  // check if project name folder exists
  if _, err := os.Stat(".scaf"); os.IsNotExist(err) {
    return nil, errors.New("Project folder not found")
  }
  // check if project.json exists
  if _, err := os.Stat(".scaf/project.json"); os.IsNotExist(err) {
    return nil, errors.New("Project not found")
  }
  // read project.json
  projectBodyString, err := os.ReadFile(".scaf/project.json")
  if err != nil {
    return nil, err
  }
  var projectBody map[string]interface{}
  err = json.Unmarshal(projectBodyString, &projectBody)
  if err != nil {
    return nil, err
  }
  return projectBody, nil
}

func PullProjectFromRemote() (string, error) {
  log.Println("pullProjectFromRemote")

  // get local project
  localProject, err := GetLocalProject()
  if err != nil {
    return "", err
  }
  projectAuthor, ok := localProject["author"].(string)
  if !ok {
    return "", errors.New("Project author not found")
  }
  projectName, ok := localProject["name"].(string)
  if !ok {
    return "", errors.New("Project name not found")
  }
  // check if user have permission to pull project
  req, err := scafreq.NewRequest(
    "GET",
    "/user/" + projectAuthor + "/project/" + projectName,
    nil,
  )
  if err != nil {
    return "", err
  }
  resp, err := scafreq.DoRequest(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()
  if resp.StatusCode != 200 {
    return "", errors.New("Project not found")
  }
  // pull project
  os.Chdir(".scaf")
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }
  projectBodyString, err := json.Marshal(body["project"])
  if err != nil {
    return "", err
  }
  err = os.WriteFile("project.json", projectBodyString, 0777)
  if err != nil {
    return "", err
  }
  return body["message"].(string), nil
}
