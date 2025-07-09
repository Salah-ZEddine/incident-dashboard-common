package models

import "time"

	
type Alert struct {
    ID        uint      `gorm:"primaryKey"`
    CreatedAt time.Time
    Source    string    `json:"source"`
    Message   string    `json:"message"`
    Count     int       `json:"count"`
}