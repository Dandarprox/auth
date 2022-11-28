package interfaces

import "github.com/Dandarprox/golang-auth/entities"

type UserServiceI interface {
	SaveUser(user *entities.User) (int, error)
	GetUserById(id int) (*entities.User, error)
	GetUserByEmail(email string) (*entities.User, error)
}
