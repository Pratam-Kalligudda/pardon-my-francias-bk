package models

import "time"

type User struct {
	UserId       string     `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"user_id,omitempty"`
	FirstName    string     `gorm:"type:varchar(100);not null" json:"first_name,omitempty"`
	LastName     string     `gorm:"type:varchar(100);not null" json:"last_name,omitempty"`
	UserName     string     `gorm:"type:varchar(100);unique;not null" json:"user_name,omitempty"`
	Email        string     `gorm:"type:varchar(100);unique;not null" json:"email,omitempty"`
	CreatedAt    time.Time  `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt    time.Time  `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
	Password     string     `gorm:"type:varchar(255);not null" json:"password,omitempty"`
	Bio          *string    `gorm:"type:text" json:"bio,omitempty"`
	ProfilePhoto *int       `gorm:"type:int;default:-1" json:"profile_photo,omitempty"`
	Dob          *time.Time `gorm:"type:date" json:"dob,omitempty"`
}
