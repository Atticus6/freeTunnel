package models

import (
	"time"
)

type Tunnel struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"size:100;not null"`
	Port      int64     `json:"port"`
	Host      string    `json:"host" gorm:"size:100;not null"`
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
