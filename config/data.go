package config

import (
  "errors"
  "github.com/AlecAivazis/survey/v2"
  "scaf/cli/project"
)

// Config Categories constants
const (
  User string = "user"
  Project string = "project"
)

// Config Categories
var Categories = []string{
  User,
  Project,
}

// Config Fields constants
// User
const (
  Nickname string = "nickname"
  Bio string = "bio"
  Avatar string = "avatar"
)

// Project
const (
  Name string = "name"
  DevTools string = "devTools"
  DevMode string = "devMode"
  AddMember string = "addMember"
  DeleteMember string = "deleteMember"
)

func GetCategoryPrompt() (survey.Prompt, error) {
  return &survey.Select{
    Message: "Select a category:",
    Options: Categories,
  }, nil
}

func GetFieldPrompt(category string) (survey.Prompt, error) {
  message := "Select a field:"
  switch category {
  case User:
    return &survey.Select{
      Message: message,
      Options: []string{
        Nickname,
        Bio,
        Avatar,
      },
    }, nil
  case Project:
    return &survey.Select{
      Message: message,
      Options: []string{
        Name,
        DevTools,
        DevMode,
        AddMember,
        DeleteMember,
      },
    }, nil
  default:
    return nil, errors.New("Invalid category")
  }
}

func GetValuePrompt(category string, field string) (survey.Prompt, error) {
  switch category {
  case User:
    return getUserValuePrompt(field)
  case Project:
    return getProjectValuePrompt(field)
  default:
    return nil, errors.New("Invalid category")
  }
}

func getUserValuePrompt(field string) (survey.Prompt, error) {
  message := "Enter a value:"
  switch field {
  case Nickname:
    return &survey.Input{
      Message: message,
    }, nil
  case Bio:
    return &survey.Input{
      Message: message,
    }, nil
  case Avatar:
    return &survey.Input{
      Message: message,
    }, nil
  default:
    return nil, errors.New("Invalid field")
  }
}

func getProjectValuePrompt(field string) (survey.Prompt, error) {
  switch field {
  case Name:
    return &survey.Input{
      Message: "Enter New Project Name:",
    }, nil
  case DevMode:
    return &survey.Select{
      Message: "Select a dev mode:",
      Options: project.DevModes,
    }, nil
  case DevTools:
    return &survey.MultiSelect{
      Message: "Select dev tools:",
      Options: project.DevTools,
    }, nil
  case AddMember:
    return &survey.Input{
      Message: "Enter a member's email:",
    }, nil
  case DeleteMember:
    members, err := project.GetMembers()
    if err != nil {
      return nil, err
    }
    memberString := make([]string, len(members))
    for i, member := range members {
      memberString[i] = member.(string)
    }
    return &survey.Select{
      Message: "Select a member to delete:",
      Options: memberString,
    }, nil
  default:
    return nil, errors.New("Invalid field")
  }
}
