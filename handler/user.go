package handler

import (
	"net/http"

	"github.com/arif-rizal1122/bwa-startup/helper"
	"github.com/arif-rizal1122/bwa-startup/user"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.ServiceUser
}

func NewUserHandler(userService user.ServiceUser) *userHandler {
	return &userHandler{userService: userService}
}




func (h *userHandler) RegisterUser(c *gin.Context) {
	// Ambil input dari user
	var input user.RegisterUserInput



	 err := c.ShouldBindJSON(&input)
	 if err != nil {
        errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.APIResponse("account register failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	 }



	user, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("account register failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}



	// token nanti yh

	formater := helper.FormatUser(user, "")
	response := helper.APIResponse("account has been registered", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)
}



func (h *userHandler) LoginUser(c *gin.Context) {
	// user memasukan input (email dan password)
	// input ditangkap handler
	// mapping dari input user ke input struct
	// input struct passing ke service
	// di service mencari dgn bantuan repository user dengan email x'
	// mencocokan password
	var input user.LoginInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("login success", http.StatusUnprocessableEntity, "errors", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	logginUser, err := h.userService.LoginUser(input)

	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("login success", http.StatusUnprocessableEntity, "errors", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}


	formater := helper.FormatUser(logginUser, "")

	response := helper.APIResponse("login success", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)
	





}