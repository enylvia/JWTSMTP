package handler

import (
	"github.com/gin-gonic/gin"
	"jwtsmtp/entity"
	"jwtsmtp/helper"
	"jwtsmtp/service"
	"net/http"
)

type userHandler struct {
	userService service.Service
	jwtService	service.JWTService
}

func NewUserHandler(userService service.Service,jwtService	service.JWTService) *userHandler {
	return &userHandler{userService: userService, jwtService: jwtService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input entity.RegisterInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse(http.StatusUnprocessableEntity,"error","Failed to process request",errorMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	newUser, err := h.userService.Register(input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse(http.StatusBadRequest,"error","Failed to process request",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}
		response := helper.ApiResponse(http.StatusOK,"success","Successfully register",newUser)
		c.JSON(http.StatusOK,response)
}

func (h *userHandler) Login(c *gin.Context){
	var input entity.LoginInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse(http.StatusUnprocessableEntity,"error","Failed to process request",errorMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	user, err := h.userService.Login(input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse(http.StatusBadRequest,"error","Failed to login",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	validToken, err := h.jwtService.GenerateJWT(user.Username, int(user.Id))
	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse(http.StatusBadRequest,"error","Failed to generate token",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}
	user.Token = validToken
	response := helper.ApiResponse(http.StatusOK,"success","Successfully login",user)
	c.JSON(http.StatusOK,response)
}

func (h *userHandler)SendMail(c *gin.Context){
	var input entity.MailInput

	err := c.ShouldBindJSON(&input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse(http.StatusUnprocessableEntity,"error","Failed to process request",errorMessage)
		c.JSON(http.StatusUnprocessableEntity,response)
		return
	}

	mail, err := h.userService.SendMail(input)

	if err != nil {
		errorMessage := gin.H{"error": err.Error()}
		response := helper.ApiResponse(http.StatusBadRequest,"error","Failed to send mail",errorMessage)
		c.JSON(http.StatusBadRequest,response)
		return
	}

	response := helper.ApiResponse(http.StatusOK,"success","Successfully send mail",mail)
	c.JSON(http.StatusOK,response)
}