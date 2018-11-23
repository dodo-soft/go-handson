package main

import (
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 3 {
		os.Exit(1)
	}
	num1, err := strconv.Atoi(os.Args[1])
	// エラー発生(err != nil)したら、exit status 1で終了
	// num2もどうようにintに変換
	// 足した結果をfmt.Println()で出力
}
