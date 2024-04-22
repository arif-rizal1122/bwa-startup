package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Errors(c *gin.Context, err error)  {
	var errors []string

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		response := APIResponse("account register failed", http.StatusBadRequest, "error", "Validation error")
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	for _, e := range validationErrors {
		errors = append(errors, e.Error())
	}

	errorMessage := gin.H{"errors": errors}

	response := APIResponse("account register failed", http.StatusBadRequest, "error", errorMessage)
	c.JSON(http.StatusUnprocessableEntity, response)
	return
}
