package dto

type SignUpUserDto struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
	Name     string `json:"name" binding:"required"`
}

type SignInUserDto struct {
	UserId   string `json:"user_id" binding:"required"`
	Password string `json:"password" binding:"required"`
}