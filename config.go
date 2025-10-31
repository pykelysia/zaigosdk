package zaigosdk

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/pykelysia/zaigosdk/zaitype"
	"github.com/spf13/viper"
)

func MustNewConfig() zaitype.Config {
	// 加载环境变量
	godotenv.Load("./config/.env")

	// 加载配置文件
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("read file config.yaml error: %s", err))
	}

	return zaitype.Config{
		ApiKey: os.Getenv("AI_API_KEY"),
		URL:    viper.GetString("AI_URL"),
	}
}
