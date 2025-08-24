package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ProblematicHandlerは意図的に問題のあるコードを含むハンドラー
func ProblematicHandler(c *gin.Context) {
	// 未使用の変数
	unusedVar := "this variable is not used"
	anotherUnusedVar := 42
	yetAnotherUnused := make([]string, 0)

	// 冗長な変数宣言
	var result string
	result = ""
	result = result + "Starting process..."

	// 非効率的な文字列結合
	output := ""
	for i := 0; i < 100; i++ {
		output = output + "Item " + strconv.Itoa(i) + " "
	}

	// 複雑で冗長な条件分岐
	param := c.Query("type")
	var responseMessage string
	if param == "type1" {
		responseMessage = processType1()
	} else if param == "type2" {
		responseMessage = processType2()
	} else if param == "type3" {
		responseMessage = processType3()
	} else if param == "type4" {
		responseMessage = processType4()
	} else if param == "type5" {
		responseMessage = processType5()
	} else {
		responseMessage = "Unknown type"
	}

	// 不要な変数代入
	temp1 := responseMessage
	temp2 := temp1
	finalMessage := temp2

	// 適切でないエラーハンドリング
	data, _ := processData()  // エラーを無視

	c.JSON(http.StatusOK, gin.H{
		"message": finalMessage,
		"data":    data,
		"output":  output,
	})
}

// 冗長で複雑すぎる関数
func processType1() string {
	// 不要なループ
	for i := 0; i < 1; i++ {
		for j := 0; j < 1; j++ {
			// 単純な処理を複雑に書く
			var result string
			if true {
				if true {
					result = "Type 1 processed"
				}
			}
			return result
		}
	}
	return ""
}

func processType2() string {
	// 非効率的な文字列操作
	base := "Type"
	space := " "
	number := "2"
	action := "processed"
	
	result := base + space + number + space + action
	result = strings.TrimSpace(result)
	result = strings.ReplaceAll(result, "  ", " ")
	
	return result
}

func processType3() string {
	// 不要な変数宣言とマジックナンバー
	magic := 42
	another := 123
	sum := magic + another
	
	if sum > 100 {
		return "Type 3 processed"
	} else {
		return "Type 3 failed"
	}
}

func processType4() string {
	// 過度に複雑な処理
	data := make(map[string]interface{})
	data["type"] = 4
	data["status"] = "processing"
	
	var result string
	if val, ok := data["type"]; ok {
		if intVal, ok := val.(int); ok {
			if intVal == 4 {
				if status, ok := data["status"]; ok {
					if statusStr, ok := status.(string); ok {
						if statusStr == "processing" {
							result = "Type 4 processed"
						}
					}
				}
			}
		}
	}
	
	if result == "" {
		result = "Type 4 failed"
	}
	
	return result
}

func processType5() string {
	// 不要な時間処理
	now := time.Now()
	future := now.Add(1 * time.Second)
	
	// 意味のない比較
	if future.After(now) {
		return "Type 5 processed"
	} else {
		return "Type 5 failed"  // 実際には到達しないコード
	}
}

// 適切でないエラーハンドリング
func processData() ([]string, error) {
	// エラーが発生する可能性があるが、適切に処理されていない
	data := make([]string, 0)
	
	// 意図的に複雑で非効率的な処理
	for i := 0; i < 10; i++ {
		value := fmt.Sprintf("data-%d", i)
		data = append(data, value)
		
		// 無意味な条件
		if i%2 == 0 {
			if i%4 == 0 {
				if i%8 == 0 {
					// 何もしない
					continue
				}
			}
		}
	}
	
	return data, nil
}

// DeepNestingHandlerは深いネストを持つ問題のあるハンドラー
func DeepNestingHandler(c *gin.Context) {
	// 未使用のインポート要因となる変数
	timeNow := time.Now()
	_ = timeNow  // 使用していることにするための無意味な代入

	param1 := c.Query("level1")
	if param1 != "" {
		param2 := c.Query("level2")
		if param2 != "" {
			param3 := c.Query("level3")
			if param3 != "" {
				param4 := c.Query("level4")
				if param4 != "" {
					param5 := c.Query("level5")
					if param5 != "" {
						// 深すぎるネスト
						result := "Deep nesting completed"
						c.JSON(http.StatusOK, gin.H{"result": result})
						return
					} else {
						c.JSON(http.StatusBadRequest, gin.H{"error": "level5 required"})
						return
					}
				} else {
					c.JSON(http.StatusBadRequest, gin.H{"error": "level4 required"})
					return
				}
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "level3 required"})
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": "level2 required"})
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "level1 required"})
		return
	}
}

// DuplicatedCodeHandler1は重複したコードを含むハンドラー1
func DuplicatedCodeHandler1(c *gin.Context) {
	// 重複したロジック
	name := c.Query("name")
	if name == "" {
		name = "default"
	}
	
	email := c.Query("email")
	if email == "" {
		email = "default@example.com"
	}
	
	phone := c.Query("phone")
	if phone == "" {
		phone = "000-0000-0000"
	}
	
	c.JSON(http.StatusOK, gin.H{
		"name":  name,
		"email": email,
		"phone": phone,
		"type":  "handler1",
	})
}

// DuplicatedCodeHandler2は重複したコードを含むハンドラー2
func DuplicatedCodeHandler2(c *gin.Context) {
	// 重複したロジック（上と同じ）
	name := c.Query("name")
	if name == "" {
		name = "default"
	}
	
	email := c.Query("email")
	if email == "" {
		email = "default@example.com"
	}
	
	phone := c.Query("phone")
	if phone == "" {
		phone = "000-0000-0000"
	}
	
	c.JSON(http.StatusOK, gin.H{
		"name":  name,
		"email": email,
		"phone": phone,
		"type":  "handler2",
	})
}
