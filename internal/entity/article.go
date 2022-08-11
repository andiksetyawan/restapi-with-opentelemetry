package entity

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Article struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	UserID    uint           `json:"user_id"`
	User      User           `json:"user"`
	Slug      string         `gorm:"not null" json:"slug"`
	Title     string         `gorm:"not null" json:"title"`
	Content   string         `gorm:"null" json:"content"`
	CreatedAt int64          `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt int64          `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

func (a *Article) ToJSON() string {
	b, _ := json.Marshal(a)
	return string(b)
}
