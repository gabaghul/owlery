package models

import "time"

type EmailingConfigs struct {
	ClientID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Active    bool
}

type ContactLists struct {
	ListID    string
	ClientID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Active    bool
}

func (m EmailingConfigs) IsZero() bool {
	return m.ClientID == 0
}

func (m ContactLists) IsZero() bool {
	return m.ClientID == 0 || m.ListID == ""
}
