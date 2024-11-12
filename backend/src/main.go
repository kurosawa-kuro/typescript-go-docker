package main

import (
	"backend/src/config"
	"backend/src/router"
)

// @title           Your API Title
// @version         1.0
// @description     Your API Description
// @host           localhost:8000
// @BasePath       /
func main() {
	// Swaggerの初期化
	config.SetupSwagger()

	// ルーターの設定
	r := router.Setup()

	// サーバーの起動
	r.Run(":8000")
}
