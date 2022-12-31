package action

import (
  "fmt"
  "errors"
  "github.com/urfave/cli/v2"
  "github.com/AlecAivazis/survey/v2"
  "scaf/cli/project"
  "scaf/cli/scafio"
)

func ListRepoAction(c *cli.Context) error {
  // get local repo list
  localProject, err := project.GetLocalProject()
  if err != nil {
    return err
  }
  repoList, ok := localProject["repos"].([]interface{})
  if !ok {
    fmt.Println("No repo found")
    return nil
  }
  for _, repo := range repoList {
    repoMap, ok := repo.(map[string]interface{})
    if !ok {
      continue
    }
    repoString := scafio.RepoToString(repoMap)
    fmt.Println(repoString)
  }
  return nil
}

func AddRepoAction(c *cli.Context) error {
  // get input
  var err error
  questions := []*survey.Question{}
  answers := struct {
    RepoName string
    RepoUrl string
  }{}
  answers.RepoName, err = scafio.GetArg(c, 0)
  if err != nil {
    questions = append(questions, repoNameQuestion)
  }
  answers.RepoUrl, err = scafio.GetArg(c, 1)
  if err != nil {
    questions = append(questions, repoUrlQuestion)
  }
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  // add repo
  message, err := project.AddRepo(answers.RepoName, answers.RepoUrl)
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func UpdateRepoAction(c *cli.Context) error {
  // select repo
  selectRepoMap, err := selectRepo()
  // get input
  questions := []*survey.Question{
    newRepoNameQuestion,
    newRepoUrlQuestion,
  }
  answers := struct {
    NewRepoName string
    NewRepoUrl string
  }{}
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  if answers.NewRepoName == "" && answers.NewRepoUrl == "" {
    fmt.Println("No update")
    return nil
  }
  if answers.NewRepoName == "" {
    answers.NewRepoName = selectRepoMap["name"].(string)
  }
  if answers.NewRepoUrl == "" {
    answers.NewRepoUrl = selectRepoMap["url"].(string)
  }
  // update repo
  message, err := project.UpdateRepo(selectRepoMap["id"].(string), answers.NewRepoName, answers.NewRepoUrl)
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}

func selectRepo() (map[string]interface{}, error) {
  localProject, err := project.GetLocalProject()
  if err != nil {
    return nil, err
  }
  repoList := localProject["repos"].([]interface{})
  repoStringList := []string{}
  for _, repo := range repoList {
    repoMap, ok := repo.(map[string]interface{})
    if !ok {
      continue
    }
    repoString := scafio.RepoToString(repoMap)
    repoStringList = append(repoStringList, repoString)
  }
  selectRepoQuestion := &survey.Select{
    Message: "Select repo:",
    Options: repoStringList,
  }
  var selectRepo string
  err = survey.AskOne(selectRepoQuestion, &selectRepo)
  if err != nil {
    return nil, err
  }
  var selectRepoMap map[string]interface{}
  for _, repo := range repoList {
    repoMap, ok := repo.(map[string]interface{})
    if !ok {
      continue
    }
    repoString := scafio.RepoToString(repoMap)
    if repoString == selectRepo {
      selectRepoMap = repoMap
      break
    }
  }
  return selectRepoMap, nil
}

func DeleteRepoAction(c *cli.Context) error {
  // select repo
  selectRepoMap, err := selectRepo()
  // delete repo
  message, err := project.DeleteRepo(selectRepoMap["id"].(string))
  if err != nil {
    return err
  }
  fmt.Println(message)
  return nil
}
