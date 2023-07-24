package models

import "time"

type EmailingConfig struct {
	ClientID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Active    bool
}

type ContactList struct {
	ListID    string
	ClientID  int64
	CreatedAt time.Time
	UpdatedAt time.Time
	Active    bool
}

type Contact struct {
	ID        string
	ClientID  string
	Firstname string
	Lastname  string
	Email     string
	Status    string
}

func (m EmailingConfig) IsZero() bool {
	return m.ClientID == 0
}

func (m ContactList) IsZero() bool {
	return m.ClientID == 0 || m.ListID == ""
}
