package handler

import (
	"fmt"
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
	// tangkap input dari user
	// simpan gambarnya di folder "images/"
	// di service kita panggil repo
	// jwt (sementara hardcode dulu)
	// repo ambil data user yg ID
	// repo update data user simpan lokasi file

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}

		response := helper.APIResponse("failed to upload image", http.StatusBadRequest, "errors", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userID := 76
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