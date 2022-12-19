package project

import (
  "log"
  "os"
  "net/http"
  "io/ioutil"
  "encoding/json"
)

func getProjects(email string) ([]interface{}, error) {
  client := &http.Client{}
  req, err := http.NewRequest(
    "GET",
    os.Getenv("SCAF_BACKEND_URL") + "/" + email + "/project",
    nil,
  )
  if err != nil {
    return nil, err
  }

  resp, err := client.Do(req)
  if err != nil {
    return nil, err
  }

  log.Println(resp.StatusCode)
  data, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    return nil, err
  }

  dataJSON := make(map[string]interface{})
  json.Unmarshal(data, &dataJSON)

  return dataJSON["projects"].([]interface{}), nil

  // dataIndent, err := json.MarshalIndent(dataJSON, "", "  ")
  // log.Println(string(dataIndent))
}
