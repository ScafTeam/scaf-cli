package project

// dev mode constants
const (
  Waterfall string = "waterfall"
  Scrum string = "scrum"
)

// dev modes
var DevModes = []string{
  Waterfall,
  Scrum,
}


// dev tools constants
type DevToolType string
const (
  GoogleCalendar string = "google_calendar"
  Kanban string = "kanban"
)

// dev tools
var DevTools = []string{
  GoogleCalendar,
  Kanban,
}
