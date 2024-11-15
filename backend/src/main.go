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
	// データベースの初期化
	db := config.SetupDB()

	// Swaggerの初期化
	config.SetupSwagger()

	// ルーターの設定
	r := router.Setup(db)

	// サーバーの起動
	r.Run(":8000")
}
