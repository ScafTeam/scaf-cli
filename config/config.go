package config

import (
  "errors"
  "scaf/cli/user"
  "scaf/cli/project"
)

func SetConfig(category string, field string, value interface{}) (string, error) {
  switch category {
  case User:
    return setUserConfig(field, value)
  case Project:
    return setProjectConfig(field, value)
  default:
    return "", errors.New("Invalid category")
  }
}

func setUserConfig(field string, value interface{}) (string, error) {
  return user.UpdateUser(map[string]interface{}{
    field: value,
  })
}

func setProjectConfig(field string, value interface{}) (string, error) {
  switch field {
  case AddMember:
    return project.AddMember(value.(string))
  case DeleteMember:
    return project.DeleteMember(value.(string))
  default:
    return project.UpdateLocalProject(map[string]interface{}{
      field: value,
    })
  }
}
