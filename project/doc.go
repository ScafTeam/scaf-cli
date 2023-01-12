package project

import (
  "encoding/json"
  "scaf/cli/scafio"
  "scaf/cli/scafreq"
)

func GetDocs(projectAuthor string, projectName string) ([]interface{}, error) {
  req, err := scafreq.NewRequest(
    "GET",
    "/user/" + projectAuthor + "/project/" + projectName + "/docs",
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
  return body["docs"].(map[string]interface{})["docs"].([]interface{}), nil
}

func AddDoc(projectAuthor, projectName, title, content string) (string, error) {
  // add doc
  addDocReq := map[string]interface{}{
    "title": title,
    "content": content,
  }
  addDocReqJson, err := json.Marshal(addDocReq)
  req, err := scafreq.NewRequest(
    "POST",
    "/user/" + projectAuthor + "/project/" + projectName + "/docs",
    addDocReqJson,
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

func UpdateDoc(projectAuthor, projectName, docId, title, content string) (string, error) {
  // update doc
  updateDocReq := map[string]interface{}{
    "id": docId,
    "title": title,
    "content": content,
  }
  updateDocReqJson, err := json.Marshal(updateDocReq)
  req, err := scafreq.NewRequest(
    "PUT",
    "/user/" + projectAuthor + "/project/" + projectName + "/docs/",
    updateDocReqJson,
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

func DeleteDoc(projectAuthor, projectName, docId string) (string, error) {
  // delete doc
  deleteDocReq := map[string]interface{}{
    "id": docId,
  }
  deleteDocReqJson, err := json.Marshal(deleteDocReq)
  req, err := scafreq.NewRequest(
    "DELETE",
    "/user/" + projectAuthor + "/project/" + projectName + "/docs/",
    deleteDocReqJson,
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
