package server

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

type GinServer interface {
	Start(ctx context.Context, httpAddress string) error
	Shutdown(ctx context.Context) error
}

type GinServerBuilder struct {
}

// This represents the Gin server
type ginServer struct {
	engine *gin.Engine // The engine for server
	server *http.Server
}

// To initialize Gin server
func NewGinServerBuilder() *GinServerBuilder {
	return &GinServerBuilder{}
}

// Build Gin serve and return the instance with created engine
func (b *GinServerBuilder) Build() GinServer {
	engine := gin.Default()
	return &ginServer{engine: engine}
}

// To start server
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
