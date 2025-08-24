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
	// 未使用の変数
	unusedString := "not used"
	unusedInt := 999
	unusedSlice := make([]int, 0)
	
	// 値の初期化
	value := 10

	// 冗長なポインタ操作
	ptr := &value
	ptrToPtr := &ptr
	actualPtr := *ptrToPtr

	// ポインタを使って初期値を出力
	initialValue := *actualPtr

	// ポインタを渡して関数で値を変更
	incrementValue(actualPtr)
	incrementedValue := *actualPtr

	// さらに値を変更
	incrementValue(actualPtr)
	finalValue := *actualPtr

	// レスポンスとして変更後の値を返す
	c.JSON(http.StatusOK, gin.H{
		"initial_value":     initialValue,
		"incremented_value": incrementedValue,
		"final_value":       finalValue,
	})
}
