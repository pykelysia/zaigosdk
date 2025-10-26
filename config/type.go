package config

type (
	Config struct {
		// 你的智谱 Api-Key值。记录在文件 `/config/.env` 中的 AI_API_KEY
		ApiKey string
		// 调用智谱 Api 的基本 url。记录在文件 `/config/config.yaml` 中的 AI_URL
		URL string
	}

	ModelConfig struct {
		Config
		// AI 模型的型号。
		Model string
		// 调用的 API-url。
		API API
	}

	API struct {
		// 对话补全 Api 的 url
		Chat string

		// 异步对话补全 Api 的 url，需要通过 ChatAsyncGet + "id" 的方法获取结果。
		// id 来自对应异步请求的返回。
		ChatAsyncPost string

		// 获取异步补全的结果 Api 的 url，通过 ChatAsyncGet + "id" 的方法获取结果。
		// id 来自对应异步请求的返回。
		// 建议用: fmt.Sprintf("%s/%s", ChatAsyncGet, id)
		ChatAsyncGet string

		// 生成图片 Api 的 url，生成的图片会返回 url。该 url 保存 30 天的时间。
		Image string

		// 语音转文本 Api 的url。
		AudioToTxt string

		// 文本转语音 Api 的 url。将会返回 .wav 文件。
		TxtToAudio string

		// 为其他方向的 api 做准备。
		NewFunction string
	}
)

var (
	// 存储可能用到的大部分 Api 的 url
	ApiConfig = API{
		Chat:          "/paas/v4/chat/completions",
		ChatAsyncPost: "/paas/v4/async/chat/completions",
		ChatAsyncGet:  "/paas/v4/async-result",
		Image:         "/paas/v4/images/generations",
		AudioToTxt:    "/paas/v4/audio/transcriptions",
		TxtToAudio:    "/paas/v4/audio/speech",
		NewFunction:   "",
	}
)
