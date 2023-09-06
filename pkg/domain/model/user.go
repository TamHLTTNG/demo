package model

import "time"

type User struct {
	ID        uint      `gorm:"primary_key" json:"-"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"unique;not null" json:"email"`
	ImageURL  string    `json:"imageURL"`
	Password  string    `gorm:"not null" json:"password"`
	RoleID    uint      `gorm:"not null" json:"roleID"`
	CreatedAt time.Time `json:"-" time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `json:"-" time_format:"2006-01-02 15:04:05"`
}
