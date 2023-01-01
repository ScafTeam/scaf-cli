package action

import (
  "github.com/AlecAivazis/survey/v2"
  "scaf/cli/project"
  "scaf/cli/config"
)

// user questions
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

// config questions
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

// repo questions
var (
  repoNameQuestion = &survey.Question{
    Name: "RepoName",
    Prompt: &survey.Input{
      Message: "Please input repo name:",
    },
    Validate: survey.Required,
  }
  repoUrlQuestion = &survey.Question{
    Name: "RepoUrl",
    Prompt: &survey.Input{
      Message: "Please input repo url:",
    },
    Validate: survey.Required,
  }
  newRepoNameQuestion = &survey.Question{
    Name: "NewRepoName",
    Prompt: &survey.Input{
      Message: "Please input new repo name (empty to keep the same):",
    },
  }
  newRepoUrlQuestion = &survey.Question{
    Name: "NewRepoUrl",
    Prompt: &survey.Input{
      Message: "Please input new repo url (empty to keep the same):",
    },
  }
)

// workflow questions
var (
  workflowNameQuestion = &survey.Question{
    Name: "WorkflowName",
    Prompt: &survey.Input{
      Message: "Please input board name:",
    },
    Validate: survey.Required,
  }
)
