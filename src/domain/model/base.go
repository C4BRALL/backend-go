package model

import "time"

type Base struct {
	ID        int        `json:"id"`
	Uuid      string     `json:"uuid"`
	Version   int        `json:"version"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
	DeletedAt *time.Time `json:"-"`
}
