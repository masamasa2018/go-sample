package main

import (
	"gin-goroutine/handlers" // handlersパッケージをインポート

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのデフォルトのルーターを作成
	router := gin.Default()
	// /async エンドポイントを定義
	router.GET("/async", handlers.AsyncHandler)
	// /sync エンドポイントを定義
	router.GET("/sync", handlers.SyncHandler)

	// サーバーの起動
	router.Run(":8080")
}
