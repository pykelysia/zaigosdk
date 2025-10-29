package zaigosdk

type (
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

	ImageQualityModel struct {
		Standard string
		HD       string
	}
)

var (
	ApiConfig = API{
		Chat:          "/paas/v4/chat/completions",
		ChatAsyncPost: "/paas/v4/async/chat/completions",
		ChatAsyncGet:  "/paas/v4/async-result",
		Image:         "/paas/v4/images/generations",
		AudioToTxt:    "/paas/v4/audio/transcriptions",
		TxtToAudio:    "/paas/v4/audio/speech",
	}
	ImageQuality = ImageQualityModel{
		Standard: "standard",
		HD:       "hd",
	}
	ImageSiae = []string{
		"1024x1024",
		"768x1344",
		"864x1152",
		"1344x768",
		"1152x864",
		"1440x720",
		"720x1440",
	}
)

const (
	ROLESYSTEM    = "system"
	ROLEUSER      = "user"
	ROLEASSISTANT = "assistant"
	charset       = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	GLM4_6             = "glm-4.6"
	GLM4_5             = "glm-4.5"
	GLM4_5_AIR         = "glm-4.5-air"
	GLM4_5_X           = "glm-4.5-x"
	GLM4_5_AIRX        = "glm-4.5-airx"
	GLM4_5_FLASH       = "glm-4.5-flash"
	GLM4_PLUS          = "glm-4-plus"
	GLM4_AIR_250414    = "glm-4-air-250414"
	GLM4_AIRX          = "glm-4-airx"
	GLM4_FLASHX        = "glm-4-flashx"
	GLM4_FLASHX_250414 = "glm-4-flashx-250414"
)
