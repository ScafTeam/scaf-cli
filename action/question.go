package action

import (
  "github.com/AlecAivazis/survey/v2"
  "scaf/cli/project"
  "scaf/cli/config"
)

var (
  emailQuestion = &survey.Question{
    Name: "Email",
    Prompt: &survey.Input{ Message: "Please input email:" },
    Validate: survey.Required,
  }
  passwordQuestion = &survey.Question{
    Name: "Password",
    Prompt: &survey.Password{ Message: "Please input your password:" },
    Validate: survey.Required,
  }
  passwordConfirmQuestion = &survey.Question{
    Name: "PasswordConfirm",
    Prompt: &survey.Password{ Message: "Please confirm your password:" },
    Validate: survey.Required,
  }
  newPasswordQuestion = &survey.Question{
    Name: "NewPassword",
    Prompt: &survey.Password{
      Message: "Please input your new password:",
    },
    Validate: survey.Required,
  }
  oldPasswordQuestion = &survey.Question{
    Name: "OldPassword",
    Prompt: &survey.Password{
      Message: "Please input your old password:",
    },
    Validate: survey.Required,
  }
)

// project questions
var (
  projectNameQuestion = &survey.Question{
    Name: "ProjectName",
    Prompt: &survey.Input{ Message: "Please input your project name:" },
    Validate: survey.Required,
  }
  devModeQuestion = &survey.Question{
    Name: "DevMode",
    Prompt: &survey.Select{
      Message: "Please select your development mode:",
      Options: project.DevModes,
    },
    Validate: survey.Required,
  }
  devToolsQuestion = &survey.Question{
    Name: "DevTools",
    Prompt: &survey.MultiSelect{
      Message: "Please select your dev tools:",
      Options: project.DevTools,
    },
  }
)

var (
  configCategoryQuestion = &survey.Question{
    Name: "Category",
    Prompt: &survey.Select{
      Message: "Please select your config category:",
      Options: config.Categories,
    },
    Validate: survey.Required,
  }
  configFieldQuestion = &survey.Question{
    Name: "Field",
    Prompt: &survey.Input{
      Message: "Please input your config field:",
    },
    Validate: survey.Required,
  }
  valueQuestion = &survey.Question{
    Name: "Value",
    Prompt: &survey.Input{
      Message: "Please input your value:",
    },
    Validate: survey.Required,
  }
  boolQuestion = &survey.Question{
    Name: "Value",
    Prompt: &survey.Select{
      Message: "Please select your value:",
      Options: []string{"true", "false"},
    },
    Validate: survey.Required,
  }
)
