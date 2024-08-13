package schemas

import "time"

type ActiveEventSchema struct {
	Name         string    `json:"name"`
	Map          MapSchema `json:"map"`
	PreviousSkin string    `json:"previous_skin"`
	Duration     int32     `json:"duration"`
	Expiration   time.Time `json:"expiration"`
	CreatedAt    time.Time `json:"created_at"`
}
