package repository

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Dandarprox/golang-auth/entities"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	db map[string]*entities.User
}

func NewUserService() *UserService {
	return &UserService{
		db: make(map[string]*entities.User),
	}
}

func (s *UserService) SaveUser(user *entities.User) (int, error) {
	rand.Seed(time.Now().Unix())

	if user, ok := s.db[user.Email]; ok {
		return 0, fmt.Errorf("the user with mail (%s) already exists", user.Email)
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)

	if err != nil {
		return 0, err
	}

	id := rand.Int()
	s.db[user.Email] = &entities.User{
		Id:       uint(id),
		Email:    user.Email,
		Password: string(hash),
	}

	return id, nil
}

func (s *UserService) GetUserByEmail(email string) (*entities.User, error) {
	userDb, ok := s.db[email]

	if !ok {
		return nil, fmt.Errorf("the user with mail (%s) doesnt exist", email)
	}

	return userDb, nil
}

func (s *UserService) GetUserById(id int) (*entities.User, error) {
	for _, v := range s.db {
		if v.Id == uint(id) {
			return v, nil
		}
	}

	return nil, fmt.Errorf("the user with id (%d) already exists", id)
}
