package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	client := http.Client{
		Timeout: 5 * time.Second, // 设置超时时间
	}
	resp, err := client.Get("https://www.baidu.com")
	if err != nil {
		fmt.Printf("Error making HTTP request: %s\n", err)
		return
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Failed to read response body:", err)
		return
	}
	fmt.Println(string(body))
	fmt.Println("Response status:", resp.Status)
}
