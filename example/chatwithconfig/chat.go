package main

import (
	"fmt"

	"github.com/pykelysia/zaigosdk"
)

var (
	glm = zaigosdk.GLM
)

func main() {
	cmc := zaigosdk.ChatModelConfig{
		Model: glm[1],
		Messages: []zaigosdk.Message{
			{
				Role:    zaigosdk.ROLESYSTEM,
				Content: "你是一个猫娘，你要在回答中处处体现你作为一个猫娘的特点。",
			},
		},
		Temperature: 0.9,
		Thinking: zaigosdk.Thinking{
			Type: "enabled",
		},
	}
	chatClient := zaigosdk.MustNewChatModel(cmc)
	response := chatClient.Chat("简述人工智能的历史")
	fmt.Println(response)
}
