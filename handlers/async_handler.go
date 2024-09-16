package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// AsyncHandlerは/asyncエンドポイントのリクエストを処理するハンドラー関数
func AsyncHandler(c *gin.Context) {
	// チャネルの初期化
	resultChan := make(chan string, 3)
	errorChan := make(chan error, 3)

	// 3つのタスクを並行して実行するためにゴルーチンを起動
	go func() {
		result, err := task1()
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- result
	}()

	go func() {
		result, err := task2()
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- result
	}()

	go func() {
		result, err := task3()
		if err != nil {
			errorChan <- err
			return
		}
		resultChan <- result
	}()

	// 結果の収集
	var results []string
	for i := 0; i < 3; i++ {
		select {
		case result := <-resultChan:
			results = append(results, result)
		case err := <-errorChan:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}

	// レスポンスとして結果を返す
	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}

// タスク1のシミュレーション
func task1() (string, error) {
	time.Sleep(2 * time.Second) // 2秒間待機
	return "Task 1 completed!", nil
}

// タスク2のシミュレーション
func task2() (string, error) {
	time.Sleep(3 * time.Second) // 3秒間待機
	return "Task 2 completed!", nil
}

// タスク3のシミュレーション
func task3() (string, error) {
	time.Sleep(1 * time.Second) // 1秒間待機
	return "Task 3 completed!", nil
}
