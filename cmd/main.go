package main

import (
	"log"

	"github.com/Dandarprox/golang-auth/authController"
	"github.com/Dandarprox/golang-auth/repository"
	"github.com/Dandarprox/golang-auth/router"
	"github.com/Dandarprox/golang-auth/service/jwt"
)

func main() {
	authController := authController.NewAuthController(repository.NewUserService(), jwt.NewJwt())

	r := router.NewRouter(authController)

	if err := r.Start(); err != nil {
		log.Fatal(err)
	}
}
