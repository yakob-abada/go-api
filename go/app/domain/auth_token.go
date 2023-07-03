package domain

// TokenService defines methods the handler layer expects to interact
// with in regards to producing JWTs as string
type IAuthToken interface {
	GenerateToken(username string, userId int8) (*TokenResponse, error)
	Validate(string) (*Claims, error)
}
