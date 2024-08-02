package responses

import (
	"time"
)

type AnnouncementSchema struct {
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type StatusSchema struct {
	Status           string               `json:"status"`
	Version          string               `json:"version"`
	CharactersOnline uint32               `json:"characters_online"`
	Announcements    []AnnouncementSchema `json:"announcements"`
	LastWipe         string               `json:"last_wipe"`
	NextWipe         string               `json:"next_wipe"`
}
