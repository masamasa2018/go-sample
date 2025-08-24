package main

import (
	"fmt"
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

	// 問題のあるコードのテスト用エンドポイント
	router.GET("/problematic", handlers.ProblematicHandler)
	router.GET("/deep-nesting", handlers.DeepNestingHandler)
	router.GET("/duplicate1", handlers.DuplicatedCodeHandler1)
	router.GET("/duplicate2", handlers.DuplicatedCodeHandler2)

	// サーバー起動前のメッセージ
	fmt.Println("🚀 サーバーを http://localhost:8080 で起動します...")

	// サーバーの起動
	router.Run(":8080")
}
