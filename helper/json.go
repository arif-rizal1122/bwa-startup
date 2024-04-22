package helper

import (

	"github.com/arif-rizal1122/bwa-startup/user"
	"github.com/go-playground/validator/v10"
)

type ResponseJSON struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func APIResponse(Message string, Code int, Status string, Data interface{}) ResponseJSON {
	Meta := Meta{
		Message: Message,
		Code:    Code,
		Status:  Status,
	}

	response := ResponseJSON{
		Meta: Meta,
		Data: Data,
	}

	return response
}

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user user.User, token string) UserFormatter {
    formatter := UserFormatter{
		ID: user.ID,
		Name: user.Name,
		Occupation: user.Occupation,
		Email: user.Email,
		Token: token,
	}

	return formatter
}


func FormatValidationError(err error) []string {

	var errors []string

	for _, e := range err.(validator.ValidationErrors){
		errors = append(errors, e.Error())
	}

	return errors

}
