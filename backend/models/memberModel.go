package models

import "time"

type Member struct {
	GroupId   int64     `json:"group_id" gorm:"primaryKey"`
	UserId    int64     `json:"user_id" gorm:"primaryKey"`
	Grade     int8      `json:"grade" gorm:"not null"` // 0: 마스터, 1: 어드, 2: 노멀
	Color     string    `json:"color" gorm:"size:7;not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`

	Schedule []Schedule `gorm:"foreignkey:GroupId, UserId;references:GroupId, UserId"`
}