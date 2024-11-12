package router

import (
	"backend/src/handler"
	"backend/src/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Setup initializes the router and its routes
func Setup() *gin.Engine {
	r := gin.Default()

	// ミドルウェアの設定
	r.Use(middleware.CORS())

	// Swaggerの設定
	setupSwagger(r)

	// ルートの設定
	setupRoutes(r)

	return r
}

// setupSwagger configures swagger documentation
func setupSwagger(r *gin.Engine) {
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}

// setupRoutes configures all application routes
func setupRoutes(r *gin.Engine) {
	r.GET("/ping", handler.PingHandler)
	// 他のルートをここに追加
}
