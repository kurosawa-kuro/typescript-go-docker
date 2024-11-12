package main

import (
	"backend/src/handler" // ハンドラーをインポート

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// CORSの設定
	r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Content-Type"},
	}))

	// ハンドラーを登録
	r.GET("/ping", handler.PingHandler)

	// ... 他の設定
	r.Run(":8000")
}
