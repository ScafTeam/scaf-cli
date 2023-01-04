package project

import (
  "errors"
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

func UpdateWorkflow(projectAuthor, projectName, workflowID, workflowName string) (string, error) {
  log.Println("updateWorkflow:", projectAuthor, projectName, workflowName)
  // update workflow
  updateWorkflowReq := map[string]interface{}{
    "id": workflowID,
    "name": workflowName,
  }
  updateWorkflowReqJson, err := json.Marshal(updateWorkflowReq)
  req, err := scafreq.NewRequest(
    "PUT",
    "/user/" + projectAuthor + "/project/" + projectName + "/kanban",
    updateWorkflowReqJson,
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

func DeleteWorkflow(projectAuthor, projectName, workflowID string) (string, error) {
  log.Println("deleteWorkflow:", projectAuthor, projectName, workflowID)
  // delete workflow
  deleteWorkflowReq := map[string]interface{}{
    "id": workflowID,
  }
  deleteWorkflowReqJson, err := json.Marshal(deleteWorkflowReq)
  req, err := scafreq.NewRequest(
    "DELETE",
    "/user/" + projectAuthor + "/project/" + projectName + "/kanban",
    deleteWorkflowReqJson,
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

func GetTasks(projectAuthor, projectName, workflowID string) ([]interface{}, error) {
  log.Println("getTasks:", projectAuthor, projectName, workflowID)
  workflowList, err := GetWorkflows(projectAuthor, projectName)
  if err != nil {
    return nil, err
  }
  for _, workflow := range workflowList {
    workflowMap := workflow.(map[string]interface{})
    if workflowMap["id"].(string) == workflowID {
      if workflowMap["tasks"] == nil {
        return []interface{}{}, nil
      }
      return workflowMap["tasks"].([]interface{}), nil
    }
  }
  return nil, errors.New("workflow not found")
}

func AddTask(projectAuthor string, projectName string, workflowID string, task map[string]interface{}) (string, error) {
  log.Println("addTask:", projectAuthor, projectName, workflowID, task)
  // add task
  addTaskReq := map[string]interface{}{
    "workflowId": workflowID,
    "name": task["Name"],
    "description": task["Description"],
  }
  addTaskReqJson, err := json.Marshal(addTaskReq)
  req, err := scafreq.NewRequest(
    "POST",
    "/user/" + projectAuthor + "/project/" + projectName + "/kanban/task",
    addTaskReqJson,
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

func UpdateTask(projectAuthor string, projectName string, workflowID string, taskID string, task map[string]interface{}) (string, error) {
  log.Println("updateTask:", projectAuthor, projectName, workflowID, task)
  // update task
  updateTaskReq := map[string]interface{}{
    "id": taskID,
    "workflowId": workflowID,
    "name": task["Name"],
    "description": task["Description"],
  }
  updateTaskReqJson, err := json.Marshal(updateTaskReq)
  req, err := scafreq.NewRequest(
    "PUT",
    "/user/" + projectAuthor + "/project/" + projectName + "/kanban/task",
    updateTaskReqJson,
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

func DeleteTask(projectAuthor string, projectName string, workflowID string, taskID string) (string, error) {
  log.Println("deleteTask:", projectAuthor, projectName, workflowID, taskID)
  // delete task
  deleteTaskReq := map[string]interface{}{
    "id": taskID,
    "workflowId": workflowID,
  }
  deleteTaskReqJson, err := json.Marshal(deleteTaskReq)
  req, err := scafreq.NewRequest(
    "DELETE",
    "/user/" + projectAuthor + "/project/" + projectName + "/kanban/task",
    deleteTaskReqJson,
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

func MoveTask(projectAuthor string, projectName string, workflowID string, taskID string, toWorkflowID string) (string, error) {
  log.Println("moveTask:", projectAuthor, projectName, workflowID, taskID, toWorkflowID)
  // move task
  moveTaskReq := map[string]interface{}{
    "id": taskID,
    "workflowId": workflowID,
    "newWorkflowId": toWorkflowID,
  }
  moveTaskReqJson, err := json.Marshal(moveTaskReq)
  req, err := scafreq.NewRequest(
    "PATCH",
    "/user/" + projectAuthor + "/project/" + projectName + "/kanban/task",
    moveTaskReqJson,
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
