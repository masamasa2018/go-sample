package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// AsyncHandlerは/asyncエンドポイントのリクエストを処理するハンドラー関数
func AsyncHandler(c *gin.Context) {
	// チャネルの初期化
	resultChan := make(chan string, 2)

	// 2つのタスクを並行して実行するためにゴルーチンを起動
	go func() {
		resultChan <- task1()
	}()

	go func() {
		resultChan <- task2()
	}()

	// 並行処理の完了を待つ
	result1 := <-resultChan
	result2 := <-resultChan

	// レスポンスとして結果を返す
	c.JSON(http.StatusOK, gin.H{
		"task1": result1,
		"task2": result2,
	})
}

// タスク1のシミュレーション（時間のかかる処理）
func task1() string {
	time.Sleep(2 * time.Second) // 2秒間待機
	return "Task 1 completed!"
}

// タスク2のシミュレーション（時間のかかる処理）
func task2() string {
	time.Sleep(3 * time.Second) // 3秒間待機
	return "Task 2 completed!"
}
