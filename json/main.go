package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

// camelCase → snake_case 変換
func toSnakeCase(str string) string {
	// 大文字の前にアンダースコアを入れる
	re := regexp.MustCompile(`([a-z0-9])([A-Z])`)
	snake := re.ReplaceAllString(str, "${1}_${2}")
	return strings.ToLower(snake)
}

// JSONのキーを再帰的にsnake_caseへ変換
func convertKeysToSnakeCase(m map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{}, len(m))
	for k, v := range m {
		newKey := toSnakeCase(k)
		switch vv := v.(type) {
		case map[string]interface{}:
			newMap[newKey] = convertKeysToSnakeCase(vv)
		case []interface{}:
			// 配列の中もmapなら再帰処理
			arr := make([]interface{}, len(vv))
			for i, elem := range vv {
				if subMap, ok := elem.(map[string]interface{}); ok {
					arr[i] = convertKeysToSnakeCase(subMap)
				} else {
					arr[i] = elem
				}
			}
			newMap[newKey] = arr
		default:
			newMap[newKey] = v
		}
	}
	return newMap
}

func main() {
	var hogeMap map[string]interface{}
	if err := json.Unmarshal([]byte(jsonA), &hogeMap); err != nil {
		panic(err)
	}
	fmt.Println(hogeMap["profile"].(map[string]interface{})["firstName"].(string))

	// キーを snake_case に変換
	snakeMap := convertKeysToSnakeCase(hogeMap)

	// 整形して出力
	pretty, _ := json.MarshalIndent(snakeMap, "", "  ")
	fmt.Println(string(pretty))
}

var jsonA = `
{
  "userId": 12345,
  "profile": {
    "firstName": "Taro",
    "lastName": "Yamada"
  },
  "favoriteFoods": [
    "sushi",
    "ramen"
  ],
  "recentOrders": [
    {
      "orderId": "A1001",
      "orderDate": "2025-09-04",
      "items": [
        { "itemName": "salmonSushi", "quantity": 8, "price": 12.5 },
        { "itemName": "misoSoup", "quantity": 2, "price": 3.0 }
      ]
    },
    {
      "orderId": "A1002",
      "orderDate": "2025-09-02",
      "items": [
        { "itemName": "ramenShoyu", "quantity": 1, "price": 9.5 }
      ]
    }
  ]
}
`

var jsonB = `
{
  "favorite_foods": [
    "sushi",
    "ramen"
  ],
  "profile": {
    "first_name": "Taro",
    "last_name": "Yamada"
  },
  "recent_orders": [
    {
      "items": [
        {
          "item_name": "salmonSushi",
          "price": 12.5,
          "quantity": 8
        },
        {
          "item_name": "misoSoup",
          "price": 3,
          "quantity": 2
        }
      ],
      "order_date": "2025-09-04",
      "order_id": "A1001"
    },
    {
      "items": [
        {
          "item_name": "ramenShoyu",
          "price": 9.5,
          "quantity": 1
        }
      ],
      "order_date": "2025-09-02",
      "order_id": "A1002"
    }
  ],
  "user_id": 12345
}
`
