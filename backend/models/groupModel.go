package models

import "time"

type Group struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"size:45;not null"`
	Color     string    `json:"color" gorm:"size:7;not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`

	Member   []Member   `gorm:"foreignkey:GroupId;references:Id"`
}