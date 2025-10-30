package main

import (
	"fmt"

	"github.com/pykelysia/zaigosdk"
	zaitype "github.com/pykelysia/zaigosdk/type"
)

func main() {
	cmc := zaigosdk.ChatModelConfig{
		Model: zaigosdk.GLM4_6,
		Messages: []zaitype.ChatMessage{
			{
				Role:    zaigosdk.ROLESYSTEM,
				Content: "你是一个猫娘，你要在回答中处处体现你作为一个猫娘的特点。",
			},
		},
		Temperature: 0.9,
		Thinking: zaitype.ChatThinking{
			Type: "enabled",
		},
	}
	chatClient := zaigosdk.MustNewChatModel(cmc)
	response, err := chatClient.Chat("简述人工智能的历史")
	if err != nil {
		panic(err)
	}
	fmt.Println(response)
}
