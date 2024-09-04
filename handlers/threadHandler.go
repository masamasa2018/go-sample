package handlers

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// ThreadHandlerは/threadエンドポイントのリクエストを処理するハンドラー関数
func ThreadHandler(c *gin.Context) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	results := make([]string, 0)

	// 3つのタスクを並行して実行する
	wg.Add(3)

	go func() {
		defer wg.Done()
		result := threadTask1()
		mu.Lock()
		results = append(results, result)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		result := threadTask2()
		mu.Lock()
		results = append(results, result)
		mu.Unlock()
	}()

	go func() {
		defer wg.Done()
		result := threadTask3()
		mu.Lock()
		results = append(results, result)
		mu.Unlock()
	}()

	// 全てのゴルーチンが完了するまで待機
	wg.Wait()

	// レスポンスとして結果を返す
	c.JSON(http.StatusOK, gin.H{
		"results": results,
	})
}

// タスク1のシミュレーション（時間のかかる処理）
func threadTask1() string {
	time.Sleep(2 * time.Second) // 2秒間待機
	return "Task 1 completed!"
}

// タスク2のシミュレーション（時間のかかる処理）
func threadTask2() string {
	time.Sleep(3 * time.Second) // 3秒間待機
	return "Task 2 completed!"
}

// タスク3のシミュレーション（時間のかかる処理）
func threadTask3() string {
	time.Sleep(1 * time.Second) // 1秒間待機
	return "Task 3 completed!"
}
