package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/arif-rizal1122/bwa-startup/auth"
	"github.com/arif-rizal1122/bwa-startup/campaign"
	"github.com/arif-rizal1122/bwa-startup/handler"
	"github.com/arif-rizal1122/bwa-startup/helper"
	"github.com/arif-rizal1122/bwa-startup/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/bwa-api?charset=utf8mb4&parseTime=True&loc=Local"

	// Menggunakan driver MySQL dari GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.IfError(err)

	// users
	userRepository := user.NewRepositoryUser(db)
	userService    := user.NewServiceUser(userRepository)
	authService := auth.NewJWTservice()
	userHandler := handler.NewUserHandler(userService, authService)
	// users handler

	// campaigns
	campaignRepository := campaign.NewRepository(db)
    campaignService    := campaign.NewService(campaignRepository)

	camp, _ := campaignService.FindCampaings(0)
	fmt.Println(len(camp)) 

	router := gin.Default()
	api := router.Group("/api/v1/")  
	api.POST("/register", userHandler.RegisterUser)
	api.POST("/login", userHandler.LoginUser)
	api.POST("/email_checker", userHandler.CheckEmailAvailability)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	



	router.Run()
}




func authMiddleware(authService auth.Service, userService user.ServiceUser) gin.HandlerFunc {
	return func (c *gin.Context)  {
		authHeader := c.GetHeader("Authorization")
	
		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return 
		}
	
		// bearer token
		tokenString := " "
		arrayToken := strings.Split(authHeader, " ")
		// dikasih dua karena ada 2 nilai nanti yaitu bearer[0] dan token[1]
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}
	
		// disini cek dengan token
		token, err := authService.ValidationToken(tokenString)
		if err != nil {
			response := helper.APIResponse("unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return 
		}

		claim, ok := token.Claims.(jwt.MapClaims)

		if !ok || !token.Valid {
			response := helper.APIResponse("unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return 
		}

		userID := int(claim["user_id"].(float64))
		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("unathorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return 
		}

		c.Set("currentUser", user)
	}
}

