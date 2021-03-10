package events

import "encoding/json"

type SpawnEvent struct {
	PlayerId string `json:"player_id"`
}

func (spawnEvent SpawnEvent) ToJson() string {
	event := Event{
		Event:   "spawn",
		Content: spawnEvent,
	}
	jsonEvent, _ := json.Marshal(event)
	return string(jsonEvent)
}
