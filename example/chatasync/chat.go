package main

import (
	"fmt"
	"time"

	"github.com/pykelysia/zaigosdk"
)

func main() {
	chatClient := zaigosdk.MustDefaultChatModel()
	response, err := chatClient.ChatAsync("任意回答事或否")
	if err != nil {
		panic(err)
	}
	time.Sleep(time.Minute)
	gets, err := zaigosdk.CheckResult(response.ID)

	fmt.Println(response)
	fmt.Println(gets)
}
