package router

import (
	"github.com/Dandarprox/golang-auth/authController"
	"github.com/gin-gonic/gin"
)

type Router struct {
	authController *authController.AuthController
	_              any
}

func NewRouter(authcontroller *authController.AuthController) *Router {
	return &Router{
		authController: authcontroller,
	}
}

func (router *Router) Start() error {
	r := gin.Default()

	r.POST("/login", router.authController.Login)
	r.POST("/sigin", router.authController.Signin)
	return r.Run()
}
