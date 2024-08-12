package schemas

import (
	"fmt"
	"time"
)

type AnnouncementSchema struct {
	Message   string     `json:"message"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

func (a AnnouncementSchema) String() string {
	return fmt.Sprintf("{Message:%s, CreatedAt: %v}", a.Message, *a.CreatedAt)
}

type StatusSchema struct {
	Status           string                `json:"status"`
	Version          *string               `json:"version,omitempty"`
	CharactersOnline *int32                `json:"characters_online,omitempty"`
	ServerTime       *time.Time            `json:"server_time,omitempty"`
	Accouncements    []*AnnouncementSchema `json:"announcements,omitempty"`
	LastWipe         string                `json:"last_wipe"`
	NextWipe         string                `json:"next_wipe"`
}

func (s StatusSchema) String() string {
	announcements := ""
	first := true
	for _, announcement := range s.Accouncements {
		if first {
			announcements += fmt.Sprintf("%v", *announcement)
			continue
		}
		announcements += fmt.Sprintf(", %v", *announcement)
	}
	return fmt.Sprintf(
		"{Status: %s, Version: %s, CharactersOnline: %d, ServerTime: %v, Announcements: [%s], LastWipe: %s, NextWipe: %s}",
		s.Status,
		*s.Version,
		*s.CharactersOnline,
		*s.ServerTime,
		announcements,
		s.LastWipe,
		s.NextWipe,
	)
}
