# SCAF-CLI

## Usage

```bash
$ scaf <command> [subcommand] [flags...] [arguments...]
```

If required command argument was not provided, scaf will require user input

---

### Auth

#### `scaf user <user email>`

show user data

#### `scaf signin <user email>`

user signin, input password

#### `scaf signup <user email>`

user signup, input password

#### `scaf signout`

user signout

#### `scaf whoami`

show current signed in user

---

### Config

#### `scaf config set <category> <field> <value>`

update config

#### `scaf config get <category> <field>`

get current config value

#### `scaf config password`

change user password, input old and new password to change

---

### Project

#### `scaf project list <--oneline> <user email>`

show all projects of a specific user
`--oneline` will compress output

#### `scaf project create <project name>`

input development mode and development tool, create a new project under current user
and clone it into local

#### `scaf project clone <project author> <project name>`

clone project into local folder

#### `scaf project pull` (inside project folder)

pull newest project state

---

### Repo (inside project folder)

#### `scaf repo list`

list all repo in project

#### `scaf repo add <repo name> <repo url>`

add a repo into project

#### `scaf repo update`

select a repo to update it data

#### `scaf repo delete`

select a repo to delete it

---

## Run source code

go version must over 1.11

```bash
$ go run <project foler path>
```
