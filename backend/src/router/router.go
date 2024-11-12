package router

import (
	"backend/src/handler"
	"backend/src/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Setup initializes the router and its routes
func Setup(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// ミドルウェアの設定
	r.Use(middleware.CORS())

	// ハンドラーの初期化
	micropostHandler := handler.NewMicropostHandler(db)

	// ルートの設定
	r.GET("/ping", handler.PingHandler)

	microposts := r.Group("/microposts")
	{
		microposts.POST("", micropostHandler.Create)
		microposts.GET("", micropostHandler.FindAll)
	}

	return r
}
