package events

import "encoding/json"

type DisconnectEvent struct {
	PlayerId string `json:"player_id"`
}

func (disconnectEvent DisconnectEvent) ToJson() string {
	event := Event{
		Event:   "disconnect",
		Content: disconnectEvent,
	}
	jsonEvent, _ := json.Marshal(event)
	return string(jsonEvent)
}
