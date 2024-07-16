package main

import (
	"encoding/json"
	"fmt"
)

type Rates struct {
	Base   string `json:"base currency"`
	Symbol string `json:"destination currency"`
}

func main() {
	jsonString := `
        {
            "success": true,
            "timestamp": 1588779306,
            "base": "USD",
            "date": "2022-01-15",
            "rates": {
                "BNB": 0.00225,
                "BTC": 0.000020,
                "EUR": 0.879,
                "GBP": 0.733,
                "CNY": 6.36
          } 
        }`
	// 直接定义整个与json对应的结构体的话，扩展性不好，比如再想增加新的目的国家汇率就需要修改结构体的定义
	// 因为json本质每一个字段都是一个key:value，所以用一个value为interface{}类型的map存储
	// 这样做，当遇到嵌套结构的json时，里面嵌套的也会解析为map
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonString), &result)
	if err == nil {
		fmt.Println(result["success"])
		rates := result["rates"]
		m, ok := rates.(map[string]interface{})
		if ok {
			for key, value := range m {
				fmt.Printf("%s:%v\n", key, value)
			}
		} else {
			fmt.Println(rates)
		}

	} else {
		fmt.Println(err)
	}
}
