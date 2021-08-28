package user

type RequestUser struct {
	Email     string `json:"email" binding:"required,email"`
	Password string `form:"password" json:"password"`
}

type ResponseUser struct {
	Id       int    `form:"id" json:"id"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}
