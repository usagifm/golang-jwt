package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/usagifm/golang-jwt/dto"
	"github.com/usagifm/golang-jwt/entity"
	"github.com/usagifm/golang-jwt/helper"
	"github.com/usagifm/golang-jwt/service"
)

type DiaryController interface {
	All(context *gin.Context)
	FindById(context *gin.Context)
	Insert(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
}

type diaryController struct {
	diaryService service.DiaryService
	jwtService   service.JWTService
}

func NewDiaryController(diaryServ service.DiaryService, jwtServ service.JWTService) DiaryController {
	return &diaryController{
		diaryService: diaryServ,
		jwtService:   jwtServ,
	}
}

func (c *diaryController) All(context *gin.Context) {
	var diaries []entity.Diary = c.diaryService.All()
	res := helper.BuildResponse(true, "OK", diaries)
	context.JSON(http.StatusOK, res)
}

func (c *diaryController) FindById(context *gin.Context) {
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		res := helper.BuildErrorResponse("Diary ID parameter not found !", err.Error(), helper.EmptyObj{})
		context.AbortWithStatusJSON(http.StatusBadRequest, res)
		return
	}

	var diary entity.Diary = c.diaryService.FindById(id)
	if (diary == entity.Diary{}) {
		res := helper.BuildErrorResponse("Data not found !", "No Data given with Id ", helper.EmptyObj{})
		context.JSON(http.StatusNotFound, res)

	} else {
		res := helper.BuildResponse(true, "OK", diary)
		context.JSON(http.StatusOK, res)
	}
}

func (c *diaryController) Insert(context *gin.Context) {
	var diaryCreateDTO dto.DiaryCreateDTO
	errDTO := context.ShouldBind(&diaryCreateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request ", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
	} else {
		authHeader := context.GetHeader("Authorization")

		splitToken := strings.Split(authHeader, "Bearer ")
		authHeader = splitToken[1]
		userID := c.getUserIDByToken(authHeader)
		intUserID, err := strconv.ParseUint(userID, 10, 64)
		if err == nil {
			diaryCreateDTO.UserID = intUserID
		}
		result := c.diaryService.Insert(diaryCreateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusCreated, response)
	}

}

func (c *diaryController) Update(context *gin.Context) {
	var diaryUpdateDTO dto.DiaryUpdateDTO
	errDTO := context.ShouldBind(&diaryUpdateDTO)
	if errDTO != nil {
		res := helper.BuildErrorResponse("Failed to process request ", errDTO.Error(), helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, res)
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
	userID := fmt.Sprintf("%v", claims["user_id"])

	if c.diaryService.IsAllowedToEdit(userID, diaryUpdateDTO.ID) {
		id, errID := strconv.ParseUint(userID, 10, 64)
		if errID == nil {
			diaryUpdateDTO.UserID = id
		}
		result := c.diaryService.Update(diaryUpdateDTO)
		response := helper.BuildResponse(true, "OK", result)
		context.JSON(http.StatusOK, response)

	} else {

		response := helper.BuildErrorResponse("Updating other person diaries is not permitted", "You are not the owner of this diary", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, response)
	}

}

func (c *diaryController) Delete(context *gin.Context) {
	var diary entity.Diary
	id, err := strconv.ParseUint(context.Param("id"), 0, 0)
	if err != nil {
		response := helper.BuildErrorResponse("Diary ID parameter not found !", "No param id found !", helper.EmptyObj{})
		context.JSON(http.StatusBadRequest, response)
	}
	diary.ID = id

	authHeader := context.GetHeader("Authorization")

	splitToken := strings.Split(authHeader, "Bearer ")
	authHeader = splitToken[1]
	token, errToken := c.jwtService.ValidateToken(authHeader)
	if errToken != nil {
		panic(errToken.Error())
	}
	claims := token.Claims.(jwt.MapClaims)
	userID := fmt.Sprintf("%v", claims["user_id"])
	if c.diaryService.IsAllowedToEdit(userID, diary.ID) {
		c.diaryService.Delete(diary)
		res := helper.BuildResponse(true, "Deleted", helper.EmptyObj{})
		context.JSON(http.StatusOK, res)
	} else {
		res := helper.BuildErrorResponse("Deleting other person diaries is not permitted", "You are not the owner of this diary", helper.EmptyObj{})
		context.JSON(http.StatusForbidden, res)
	}

}

func (c *diaryController) getUserIDByToken(token string) string {
	givenToken, err := c.jwtService.ValidateToken(token)
	if err != nil {
		panic(err.Error())
	}
	claims := givenToken.Claims.(jwt.MapClaims)
	id := fmt.Sprintf("%v", claims["user_id"])
	return id
}
