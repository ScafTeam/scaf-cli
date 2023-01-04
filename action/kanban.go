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

func selectWorkflow(projectAuthor, projectName string) (map[string]interface{}, error) {
  workflowList, err := project.GetWorkflows(projectAuthor, projectName)
  if err != nil {
    return nil, err
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
    return nil, err
  }
  var selectedWorkflow map[string]interface{}
  for _, workflow := range workflowList {
    workflowMap, ok := workflow.(map[string]interface{})
    if !ok {
      continue
    }
    if scafio.WorkflowToString(workflowMap, true) == selectWorkflowString {
      selectedWorkflow = workflowMap
      break
    }
  }
  return selectedWorkflow, nil
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
  selectedWorkflow, err := selectWorkflow(projectAuthor, projectName)
  if err != nil {
    return err
  }
  // update workflow
  workflowNamePrompt := &survey.Input{
    Message: "Workflow name:",
    Default: selectedWorkflow["name"].(string),
  }
  newWorkflowName := ""
  err = survey.AskOne(workflowNamePrompt, &newWorkflowName)
  if err != nil {
    return err
  }
  message, err := project.UpdateWorkflow(
    projectAuthor,
    projectName,
    selectedWorkflow["id"].(string),
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
  selectedWorkflow, err := selectWorkflow(projectAuthor, projectName)
  if err != nil {
    return err
  }
  // delete workflow
  message, err := project.DeleteWorkflow(
    projectAuthor,
    projectName,
    selectedWorkflow["id"].(string),
  )
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func ListTaskAction(c *cli.Context) error {
  // get local project
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // select workflow
  selectedWorkflow, err := selectWorkflow(projectAuthor, projectName)
  if err != nil {
    return err
  }
  // get task list
  taskList, err := project.GetTasks(
    projectAuthor,
    projectName,
    selectedWorkflow["id"].(string),
  )
  if err != nil {
    return err
  }
  for _, task := range taskList {
    taskMap, ok := task.(map[string]interface{})
    if !ok {
      continue
    }
    taskString := scafio.TaskToString(taskMap, c.Bool("oneline"))
    fmt.Println(taskString)
  }
  return nil
}

func AddTaskAction(c *cli.Context) error {
  // get local project
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // select workflow
  selectedWorkflow, err := selectWorkflow(projectAuthor, projectName)
  if err != nil {
    return err
  }
  // input
  questions := []*survey.Question{
    {
      Name: "Name",
      Prompt: &survey.Input{
        Message: "Task name:",
      },
      Validate: survey.Required,
    },
    {
      Name: "Description",
      Prompt: &survey.Input{
        Message: "Task description:",
      },
    },
  }
  answers := map[string]interface{}{}
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  // add task
  message, err := project.AddTask(
    projectAuthor,
    projectName,
    selectedWorkflow["id"].(string),
    answers,
  )
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func selectTask(projectAuthor, projectName, workflowID string) (map[string]interface{}, error) {
  taskList, err := project.GetTasks(projectAuthor, projectName, workflowID)
  if err != nil {
    return nil, err
  }
  taskStringList := []string{}
  for _, task := range taskList {
    taskMap, ok := task.(map[string]interface{})
    if !ok {
      continue
    }
    taskStringList = append(
      taskStringList, 
      scafio.TaskToString(taskMap, true),
    )
  }
  selectTaskPrompt := &survey.Select{
    Message: "Select task:",
    Options: taskStringList,
  }
  selectTaskString := ""
  err = survey.AskOne(selectTaskPrompt, &selectTaskString)
  if err != nil {
    return nil, err
  }
  var selectTask map[string]interface{}
  for _, task := range taskList {
    taskMap, ok := task.(map[string]interface{})
    if !ok {
      continue
    }
    if scafio.TaskToString(taskMap, true) == selectTaskString {
      selectTask = taskMap
      break
    }
  }
  return selectTask, nil
}

func UpdateTaskAction(c *cli.Context) error {
  // get local project
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // select workflow
  selectedWorkflow, err := selectWorkflow(projectAuthor, projectName)
  if err != nil {
    return err
  }
  // select task
  selectedTask, err := selectTask(
    projectAuthor,
    projectName,
    selectedWorkflow["id"].(string),
  )
  if err != nil {
    return err
  }
  // update task
  questions := []*survey.Question{
    {
      Name: "Name",
      Prompt: &survey.Input{
        Message: "Task name:",
        Default: selectedTask["name"].(string),
      },
      Validate: survey.Required,
    },
    {
      Name: "Description",
      Prompt: &survey.Input{
        Message: "Task description:",
        Default: selectedTask["description"].(string),
      },
    },
  }
  answers := map[string]interface{}{}
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  message, err := project.UpdateTask(
    projectAuthor,
    projectName,
    selectedWorkflow["id"].(string),
    selectedTask["id"].(string),
    answers,
  )
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func DeleteTaskAction(c *cli.Context) error {
  // get local project
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // select workflow
  selectedWorkflow, err := selectWorkflow(projectAuthor, projectName)
  if err != nil {
    return err
  }
  // select task
  selectedTask, err := selectTask(
    projectAuthor,
    projectName,
    selectedWorkflow["id"].(string),
  )
  if err != nil {
    return err
  }
  // delete task
  message, err := project.DeleteTask(
    projectAuthor,
    projectName,
    selectedWorkflow["id"].(string),
    selectedTask["id"].(string),
  )
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func MoveTaskAction(c *cli.Context) error {
  // get local project
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // select workflow
  selectedWorkflow, err := selectWorkflow(projectAuthor, projectName)
  if err != nil {
    return err
  }
  // select task
  selectedTask, err := selectTask(
    projectAuthor,
    projectName,
    selectedWorkflow["id"].(string),
  )
  if err != nil {
    return err
  }
  // select workflow to move
  fmt.Println("Select workflow to move:")
  selectedWorkflowToMove, err := selectWorkflow(projectAuthor, projectName)
  // move task
  message, err := project.MoveTask(
    projectAuthor,
    projectName,
    selectedWorkflow["id"].(string),
    selectedTask["id"].(string),
    selectedWorkflowToMove["id"].(string),
  )
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

