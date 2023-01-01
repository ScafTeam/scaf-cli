package action

import (
  "fmt"
  "strings"
  "github.com/urfave/cli/v2"
  "scaf/cli/project"
  "scaf/cli/scafio"
)

func ListWorkflowAction(c *cli.Context) error {
  // get local project
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // get workflow list
  workflowList, err := project.GetWorkflows(projectAuthor, projectName)
  if err != nil {
    return err
  }
  // check if user specify workflow name
  workflowName, err := scafio.GetArg(c, 0)
  if err != nil {
    // print workflow list
    for _, workflow := range workflowList {
      workflowMap, ok := workflow.(map[string]interface{})
      if !ok {
        continue
      }
      workflowString := scafio.WorkflowToString(workflowMap, c.Bool("oneline"))
      fmt.Println(workflowString)
    }
  } else {
    // print only one workflow
    for _, workflow := range workflowList {
      workflowMap, ok := workflow.(map[string]interface{})
      if !ok {
        continue
      }
      if strings.ToLower(workflowMap["name"].(string)) ==
         strings.ToLower(workflowName) {
        workflowString := scafio.WorkflowToString(workflowMap, c.Bool("oneline"))
        fmt.Println(workflowString)
        break
      }
    }
  }
  return nil
}
