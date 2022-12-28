package scafio

import (
  "fmt"
  "errors"
  "net/http"
  "encoding/json"
  "io"
  "bytes"
  "github.com/AlecAivazis/survey/v2"
  "github.com/urfave/cli/v2"
)

var (
  EmailQuestion = &survey.Question{
    Name: "Email",
    Prompt: &survey.Input{ Message: "Please input email:" },
    Validate: survey.Required,
  }
  PasswordQuestion = &survey.Question{
    Name: "Password",
    Prompt: &survey.Password{ Message: "Please input your password:" },
    Validate: survey.Required,
  }
  PasswordConfirmQuestion = &survey.Question{
    Name: "PasswordConfirm",
    Prompt: &survey.Password{ Message: "Please confirm your password:" },
    Validate: survey.Required,
  }
  ProjectNameQuestion = &survey.Question{
    Name: "ProjectName",
    Prompt: &survey.Input{ Message: "Please input your project name:" },
    Validate: survey.Required,
  }
)

func GetArg(c *cli.Context, index int) (string, error) {
  if c.NArg() > index {
    return c.Args().Get(index), nil
  }
  return "", errors.New("argument not found")
}

func PrintProject(projectMap map[string]interface{}) {
  fmt.Printf("* [%s] %s (%s)\n", projectMap["Id"], projectMap["Name"], projectMap["Author"])
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
