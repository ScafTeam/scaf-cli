package project

import (
  "os/exec"
  "errors"
  "encoding/json"
  "scaf/cli/scafio"
  "scaf/cli/scafreq"
)

func AddRepo(repoName, repoUrl string) (string, error) {

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

func PullRepo(repoId string) (string, error) {
  // get local project
  localProject, err := GetLocalProject()
  if err != nil {
    return "", err
  }
  repoList := localProject["repos"].([]interface{})
  var repo map[string]interface{}
  for _, repoItem := range repoList {
    repoItemMap, ok := repoItem.(map[string]interface{})
    if !ok {
      continue
    }
    if repoItemMap["id"].(string) == repoId {
      repo = repoItemMap
      break
    }
  } 
  out, err := exec.Command("git", "clone", repo["url"].(string)).CombinedOutput()
  if err != nil {
    return "", err
  }
  return string(out), nil
}
