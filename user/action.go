package user

import (
  "fmt"
  "sort"
  "github.com/urfave/cli/v2"
  "github.com/AlecAivazis/survey/v2"
  "scaf/cli/scafio"
)

func GetUserAction(c *cli.Context) error {
  var err error
  questions := []*survey.Question{}
  answers := struct {
    Email string
  }{}
  answers.Email, err = scafio.GetArg(c, 0)
  if err != nil {
    questions = append(questions, scafio.EmailQuestion)
  }
  err = survey.Ask(questions, &answers)
  if err != nil {
    return err
  }
  userData, err := GetUser(answers.Email)
  if err != nil {
    fmt.Println("what", err)
    return nil
  }

  keys := make([]string, 0, len(userData))
  for k := range userData {
    keys = append(keys, k)
  }
  sort.Strings(keys)
  for _, key := range keys {
    fmt.Printf("%s: %s\n", key, userData[key])
  }

  return nil
}
