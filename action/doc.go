package action

import (
  "errors"
  "log"
  "fmt"
  "github.com/urfave/cli/v2"
  "github.com/AlecAivazis/survey/v2"
  "scaf/cli/project"
  "scaf/cli/scafio"
)

func selectDoc(projectAuthor, projectName string) (map[string]interface{}, error) {
  docList, err := project.GetDocs(projectAuthor, projectName)
  if err != nil {
    return nil, err
  }
  docListStringList := []string{}
  for _, doc := range docList {
    docMap, ok := doc.(map[string]interface{})
    if !ok {
      continue;
    }
    docListStringList = append(
      docListStringList,
      scafio.DocToString(docMap, true),
    )
  }
    
  selectDocPrompt := &survey.Select{
    Message: "Select a doc:",
    Options: docListStringList,
  }
  selectedDocString := ""
  err = survey.AskOne(selectDocPrompt, &selectedDocString)
  if err != nil {
    return nil, err
  }
  for _, doc := range docList {
    docMap, ok := doc.(map[string]interface{})
    if !ok {
      continue;
    }
    if selectedDocString == scafio.DocToString(docMap, true) {
      return docMap, nil
    }
  }
  return nil, errors.New("doc not found")
}

func ShowDocAction(c *cli.Context) error {
  log.Println("ListDocAction")
  // get local project
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // select doc
  docMap, err := selectDoc(projectAuthor, projectName)
  if err != nil {
    return err
  }
  // print doc
  fmt.Println(scafio.DocToString(docMap, false))
  return nil
}

func AddDocAction(c *cli.Context) error {
  log.Println("AddDocAction")
  // get local project
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // get input
  questions := []*survey.Question{
    {
      Name: "Title",
      Prompt: &survey.Input{
        Message: "Doc Title:",
      },
      Validate: survey.Required,
    },
    {
      Name: "Content",
      Prompt: &survey.Editor{
        Message: "Doc Content:",
        FileName: "*.md",
      },
    },
  }
  answers := struct {
    Title string
    Content string
  }{}
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  // add doc
  message, err := project.AddDoc(projectAuthor, projectName, answers.Title, answers.Content)
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func UpdateDocAction(c *cli.Context) error {
  log.Println("UpdateDocAction")
  // get local project
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // select doc
  docMap, err := selectDoc(projectAuthor, projectName)
  if err != nil {
    return err
  }
  // get input
  questions := []*survey.Question{
    {
      Name: "Title",
      Prompt: &survey.Input{
        Message: "Doc Title:",
        Default: docMap["title"].(string),
      },
      Validate: survey.Required,
    },
    {
      Name: "Content",
      Prompt: &survey.Editor{
        Message: "Doc Content:",
        FileName: "*.md",
        Default: docMap["content"].(string),
        HideDefault: true,
        AppendDefault: true,
      },
    },
  }
  answers := struct {
    Title string
    Content string
  }{}
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  // update doc
  message, err := project.UpdateDoc(projectAuthor, projectName, docMap["id"].(string), answers.Title, answers.Content)
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func DeleteDocAction(c *cli.Context) error {
  log.Println("DeleteDocAction")
  // get local project
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  projectAuthor := localProject["author"].(string)
  projectName := localProject["name"].(string)
  // select doc
  docMap, err := selectDoc(projectAuthor, projectName)
  if err != nil {
    return err
  }
  // delete doc
  message, err := project.DeleteDoc(projectAuthor, projectName, docMap["id"].(string))
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}
