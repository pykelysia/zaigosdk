package main

import (
	"fmt"

	"github.com/pykelysia/zaigosdk"
)

func main() {
	chatClient := zaigosdk.MustDefaultChatModel()
	response := chatClient.Chat("任意回答事或否")
	fmt.Println(response)
}
