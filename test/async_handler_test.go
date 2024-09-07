package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-sample/handlers"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAsyncHandler(t *testing.T) {
	// Ginのテストモードを設定
	gin.SetMode(gin.TestMode)

	// ルーターをセットアップ
	r := gin.New()
	r.GET("/async", handlers.AsyncHandler)

	// テストリクエストを作成
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/async", nil)

	// ハンドラーを実行
	r.ServeHTTP(w, req)

	// ステータスコードを確認
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスボディをパース
	var response map[string][]string // "results" は string の配列
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// レスポンスの内容を確認
	assert.Contains(t, response, "results")
	assert.Len(t, response["results"], 3) // 3つの結果が返ってくることを確認

	// 各タスクが完了していることを確認
	expectedResults := []string{"Task 1 completed!", "Task 2 completed!", "Task 3 completed!"}
	assert.ElementsMatch(t, expectedResults, response["results"])
}
