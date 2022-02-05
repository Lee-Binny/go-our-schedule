package models

import "time"

type User struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId    string    `json:"user_id" gorm:"size:15;not null"`
	Password  string    `json:"password" gorm:"size:100;not null"`
	Nickname  string    `json:"nickname" gorm:"size:15;not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`

	Member   []Member   `gorm:"foreignkey:UserId;references:Id"`
}