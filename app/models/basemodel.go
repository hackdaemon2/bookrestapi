package models

import "time"

type BaseModel struct {
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at;type:timestamp;not null"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at;type:timestamp; not null"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;type:timestamp;null"`
	Deleted   int        `json:"deleted" gorm:"column:deleted;type:bit(1);not null"`
}
