package api

import (
	"golay/internal/domain/user/handler"

	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	handler *handler.UserHandler
}

func NewUserAPI(service *handler.UserHandler) *UserAPI {
	return &UserAPI{handler: service}
}

func (api *UserAPI) SetupRoutes(router *gin.Engine) {
	router.GET("/user/:id", api.handler.GetUser)
	router.POST("/user", api.handler.CreateUser)
	router.PUT("/user/:id", api.handler.UpdateUser)
	router.DELETE("/user/:id", api.handler.DeleteUser)
}
