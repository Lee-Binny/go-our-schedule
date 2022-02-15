package dto

type CreateGroupDto struct {
	Name      string `json:"name" binding:"required"`
	Color     string `json:"color" binding:"required"`
	UserColor string `json:"user_color" binding:"required"`
}

type UpdateGroupDto struct {
	Id    int64  `json:"id" binding:"required"`
	Name  string `json:"name" binding:"required"`
	Color string `json:"color" binding:"required"`
}