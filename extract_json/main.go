package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// Group represents a single group in the JSON structure
type Group struct {
	Name string `json:"name"`
}

// Data represents the entire JSON structure
type Data struct {
	Groups []Group `json:"groups"`
}

func main() {
	// JSONファイルのパスを指定
	filePath := "input.json"

	// ファイルを読み込む
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("ファイルを開くことができませんでした: %v", err)
	}
	defer file.Close()

	// ファイルの内容を読み込む
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatalf("ファイルを読み込むことができませんでした: %v", err)
	}

	// JSONデータをData構造体にデコード
	var data Data
	err = json.Unmarshal(byteValue, &data)
	if err != nil {
		log.Fatalf("JSONデコードに失敗しました: %v", err)
	}

	// groups内のname一覧を改行区切りで出力（dev, stgを含むものは除外）
	for _, group := range data.Groups {
		if !strings.Contains(group.Name, "dev") && !strings.Contains(group.Name, "stg") {
			// 先頭の〇〇-と末尾の_prdを除いたものを出力
			name := group.Name

			// 末尾の_prdを除去
			if strings.HasSuffix(name, "_prd") {
				name = strings.TrimSuffix(name, "_prd")
			}

			fmt.Println(name)
		}
	}
}
