
* 起動方法
go run main.go

http://localhost:8080
にアクセス

* コードの説明
・スレッドのサンプル
http://localhost:8080/thread

　処理
　　handlers/threadHandler.go

・ゴルーチンのサンプル
http://localhost:8080/async

　処理
　　handlers/asyncHandler.go

・ポインタのサンプル
http://localhost:8080/pointer

　処理
　　handlers/pointerHandler.go


* テストの実行方法
chmod +x test.sh

./test.sh