package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/usagifm/golang-jwt/dto"
	"github.com/usagifm/golang-jwt/helper"
	"github.com/usagifm/golang-jwt/service"
)

// user controller is a contract that defines what this controller can do
type UserController interface {
	Update(context *gin.Context)
	Profile(context *gin.Context)
	MyDiaries(context *gin.Context)
}

type userController struct {
	userService service.UserService
	jwtService  service.JWTService
}

// NewUSerController is creating a new instance of UserController
func NewUserController(userService service.UserService, jwtService service.JWTService) UserController {

	return &userController{
		userService: userService,
		jwtService:  jwtService,
	}

}

func (c *userController) Update(context *gin.Context) {
	var userUpdateDTO dto.UserUpdateDTO
	errDTO := context.ShouldBind(&userUpdateDTO)
	if errDTO != nil {

		res := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	authHeader := context.GetHeader("Authorization")

	splitToken := strings.Split(authHeader, "Bearer ")
	authHeader = splitToken[1]
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {

		panic(errToken.Error())
	}

	claims := token.Claims.(jwt.MapClaims)
	id, err := strconv.ParseUint(fmt.Sprintf("%v", claims["user_id"]), 10, 64)
	if err != nil {
		panic(err.Error())
	}
	userUpdateDTO.ID = id
	u := c.userService.Update(userUpdateDTO)
	res := helper.BuildResponse(true, "OK!", u)
	context.JSON(http.StatusOK, res)

}

func (c *userController) Profile(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")

	splitToken := strings.Split(authHeader, "Bearer ")
	authHeader = splitToken[1]

	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())

	}

	claims := token.Claims.(jwt.MapClaims)
	user := c.userService.Profile(fmt.Sprintf("%v", claims["user_id"]))
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}

func (c *userController) MyDiaries(context *gin.Context) {
	authHeader := context.GetHeader("Authorization")

	splitToken := strings.Split(authHeader, "Bearer ")
	authHeader = splitToken[1]

	token, err := c.jwtService.ValidateToken(authHeader)
	if err != nil {
		panic(err.Error())

	}

	claims := token.Claims.(jwt.MapClaims)
	user := c.userService.MyDiaries(fmt.Sprintf("%v", claims["user_id"]))
	res := helper.BuildResponse(true, "OK!", user)
	context.JSON(http.StatusOK, res)
}
