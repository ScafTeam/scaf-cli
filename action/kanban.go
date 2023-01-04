package action

import (
  "fmt"
  "strings"
  "github.com/urfave/cli/v2"
  "github.com/AlecAivazis/survey/v2"
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

func AddWorkflowAction(c *cli.Context) error {
  // get local project
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // get workflow name
  questions := []*survey.Question{}
  answers := struct {
    WorkflowName string
  }{}
  answers.WorkflowName, err = scafio.GetArg(c, 0)
  if err != nil {
    questions = append(questions, workflowNameQuestion)
  }
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  // add workflow
  message, err := project.AddWorkflow(projectAuthor, projectName, answers.WorkflowName)
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func UpdateWorkflowAction(c *cli.Context) error {
  // get local project
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // select workflow
  workflowList, err := project.GetWorkflows(projectAuthor, projectName)
  if err != nil {
    return err
  }
  workflowStringList := []string{}
  for _, workflow := range workflowList {
    workflowMap, ok := workflow.(map[string]interface{})
    if !ok {
      continue
    }
    workflowStringList = append(
      workflowStringList, 
      scafio.WorkflowToString(workflowMap, true),
    )
  }
  selectWorkflowPrompt := &survey.Select{
    Message: "Select workflow:",
    Options: workflowStringList,
  }
  selectWorkflowString := ""
  err = survey.AskOne(selectWorkflowPrompt, &selectWorkflowString)
  if err != nil {
    return err
  }
  var selectWorkflow map[string]interface{}
  for _, workflow := range workflowList {
    workflowMap, ok := workflow.(map[string]interface{})
    if !ok {
      continue
    }
    if scafio.WorkflowToString(workflowMap, true) == selectWorkflowString {
      selectWorkflow = workflowMap
      break
    }
  }
  // update workflow
  workflowNamePrompt := &survey.Input{
    Message: "Workflow name:",
    Default: selectWorkflow["name"].(string),
  }
  newWorkflowName := ""
  err = survey.AskOne(workflowNamePrompt, &newWorkflowName)
  if err != nil {
    return err
  }
  message, err := project.UpdateWorkflow(
    projectAuthor,
    projectName,
    selectWorkflow["id"].(string),
    newWorkflowName,
  )
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func DeleteWorkflowAction(c *cli.Context) error {
  // get local project
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // select workflow
  workflowList, err := project.GetWorkflows(projectAuthor, projectName)
  if err != nil {
    return err
  }
  workflowStringList := []string{}
  for _, workflow := range workflowList {
    workflowMap, ok := workflow.(map[string]interface{})
    if !ok {
      continue
    }
    workflowStringList = append(
      workflowStringList, 
      scafio.WorkflowToString(workflowMap, true),
    )
  }
  selectWorkflowPrompt := &survey.Select{
    Message: "Select workflow:",
    Options: workflowStringList,
  }
  selectWorkflowString := ""
  err = survey.AskOne(selectWorkflowPrompt, &selectWorkflowString)
  if err != nil {
    return err
  }
  var selectWorkflow map[string]interface{}
  for _, workflow := range workflowList {
    workflowMap, ok := workflow.(map[string]interface{})
    if !ok {
      continue
    }
    if scafio.WorkflowToString(workflowMap, true) == selectWorkflowString {
      selectWorkflow = workflowMap
      break
    }
  }
  // delete workflow
  message, err := project.DeleteWorkflow(
    projectAuthor,
    projectName,
    selectWorkflow["id"].(string),
  )
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}
