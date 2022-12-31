package scafio

import (
  "fmt"
  "errors"
  "net/http"
  "encoding/json"
  "io"
  "bytes"
  "github.com/urfave/cli/v2"
)

func GetArg(c *cli.Context, index int) (string, error) {
  if c.NArg() > index {
    return c.Args().Get(index), nil
  }
  return "", errors.New("argument not found")
}

func PrintProject(projectMap map[string]interface{}, oneline bool) {
  if oneline {
    fmt.Printf("* %v - %v (%v)\n", projectMap["id"], projectMap["name"], projectMap["author"])
  } else {
    fmt.Printf("* %v\n", projectMap["id"])
    fields := []string{"name", "author", "createOn", "devMode", "devTools", "members", "repos"}
    for _, field := range fields {
      fmt.Printf("  %v: %v\n", field, projectMap[field])
    }
    fmt.Println()
  }
}

func PrintRepo(repoMap map[string]interface{}) {
  fmt.Printf("* %v - %v (%v)\n", repoMap["id"], repoMap["name"], repoMap["url"])
}

// read json format response body and return a map
// then restore response body (can be read again)
func ReadBody(resp *http.Response) (map[string]interface{}, error) {
  if resp.Body != nil {
    bodyBytes, err := io.ReadAll(resp.Body)
    if err != nil {
      return nil, err
    }
    resp.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
    return ReadBodyFromBytes(bodyBytes)
  }
  return nil, nil
}

func ReadBodyFromBytes(body []byte) (map[string]interface{}, error) {
  bodyMap := make(map[string]interface{})
  err := json.Unmarshal(body, &bodyMap)
  if err != nil {
    bodyMap["message"] = string(body)
  }

  return bodyMap, nil
}
