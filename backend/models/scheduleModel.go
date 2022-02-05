package models

import "time"

type Schedule struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	GroupId   int64     `json:"group_id"`
	UserId    int64     `json:"user_id"`
	Title     string    `json:"title" gorm:"size:45;not null"` // 0: 마스터, 1: 어드, 2: 노멀
	StartAt   time.Time `json:"start_at" gorm:"not null"`
	EndAt     time.Time `json:"end_at" gorm:"not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`
}