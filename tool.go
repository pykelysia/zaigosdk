package zaigosdk

import (
	"math/rand"
	"time"
)

func getRandomString(left, right int) string {
	rand.Seed(time.Now().Unix())
	length := rand.Intn(right-left) + left
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))] // 生成随机小写字母
	}
	return string(result)
}
