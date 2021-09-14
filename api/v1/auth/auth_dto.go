package auth

type RequestToken struct {
	Token string `json:"token" binding:"required"`
}
