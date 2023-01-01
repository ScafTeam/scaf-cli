package project

import (
  "encoding/json"
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

func AddWorkflow(projectAuthor, projectName, workflowName string) (string, error) {
  log.Println("addWorkflow:", projectAuthor, projectName, workflowName)
  // add workflow
  addWorkflowReq := map[string]interface{}{
    "name": workflowName,
  }
  addWorkflowReqJson, err := json.Marshal(addWorkflowReq)
  req, err := scafreq.NewRequest(
    "POST",
    "/user/" + projectAuthor + "/project/" + projectName + "/kanban",
    addWorkflowReqJson,
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
