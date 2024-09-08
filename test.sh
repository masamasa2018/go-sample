#!/bin/bash

# プロジェクトのルートディレクトリに移動
cd "$(dirname "$0")"

# すべての*Handlerテストを実行
go test -v ./test -run 'Test.*Handler'

# テスト結果の確認
if [ $? -eq 0 ]; then
    echo "すべてのHandlerテストが成功しました。"
else
    echo "一部のHandlerテストが失敗しました。"
fi