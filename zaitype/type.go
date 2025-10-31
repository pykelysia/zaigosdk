package zaitype

type (
	Config struct {
		// 你的智谱 Api-Key值。记录在文件 `/config/.env` 中的 AI_API_KEY
		ApiKey string
		// 调用智谱 Api 的基本 url。记录在文件 `/config/config.yaml` 中的 AI_URL
		URL string
	}
)
