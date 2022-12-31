package action

import (
  "fmt"
  "github.com/urfave/cli/v2"
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
    scafio.PrintRepo(repoMap)
  }
  return nil
}

// func AddRepoAction(c * cli.Context) error {
//   // get input
//   var err error
//   questions := []*survey.Question{}
//   answers := struct {
//     RepoName string
//     RepoUrl string
//   }{}
//   answers.RepoName, err = scafio.GetArg(c, 0)
//   if err != nil {
//     questions = append(questions, repoNameQuestion)
//   }
//   answers.RepoUrl, err = scafio.GetArg(c, 1)
//   if err != nil {
//     questions = append(questions, repoUrlQuestion)
//   }
//   err = survey.Ask(questions, &answers)
//   if err != nil {
//     return err
//   }
//   // add repo
//   message, err := project.AddRepo(answers.RepoName, answers.RepoUrl)
