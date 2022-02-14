package models

import (
	"go-our-schedule/dto"
	"time"
	"golang.org/x/crypto/bcrypt"
)

type Users []User
type User struct {
	Id        int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	UserId    string    `json:"user_id" gorm:"size:15;not null"`
	Password  string    `json:"password" gorm:"size:100;not null"`
	Nickname  string    `json:"nickname" gorm:"size:15;not null"`
	CreatedAt time.Time `json:"-" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"-" gorm:"autoUpdateTime"`

	Member   []Member   `gorm:"foreignkey:UserId;references:Id"`
}

func NewUser(dto *dto.SignUpUserDto) *User {
	return &User{
		UserId:    dto.UserId,
		Password:  dto.Password,
		Nickname:  dto.Name,
	}
}

func (u *User) HashPassword() error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(hashedPassword)
	return nil
}

func (u *User) IsMatchedPassword(password string) error {
	return bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
}