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

input password to signin

#### `scaf signup <user email>`

input password twice to signup

#### `scaf signout`

signout

#### `scaf whoami`

show current signed in user

---

### Config

#### `scaf config set`

update config
will prompt menu to select

can add/delete project member

#### `scaf config get <category> <field>`

get current config value

#### `scaf config password`

change user password, input old and new password to change

---

### Project

#### `scaf project list [--oneline] <user email>`

show all projects of a specific user
`--oneline` will compress output

#### `scaf project create <project name>`

input development mode and development tool, create a new project under current user
and clone it into local

#### `scaf project delete <project name>`

#### `scaf project clone <project author> <project name>`

clone project into local folder

#### `scaf project pull` (inside project folder)

pull newest project state

---

### Repository (inside project folder)

#### `scaf repo list`

list all repo in project

#### `scaf repo add <repo name> <repo url>`

add a repo into project

#### `scaf repo pull`

select a repo and pull it into local

#### `scaf repo update`

select a repo to update it data

#### `scaf repo delete`

select a repo to delete it

---

### Document (inside project folder)

#### `scaf doc show`

select a document to show content

#### `scaf doc add`

add a new document

#### `scaf doc update`

select a document to update

#### `scaf doc delete`

select a document to delete

---

### Kanban (inside project folder)

#### `scaf kanban list [--oneline] [board name]`

list all kanban boards
if board name is specified, show only that kanban
`--oneline` will compress output

#### `scaf kanban add`

add a new kanban board

#### `scaf kanban update`

update kanban board name

#### `scaf kanban delete`

delete a kanban board

---

#### `scaf kanban task list [--oneline]`

show all task in selected board

#### `scaf kanban task add`

add a task to selected board

#### `scaf kanban task update`

update a task

#### `scaf kanban task delete`

delete a task

#### `scaf kanban task move`

move a task to another board

---

### qa

#### `scaf qa`

get q&a url

## Run source code

go version must over 1.11

```bash
$ go run <project foler path>
```
