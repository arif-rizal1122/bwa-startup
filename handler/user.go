package handler

import (
	"fmt"
	"net/http"

	"github.com/arif-rizal1122/bwa-startup/auth"
	"github.com/arif-rizal1122/bwa-startup/helper"
	"github.com/arif-rizal1122/bwa-startup/user"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.ServiceUser
	authService auth.Service
}

func NewUserHandler(userService user.ServiceUser, authService auth.Service) *userHandler {
	return &userHandler{
		userService: userService,
		authService: authService,
	}
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



	token, err := h.authService.GenerateToken(user.ID)
	if err != nil {
		response := helper.APIResponse("account register failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formater := helper.FormatUser(user, token)
	response := helper.APIResponse("account has been registered", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)
}



func (h *userHandler) LoginUser(c *gin.Context) {
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

	token, err := h.authService.GenerateToken(logginUser.ID)
	formater := helper.FormatUser(logginUser, token)
	if err != nil {
		response := helper.APIResponse("account login failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.APIResponse("login success", http.StatusOK, "success", formater)

	c.JSON(http.StatusOK, response)
}



func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// Inisialisasi variabel untuk menampung input email
	var input user.CheckEmailInput

	// Mengikat JSON dari request ke struct input
	if err := c.ShouldBindJSON(&input); err != nil {
		// Jika terjadi kesalahan saat binding JSON, tangani error validasi
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		// Buat respons API dengan pesan error validasi
		response := helper.APIResponse("email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)

		// Mengirimkan respons API ke client
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	// Memeriksa ketersediaan email menggunakan service
	isEmailAvailable, err := h.userService.IsEmailAvailable(input)
	if err != nil {
		// Jika terjadi kesalahan saat memeriksa email, tangani error server
		errorMessage := gin.H{"errors": "server error"}
		response := helper.APIResponse("email checking failed", http.StatusInternalServerError, "error", errorMessage)

		// Mengirimkan respons API ke client
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	// Menyiapkan data respons
	data := gin.H{
		"is_available": isEmailAvailable,
	}

	// Menentukan pesan meta berdasarkan ketersediaan email
	metaMessage := "email has been registered"
	if isEmailAvailable {
		metaMessage = "email is available"
	}

	// Buat respons API dengan pesan meta yang sudah ditentukan
	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusOK, response)
}




func (h *userHandler) UploadAvatar(c *gin.Context) {

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("failed to upload image", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	userID := currentUser.ID
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("failed to upload image", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("failed to upload image", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	
	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("avatar successfully uploaded", http.StatusOK, "errors", data)
	c.JSON(http.StatusOK, response)

}