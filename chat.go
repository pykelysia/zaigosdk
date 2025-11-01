package zaigosdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/pykelysia/zaigosdk/zaitype"
)

type (
	ChatModel          zaitype.ChatModel
	ChatMessage        zaitype.ChatMessage
	ChatModelConfig    zaitype.ChatModelConfig
	ChatThinking       zaitype.ChatThinking
	ChatResponseFormat zaitype.ChatResponseFormat
	ChatResponse       zaitype.ChatResponse
	AsyncResponse      zaitype.AsyncResponse
)

func MustDefaultChatModel() *ChatModel {
	config := MustNewConfig()
	return &ChatModel{
		Config: config,
		URL:    config.URL + ApiConfig.Chat,
		ChatModelConfig: zaitype.ChatModelConfig{
			Model: GLM4_6,
			Messages: []zaitype.ChatMessage{
				{
					Role:    ROLESYSTEM,
					Content: "你是一个有用的AI助手。",
				},
			},
			MaxTokens: 65536,
			RequestID: getRandomString(6, 64),
			UserID:    getRandomString(6, 128),
		},
	}
}

// 需要将所有参数填写
func MustNewChatModel(cmc ChatModelConfig) *ChatModel {
	config := MustNewConfig()
	return &ChatModel{
		Config:          config,
		URL:             config.URL + ApiConfig.Chat,
		ChatModelConfig: zaitype.ChatModelConfig(cmc),
	}
}

func (cm *ChatModel) toString() (s string, err error) {
	request := cm.ChatModelConfig
	if request.RequestID == "" {
		request.RequestID = getRandomString(6, 64)
	}
	if request.UserID == "" {
		request.UserID = getRandomString(6, 128)
	}
	if request.MaxTokens == 0 {
		request.MaxTokens = 65536
	}
	bytes, err := json.Marshal(request)
	if err != nil {
		return s, err
	}
	s = string(bytes)
	return
}

func (cm *ChatModel) appendConversation(role, content string) {
	cm.Messages = append(cm.Messages, zaitype.ChatMessage(ChatMessage{
		Role:    role,
		Content: content,
	}))
}

func (cm *ChatModel) Chat(content string) (ChatResponse, error) {
	cm.appendConversation(ROLEUSER, content)
	url := cm.URL
	s, err := cm.toString()
	if err != nil {
		return ChatResponse{}, fmt.Errorf("[Error]: message: 请求参数序列化失败, %v", err)
	}
	payload := strings.NewReader(s)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cm.ApiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("[Error]: message: 远程服务器错误, %v", err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var response ChatResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ChatResponse{}, fmt.Errorf("[Error]: message: 响应获取失败, %v", err)
	}
	if response.Error.Code != "" {
		e := response.Error
		return ChatResponse{}, fmt.Errorf("[Error]: code: %s, message: %s\n", e.Code, e.Message)
	}

	ai_response := response.Choices[0].Message.Content
	cm.appendConversation(ROLEASSISTANT, ai_response)

	return response, nil
}

func (cm *ChatModel) ChatAsync(content string) (AsyncResponse, error) {
	cm.appendConversation(ROLEUSER, content)
	url := cm.Config.URL + ApiConfig.ChatAsyncPost
	s, err := cm.toString()
	if err != nil {
		return AsyncResponse{}, fmt.Errorf("[Error]: message: 请求参数序列化失败, %v", err)
	}
	payload := strings.NewReader(s)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", cm.ApiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return AsyncResponse{}, fmt.Errorf("[Error]: message: 远程服务器错误, %v", err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var response AsyncResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return AsyncResponse{}, fmt.Errorf("[Error]: message: 响应获取失败, %v", err)
	}
	if response.Error.Code != "" {
		e := response.Error
		return AsyncResponse{}, fmt.Errorf("[Error]: code: %s, message: %s\n", e.Code, e.Message)
	}

	return response, nil
}
