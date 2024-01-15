package domain

import "time"

type Model struct {
	ID        uint       `json:"id,omitempty" gorm:"primary_key"`
	CreatedAt time.Time  `json:"create_at,omitempty"`
	UpdatedAt time.Time  `json:"update_at,omitempty"`
	DeletedAt *time.Time `json:"delete_at,omitempty"  sql:"index"`
}
