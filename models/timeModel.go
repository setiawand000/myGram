package models

import "time"

type TimeModel struct {
	ID        uint       `gorm:"primarykey" json:"id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
