package zaigosdk

import "github.com/pykelysia/zaigosdk/config"

type (
	ChatModel struct {
		config.ModelConfig
		ChatModelConfig
		url string
	}

	Message struct {
		// 信息作者的角色，包括 "user", "system", "assistant"
		Role string `json:"role"`

		// 信息内容
		Content string `json:"content"`
	}

	Thinking struct {
		Type string `json:"type"`
	}

	ResponseFormat struct {
		Type string `json:"type"`
	}

	ChatModelConfig struct {

		// 调用的普通对话模型代码
		Model string `json:"model"`

		// 对话内容列表
		Messages []Message `json:"messages"`

		// 是否启用流式输出
		Stream bool `json:"stream"`

		// 是否开启思考链，仅 GLM4.5 以上模型支持
		// 支持的值为 "enabled" , "disabled"，默认状态下为 "enabled"
		// 开启之后 GLM-4.6, GLM4.5 会自动判断是否需要进行思考，GLM4.5V 为强制思考
		Thinking Thinking `json:"thinking"`

		// 是否启用采样策略来生成文本，默认为 true。
		// 开启时表示模型会启用 temperature, top_p 等参数进行随机采样，生成更多样化的输出
		// 关闭后，模型总是模型总是选择概率最高的词汇，生成更确定性的输出，此时 temperature 和 top_p 参数将被忽略。
		// 对于需要一致性和可重复性的任务（如代码生成、翻译），建议设置为 false。
		DoSample bool `json:"do_sample"`

		// 采样温度，控制输出的随机性和创造性，取值范围为 [0.0, 1.0]，限两位小数。
		// 对于GLM-4.6系列默认值为 1.0，GLM-4.5系列默认值为 0.6，GLM-4系列默认值为 0.75。
		// 较高的值（如0.8）会使输出更随机、更具创造性，适合创意写作和头脑风暴；较低的值（如0.2）会使输出更稳定、更确定，适合事实性问答和代码生成。
		// 建议根据应用场景调整 top_p 或 temperature 参数，但不要同时调整两个参数。
		Temperature float32 `json:"temperature"`

		// 核采样（nucleus sampling）参数，是temperature采样的替代方法，取值范围为 (0.0, 1.0]，限两位小数。
		// 对于GLM-4.6 GLM-4.5系列默认值为 0.95，GLM-4系列默认值为 0.9。
		// 模型只考虑累积概率达到top_p的候选词汇。
		// 例如：0.1表示只考虑前10%概率的词汇，0.9表示考虑前90%概率的词汇。
		// 较小的值会产生更集中、更一致的输出；较大的值会增加输出的多样性。
		// 建议根据应用场景调整 top_p 或 temperature 参数，但不建议同时调整两个参数。
		TopP float32 `json:"top_p"`

		// 模型输出的最大令牌token数量限制。
		// GLM-4.6最大支持128K输出长度，GLM-4.5最大支持96K输出长度，建议设置不小于1024。
		// 令牌是文本的基本单位，通常1个令牌约等于0.75个英文单词或1.5个中文字符。
		// 设置合适的max_tokens可以控制响应长度和成本，避免过长的输出。
		// 如果模型在达到max_tokens限制前完成回答，会自然结束；如果达到限制，输出可能被截断。
		// 默认值和最大值等更多详见 https://docs.bigmodel.cn/cn/guide/start/concept-param#max-tokens
		// 1 <= max_tokens <= 131072
		MaxTokens int32 `json:"max_tokens"`

		// 停止词列表，当模型生成的文本中遇到这些指定的字符串时会立即停止生成。
		// 目前仅支持单个停止词，格式为["stop_word1"]。
		// 停止词不会包含在返回的文本中。
		// 这对于控制输出格式、防止模型生成不需要的内容非常有用，例如在对话场景中可以设置["Human:"]来防止模型模拟用户发言。
		Stop []string `json:"stop"`

		// 指定模型的响应输出格式，默认为text，仅文本模型支持此字段。
		// 支持两种格式：{ "type": "text" } 表示普通文本输出模式，模型返回自然语言文本；
		// 				{ "type": "json_object" } 表示JSON输出模式，模型会返回有效的JSON格式数据，适用于结构化数据提取、API响应生成等场景。
		// 使用JSON模式时，建议在提示词中明确说明需要JSON格式输出。
		ResponseFormat ResponseFormat `json:"responae_format"`

		// 请求唯一标识符。由用户端传递，建议使用UUID格式确保唯一性，若未提供平台将自动生成。
		RequestID string `json:"request_id"`

		// 终端用户的唯一标识符。ID长度要求：最少6个字符，最多128个字符，建议使用不包含敏感信息的唯一标识。
		// Required string length: 6 - 128
		UserID string `json:"user_id"`
	}

	response struct {
		Choices []struct {
			FinishReason string `json:"finish_reason"`
			Index        int    `json:"index"`
			Message      struct {
				Content string `json:"content"`
				Role    string `json:"role"`
			} `json:"message"`
		} `json:"choices"`
		Created    int    `json:"created"`
		Id         string `json:"id"`
		Model      string `json:"model"`
		Request_id string `json:"request_id"`

		// 使用的额度
		Usage struct {
			Completion            int `json:"completion"`
			Prompt_tokens         int `json:"prompt_tokens"`
			Prompt_tokens_details struct {
				Cached_token int `json:"cached_token"`
			} `json:"prompt_tokens_details"`
			Total_tokens int `json:"total_tokens"`
		} `json:"usage"`
		Error struct {
			Code    string `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}
)
