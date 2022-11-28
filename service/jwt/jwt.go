package jwt

import "github.com/Dandarprox/golang-auth/entities"

type Jwt struct{}

func NewJwt() *Jwt {
	return &Jwt{}
}

func (j *Jwt) GenerateToken(user *entities.User) string {
	return "hola"
}

func (j *Jwt) ValidateToken(token string) (bool, error) {
	return true, nil
}

func (j *Jwt) GetTokenUser(token string) (*entities.User, error) {
	return &entities.User{
		Id:       1,
		Email:    "asdasd@gmail.com",
		Password: "asdqijqwelja",
	}, nil
}
