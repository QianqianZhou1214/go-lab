package handler

import (
	"net/http"
	"rest-api/internals/services"
	"rest-api/x/interfacesx"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type UserHandler struct {
	usrService services.UserService
	validate   *validator.Validate
}

func NewUserHandler(userService services.UserService) *UserHandler {
	return &UserHandler{
		usrService: userService,
		validate:   validator.New(),
	}
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var userRequest interfacesx.UserRegistrationRequest
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusBadRequest,
		})
		return
	}
	if err := h.validate.Struct(userRequest); err != nil {
		c.JSON(http.StatusBadRequest, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusBadRequest,
		})
		return
	}
	userData, err := h.usrService.CreateUserAccount(&userRequest)
	if err != nil {
		c.JSON(http.StatusInternalServerError, interfacesx.ErrorMessage{
			Message: err.Error(),
			Status:  interfacesx.StatusError,
			Code:    http.StatusBadRequest,
		})
		return
	}
	c.JSON(http.StatusCreated, interfacesx.UserResponse{
		Message: "User created successfully",
		Status:  interfacesx.StatusSuccess,
		Code:    http.StatusCreated,
		Data:    *userData,
	})
}
