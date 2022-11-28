package authController

import (
	"net/http"

	"github.com/Dandarprox/golang-auth/entities"
	"github.com/Dandarprox/golang-auth/service/interfaces"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthController struct {
	userService interfaces.UserServiceI
	jwtService  interfaces.JwtI
}

func NewAuthController(userService interfaces.UserServiceI, jwtService interfaces.JwtI) *AuthController {
	return &AuthController{
		userService: userService,
		jwtService:  jwtService,
	}
}

func (r *AuthController) Login(c *gin.Context) {
	var userLogin entities.UserLogin

	if err := c.Bind(&userLogin); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := r.userService.GetUserByEmail(userLogin.Email)

	if err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	if user == nil {
		handleError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLogin.Password)); err != nil {
		handleError(c, http.StatusUnauthorized, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": r.jwtService.GenerateToken(user),
	})
}

func handleError(c *gin.Context, status int, errorMessage string) {
	errorEntity := entities.NewErrorEntiy(errorMessage, c.ClientIP(), int16(status))
	c.JSON(status, errorEntity)
}

func (r *AuthController) Signin(c *gin.Context) {
	var userSignin entities.User

	if err := c.Bind(&userSignin); err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	id, err := r.userService.SaveUser(&userSignin)

	if err != nil {
		handleError(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.String(http.StatusCreated, "The user with id: (%d) was successfully created", id)
}
