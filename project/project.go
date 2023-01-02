package project

import (
  "errors"
  "log"
  "os"
  "encoding/json"
  "scaf/cli/scafio"
  "scaf/cli/scafreq"
)

func GetProjects(email string) ([]interface{}, error) {
  log.Println("getProjects:", email)
  req, err := scafreq.NewRequest("GET", "/user/" + email + "/project", nil)
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
  // bodyIndent, err := json.MarshalIndent(body, "", "  ")
  // log.Println(string(bodyIndent))
  return body["projects"].([]interface{}), nil
}

func CreateProject(name string, devMode string, devTools []string) (string, error) {
  log.Println("createProject:", name, devMode, devTools)
  // check if project name folder exists
  if _, err := os.Stat(name); !os.IsNotExist(err) {
    return "", errors.New("Project folder already exists")
  }
  // create project to remote
  createProjectRequest := map[string]interface{}{
    "name": name,
    "devMode": devMode,
    "devTools": devTools,
  }
  createProjectRequestJSON, err := json.Marshal(createProjectRequest)
  if err != nil {
    return "", err
  }
  email, err := scafreq.LoadCookieValue("email")
  if err != nil {
    return "", err
  }
  req, err := scafreq.NewRequest(
    "POST",
    "/user/" + email + "/project",
    createProjectRequestJSON,
  )
  if err != nil {
    return "", err
  }
  resp, err := scafreq.DoRequest(req)
  if err != nil {
    return "", err
  }
  if resp.StatusCode != 201 {
    return "", errors.New("Failed to create project")
  }
  defer resp.Body.Close()
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }
  // clone project
  _, err = CloneProjectIntoLocal(email, name)
  if err != nil {
    return "", err
  }
  return body["message"].(string), nil
}

func DeleteProject(name string) (string, error) {
  log.Println("deleteProject:", name)
  // delete project from remote
  email, err := scafreq.LoadCookieValue("email")
  if err != nil {
    return "", err
  }
  req, err := scafreq.NewRequest(
    "DELETE",
    "/user/" + email + "/project/" + name,
    nil,
  )
  if err != nil {
    return "", err
  }
  resp, err := scafreq.DoRequest(req)
  if err != nil {
    return "", err
  }
  if resp.StatusCode != 200 {
    return "", errors.New("Failed to delete project")
  }
  defer resp.Body.Close()
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }
  return body["message"].(string), nil
}

func CloneProjectIntoLocal(email string, name string) (string, error) {
  log.Println("cloneProjectIntoLocal:", email, name)
  // check if project name folder exists
  if _, err := os.Stat(name); !os.IsNotExist(err) {
    return "", errors.New("Project folder already exists")
  }
  //check if user have permission to clone project
  req, err := scafreq.NewRequest(
    "GET",
    "/user/" + email + "/project/" + name,
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
  // clone project
  os.MkdirAll(name + "/.scaf", 0777)
  os.Chdir(name + "/.scaf")
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
