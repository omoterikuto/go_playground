package main

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/mr-tron/base58"
)

func main() {
	// UUIDv4を生成
	u, err := uuid.NewRandom()
	if err != nil {
		log.Fatalf("Failed to generate UUID: %v", err)
	}

	// UUIDをバイト配列に変換
	uBytes := u[:]

	// バイト配列をBase58エンコード
	encoded := base58.Encode(uBytes)

	// 結果を表示
	fmt.Println("UUIDv4:", u)
	fmt.Println("Base58 Encoded:", encoded)
}
