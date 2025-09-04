package hoge

import (
	"fmt"
	"runtime"
	"strings"
)

func Hoge() {
	// skip=1: この関数Bの呼び出し元を取得
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		fmt.Println("呼び出し元を取得できませんでした")
		return
	}

	// 関数の完全修飾名（例: "main.A"）
	fn := runtime.FuncForPC(pc)
	fullName := fn.Name()

	// パッケージ名部分を抽出
	lastSlash := strings.LastIndex(fullName, "/")
	lastDot := strings.LastIndex(fullName, ".")
	if lastDot > lastSlash {
		pkgName := fullName[:lastDot]
		fmt.Printf("呼び出し元パッケージ: %s\n", pkgName)
	}

	fmt.Printf("呼び出し元関数: %s\n", fullName)
}
