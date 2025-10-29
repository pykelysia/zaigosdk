package zaitype

type (
	ImageModel struct {
		Config
		ImageModelConfig
		URL string
	}

	ImageModelConfig struct {

		// 调用的普通对话模型代码
		Model string `json:"model"`

		// 所需图像的文本描述
		Prompt string `json:"prompt"`

		// 生成图像的质量，默认为 standard。
		// hd: 生成更精细、细节更丰富的图像，整体一致性更高，耗时约20秒；
		// standard: 快速生成图像，适合对生成速度有较高要求的场景，耗时约5-10秒。
		// 此参数仅支持cogview-4-250304
		Quality string `json:"quality"`

		// 图片尺寸，推荐枚举值：1024x1024 (默认), 768x1344, 864x1152, 1344x768, 1152x864, 1440x720, 720x1440。
		// 自定义参数：长宽均需满足512px-2048px之间，需被16整除，并保证最大像素数不超过2^21px。
		Size string `json:"size"`

		// 控制AI生成图片时是否添加水印。
		// true: 默认启用AI生成的显式水印及隐式数字水印，符合政策要求。
		// false: 关闭所有水印，仅允许已签署免责声明的客户使用，签署路径：个人中心-安全管理-去水印管理
		WaterMark bool `json:"watermark_enabled"`

		// 终端用户的唯一标识符。ID长度要求：最少6个字符，最多128个字符，建议使用不包含敏感信息的唯一标识。
		// Required string length: 6 - 128
		UserID string `json:"user_id"`
	}

	ImageResponse struct {
		Created int `json:"created"`
		Data    []struct {
			URL string `json:"url"`
		}
		ContentFilter []struct {
			Role  string `json:"role"`
			Level int    `json:"level"`
		}
	}
)
