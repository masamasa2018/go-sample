package test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go-sample/handlers"
)

func TestThreadHandler(t *testing.T) {
	// Ginのテスト用モードを設定
	gin.SetMode(gin.TestMode)

	// ルーターを設定
	router := gin.New()
	router.GET("/thread", handlers.ThreadHandler)

	// リクエストを作成
	req, _ := http.NewRequest("GET", "/thread", nil)

	// レスポンスレコーダーを作成
	w := httptest.NewRecorder()

	// リクエストを実行
	start := time.Now()
	router.ServeHTTP(w, req)
	duration := time.Since(start)

	// レスポンスを検証
	assert.Equal(t, http.StatusOK, w.Code)

	// レスポンスボディをパース
	var response map[string][]string
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)

	// 結果を検証
	results := response["results"]
	assert.Len(t, results, 3)
	assert.Contains(t, results, "Task 1 completed!")
	assert.Contains(t, results, "Task 2 completed!")
	assert.Contains(t, results, "Task 3 completed!")

	// 実行時間を検証（3秒以上、4秒未満であることを確認）
	assert.GreaterOrEqual(t, duration, 3*time.Second)
	assert.Less(t, duration, 4*time.Second)
}