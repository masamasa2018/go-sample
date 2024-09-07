package main

import (
	"go-sample/handlers" // handlersパッケージをインポート
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Ginのデフォルトのルーターを作成
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World!")
	})

	router.GET("/async", handlers.AsyncHandler)

	router.GET("/thread", handlers.ThreadHandler)

	router.GET("/pointer", handlers.PointerHandler)

	// サーバーの起動
	router.Run(":8080")
}
