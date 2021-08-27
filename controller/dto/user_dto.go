package dto

type UserRequest struct {
	Email     string `json:"email" binding:"required,email"`
	Password string `form:"password" json:"password"`
}

type UserResponse struct {
	Password string `form:"password" json:"password"`
}
