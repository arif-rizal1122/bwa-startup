package main

import (

	"github.com/arif-rizal1122/bwa-startup/auth"
	"github.com/arif-rizal1122/bwa-startup/handler"
	"github.com/arif-rizal1122/bwa-startup/helper"
	"github.com/arif-rizal1122/bwa-startup/user"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa-api?charset=utf8mb4&parseTime=True&loc=Local"

	// Menggunakan driver MySQL dari GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.IfError(err)

	//
	userRepository := user.NewRepositoryUser(db)
	userService    := user.NewServiceUser(userRepository)
	authService := auth.NewJWTservice()


	userHandler := handler.NewUserHandler(userService, authService)
	router := gin.Default()

	api := router.Group("/api/v1/")  
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email_checker", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)
	

	router.Run()


}





// input dari user
// handler mapping input user menjadi struct input
// service melakukan mapping dari struct Input ke struct user
// repository save 
// db