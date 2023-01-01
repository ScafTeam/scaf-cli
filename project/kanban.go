package project

import (
  "log"
  "scaf/cli/scafio"
  "scaf/cli/scafreq"
)

func GetWorkflows(projectAuthor string, projectName string) ([]interface{}, error) {
  log.Println("getWorkflows:", projectAuthor, projectName)

  req, err := scafreq.NewRequest(
    "GET",
    "/user/" + projectAuthor + "/project/" + projectName + "/kanban",
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
  body, err := scafio.ReadBody(resp)
  if err != nil {
    return nil, err
  }
  kanban := body["kanban"].(map[string]interface{})
  return kanban["workflows"].([]interface{}), nil
}

// func AddWorkflow(workflowName string) (string, error) {
//   log.Println("addWorkflow:", workflowName)
//   // get local project
//   localProject, err := GetLocalProject()
//   if err != nil {
//     return "", err
//   }
//   projectAuthor := localProject["author"].(string)
//   projectName := localProject["name"].(string)
