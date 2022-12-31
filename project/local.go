package project

import (
  "errors"
  "os"
  "log"
  "encoding/json"
  "scaf/cli/scafio"
  "scaf/cli/scafreq"
)

// guarantee that the project has author and name
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
  // check if project.json is valid
  if projectBody["author"] == nil {
    return nil, errors.New("Project author not found")
  }
  if projectBody["name"] == nil {
    return nil, errors.New("Project name not found")
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
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
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

func AddRepo(repoName, repoUrl string) (string, error) {
  log.Println("addRepo:", repoName, repoUrl)

  // get local project
  localProject, err := GetLocalProject()
  if err != nil {
    return "", err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // add repo
  addRepoReqBody := map[string]interface{}{
    "name": repoName,
    "url": repoUrl,
  }
  addRepoReqBodyString, err := json.Marshal(addRepoReqBody)
  if err != nil {
    return "", err
  }
  req, err := scafreq.NewRequest(
    "POST",
    "/user/" + projectAuthor + "/project/" + projectName + "/repo",
    addRepoReqBodyString,
  )
  if err != nil {
    return "", err
  }
  resp, err := scafreq.DoRequest(req)
  if err != nil {
    return "", err
  }
  defer resp.Body.Close()
  if resp.StatusCode != 201 {
    return "", errors.New("Failed to add repo")
  }
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }
  // update local project
  _, err = PullProjectFromRemote()
  if err != nil {
    return "", err
  }
  return body["message"].(string), nil
}

func UpdateRepo(repoId, repoName, repoUrl string) (string, error) {
  log.Println("updateRepo:", repoId, repoName, repoUrl)
  // get local project
  localProject, err := GetLocalProject()
  if err != nil {
    return "", err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // update repo
  updateRepoReqBody := map[string]interface{}{
    "id": repoId,
    "name": repoName,
    "url": repoUrl,
  }
  updateRepoReqBodyString, err := json.Marshal(updateRepoReqBody)
  if err != nil {
    return "", err
  }
  req, err := scafreq.NewRequest(
    "PUT",
    "/user/" + projectAuthor + "/project/" + projectName + "/repo",
    updateRepoReqBodyString,
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
    return "", errors.New("Failed to update repo")
  }
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }
  // update local project
  _, err = PullProjectFromRemote()
  if err != nil {
    return "", err
  }
  return body["message"].(string), nil
}

func DeleteRepo(repoId string) (string, error) {
  log.Println("deleteRepo:", repoId)
  // get local project
  localProject, err := GetLocalProject()
  if err != nil {
    return "", err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // delete repo
  deleteRepoReqBody := map[string]interface{}{
    "id": repoId,
  }
  deleteRepoReqBodyString, err := json.Marshal(deleteRepoReqBody)
  if err != nil {
    return "", err
  }
  req, err := scafreq.NewRequest(
    "DELETE",
    "/user/" + projectAuthor + "/project/" + projectName + "/repo",
    deleteRepoReqBodyString,
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
    return "", errors.New("Failed to delete repo")
  }
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }
  // update local project
  _, err = PullProjectFromRemote()
  if err != nil {
    return "", err
  }
  return body["message"].(string), nil
}
