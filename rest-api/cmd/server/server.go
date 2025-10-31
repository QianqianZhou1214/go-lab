package server

import (
	"context"
	"net/http"
	"rest-api/x/interfacesx"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type GinServer interface {
	Start(ctx context.Context, httpAddress string) error
	Shutdown(ctx context.Context) error
	RegisterRoute(method, path string, handler gin.HandlerFunc)
	RegisterGroupRoute(path string, routes []interfacesx.RouteDefinition, middlewares ...gin.HandlerFunc)
}

type GinServerBuilder struct {
}

// This represents the Gin server
type ginServer struct {
	engine *gin.Engine // The engine for server
	server *http.Server
}

// NewGinServerBuilder To initialize Gin server
func NewGinServerBuilder() *GinServerBuilder {
	return &GinServerBuilder{}
}

// Build Gin serve and return the instance with created engine
func (b *GinServerBuilder) Build() GinServer {
	engine := gin.Default()
	return &ginServer{engine: engine}
}

// Start To start server
func (gs *ginServer) Start(ctx context.Context, httpAddress string) error {
	gs.server = &http.Server{
		Addr:    httpAddress,
		Handler: gs.engine,
	}
	go func() {
		if err := gs.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listening: %s\n", err)
		}
	}()
	logrus.Infof("Https server is running on port %s", httpAddress)
	return nil
}

func (gs *ginServer) Shutdown(ctx context.Context) error {
	if err := gs.server.Shutdown(ctx); err != nil {
		logrus.Fatalf("Server shutdown %s", err)
	}
	logrus.Info("Server exited")
	return nil
}

// RegisterRoute Method to register a single route
func (gs *ginServer) RegisterRoute(method, path string, handler gin.HandlerFunc) {
	switch method {
	case "GET":
		gs.engine.GET(path, handler)
	case "POST":
		gs.engine.POST(path, handler)
	case "PUT":
		gs.engine.PUT(path, handler)
	case "DELETE":
		gs.engine.DELETE(path, handler)
	case "PATCH":
		gs.engine.PATCH(path, handler)
	default:
		logrus.Errorf("Invalid https method")
	}
}

//api/user/register
//api/user/signup

func (gs *ginServer) RegisterGroupRoute(path string, routes []interfacesx.RouteDefinition, middlewares ...gin.HandlerFunc) {
	group := gs.engine.Group(path)
	group.Use(middlewares...)
	for _, route := range routes {
		switch route.Method {
		case "GET":
			group.GET(route.Path, route.Handler)
		case "POST":
			group.POST(route.Path, route.Handler)
		case "PUT":
			group.PUT(route.Path, route.Handler)
		case "DELETE":
			group.DELETE(route.Path, route.Handler)
		case "PATCH":
			group.PATCH(route.Path, route.Handler)
		default:
			logrus.Errorf("Invalid https method")
		}
	}
}
