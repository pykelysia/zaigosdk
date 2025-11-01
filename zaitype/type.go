package zaitype

type (
	Config struct {
		// 你的智谱 Api-Key值。记录在文件 `/config/.env` 中的 AI_API_KEY
		ApiKey string
		// 调用智谱 Api 的基本 url。记录在文件 `/config/config.yaml` 中的 AI_URL
		URL string
	}
	Error struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	Message struct {
		Role             string `json:"role"`
		Content          string `json:"content"`
		ReasoningContent string `json:"reasoning_content"`
		// 后续添加
	}

	Choices struct {
		Index   int     `json:"index"`
		Message Message `json:"message"`

		// 推理终止原因。
		// 'stop’表示自然结束或触发stop词，
		// 'tool_calls’表示模型命中函数，
		// 'length’表示达到token长度限制，
		// 'sensitive’表示内容被安全审核接口拦截（用户应判断并决定是否撤回公开内容），
		// 'network_error’表示模型推理异常。
		FinishReason string `json:"finish_reason"`
	}
	Usage struct {
		CompletionTokens      int `json:"completion_tokens"`
		Prompt_tokens         int `json:"prompt_tokens"`
		Prompt_tokens_details struct {
			Cached_token int `json:"cached_token"`
		} `json:"prompt_tokens_details"`
		Total_tokens int `json:"total_tokens"`
	}
	// 后续任务
	VideoResult   struct{}
	WebSearch     struct{}
	ContentFilter struct{}
)
