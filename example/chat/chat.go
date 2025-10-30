package main

import (
	"fmt"

	"github.com/pykelysia/zaigosdk"
)

func main() {
	chatClient := zaigosdk.MustDefaultChatModel()
	response, err := chatClient.Chat("任意回答事或否")
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
