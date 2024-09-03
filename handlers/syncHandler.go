package handlers

import (
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

func SyncHandler(c *gin.Context) {
	var wg sync.WaitGroup
	results := make(chan string, 2)

	// タスク1をゴルーチンで実行
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(2 * time.Second)
		task1 := "Task 1 complete"
		results <- task1
	}()

	// タスク2をゴルーチンで実行
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(1 * time.Second)
		task2 := "Task 2 complete"
		results <- task2
	}()

	// すべてのタスクが完了するのを待つ
	wg.Wait()
	close(results)

	// 結果を収集
	var task1Result, task2Result string
	for result := range results {
		if task1Result == "" {
			task1Result = result
		} else {
			task2Result = result
		}
	}

	// クライアントにJSONレスポンスを返す
	c.JSON(http.StatusOK, gin.H{
		"task1": task1Result,
		"task2": task2Result,
	})
}
