package driver

type AuthService interface {
	Login(token string) error
}
