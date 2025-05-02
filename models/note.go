package models

import (
	"time"

	"github.com/lib/pq"
)

type Priority int

const (
	Low Priority = iota
	Medium
	High
)

type Note struct {
	ID        string         `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id,omitempty"`
	UserID    string         `gorm:"type:uuid;not null;index" json:"user_id,omitempty"`
	User      User           `gorm:"foreignKey:UserID;references:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
	Title     string         `gorm:"type:varchar(255);not null" json:"title,omitempty"`
	Content   string         `gorm:"type:text;not null" json:"content,omitempty"`
	Tags      pq.StringArray `gorm:"type:text[]" json:"tags,omitempty"`
	Priority  *Priority      `gorm:"type:int;default:0" json:"priority,omitempty"`
	CreatedAt time.Time      `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}
