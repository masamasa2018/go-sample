package main

import (
	"gin-goroutine/handlers" // handlersパッケージをインポート

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのデフォルトのルーターを作成
	router := gin.Default()

	router.GET("/async", handlers.AsyncHandler)

	router.GET("/thread", handlers.ThreadHandler)

	router.GET("/pointer", handlers.PointerHandler)

	// サーバーの起動
	router.Run(":8080")
}
