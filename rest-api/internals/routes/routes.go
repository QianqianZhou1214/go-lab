package routes

import (
	"rest-api/cmd/server"
	"rest-api/internals/handler"
	"rest-api/x/interfacesx"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func RegisterUserRoutes(server server.GinServer, userHandler *handler.UserHandler) {
	server.RegisterGroupRoute("api/v1/user", []interfacesx.RouteDefinition{
		{Method: "POST", Path: "/register", Handler: userHandler.CreateUser},
	}, func(c *gin.Context) {
		logrus.Infof("Request on %s", c.Request.URL.Path)
	})
}
