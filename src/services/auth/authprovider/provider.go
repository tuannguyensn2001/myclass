package authprovider

import (
	"github.com/gin-gonic/gin"
	"myclass/src/services/auth/authservice"
	"myclass/src/services/auth/authtransport"
	"myclass/src/services/user/userrepository"
	"myclass/src/services/user/userservice"
)

func Boot(r *gin.Engine) error {
	service := authservice.New()

	userRepository := userrepository.New()
	userService := userservice.New(userRepository)

	service.SetUserService(userService)

	err := service.CheckHealth()
	if err != nil {
		return err
	}

	transport := authtransport.NewHttpTransport(userService)

	r.POST("/api/v1/auth/register", transport.Register)

	return nil
}
