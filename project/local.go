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

func GetMembers() ([]interface{}, error) {
  log.Println("getMembers")
  // get local project
  localProject, err := GetLocalProject()
  if err != nil {
    return nil, err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // get members
  req, err := scafreq.NewRequest(
    "GET",
    "/user/" + projectAuthor + "/project/" + projectName + "/member",
    nil,
  )
  if err != nil {
    return nil, err
  }
  resp, err := scafreq.DoRequest(req)
  if err != nil {
    return nil, err
  }
  defer resp.Body.Close()
  if resp.StatusCode != 200 {
    return nil, errors.New("Member not found")
  }
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return nil, err
  }
  return body["members"].([]interface{}), nil
}

func AddMember(member string) (string, error) {
  log.Println("AddMember")
  // get local project
  localProject, err := GetLocalProject()
  if err != nil {
    return "", err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // add member
  updateMemberJSON, err := json.Marshal(map[string]string{
    "email": member,
  })
  req, err := scafreq.NewRequest(
    "POST",
    "/user/" + projectAuthor + "/project/" + projectName + "/member",
    updateMemberJSON,
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
    return "", errors.New("cannot add member")
  }
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }
  return body["message"].(string), nil
}

func DeleteMember(member string) (string, error) {
  log.Println("DeleteMember")
  // get local project
  localProject, err := GetLocalProject()
  if err != nil {
    return "", err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // remove member
  updateMemberJSON, err := json.Marshal(map[string]string{
    "email": member,
  })
  req, err := scafreq.NewRequest(
    "DELETE",
    "/user/" + projectAuthor + "/project/" + projectName + "/member",
    updateMemberJSON,
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
    return "", errors.New("Member not found")
  }
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }
  return body["message"].(string), nil
}

func UpdateLocalProject(data map[string]interface{}) (string, error) {
  log.Println("updateLocalProject")
  // get local project
  localProject, err := GetLocalProject()
  if err != nil {
    return "", err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // update local project
  updateProjectJSON, err := json.Marshal(data)
  if err != nil {
    return "", err
  }
  req, err := scafreq.NewRequest(
    "PUT",
    "/user/" + projectAuthor + "/project/" + projectName,
    updateProjectJSON,
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
    return "", errors.New("update project failed")
  }
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return "", err
  }
  return body["message"].(string), nil
}
