package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 値をインクリメントする関数
func incrementValue(ptr *int) {
	*ptr = *ptr + 1
}

// ポインタと関数を使って値を変更するハンドラー
func PointerHandler(c *gin.Context) {
	// 値の初期化
	value := 10

	// ポインタの取得
	ptr := &value

	// ポインタを使って初期値を出力
	initialValue := *ptr

	// ポインタを渡して関数で値を変更
	incrementValue(ptr)
	incrementedValue := *ptr

	// さらに値を変更
	incrementValue(ptr)
	finalValue := *ptr

	// レスポンスとして変更後の値を返す
	c.JSON(http.StatusOK, gin.H{
		"initial_value":     initialValue,
		"incremented_value": incrementedValue,
		"final_value":       finalValue,
	})
}
