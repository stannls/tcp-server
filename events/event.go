package events

type Event struct {
	Event   string      `json:"event"`
	Content interface{} `json:"content"`
}
