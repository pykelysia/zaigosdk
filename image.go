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
	ImageModel       zaitype.ImageModel
	ImageModelConfig zaitype.ImageModelConfig
	ImageResponse    zaitype.ImageResponse
)

func MustDefaultImageModel() *ImageModel {
	config := MustNewConfig()
	return &ImageModel{
		Config: config,
		URL:    config.URL + ApiConfig.Image,
		ImageModelConfig: zaitype.ImageModelConfig{
			Model:     GLM4_6,
			Prompt:    "",
			Quality:   ImageQuality.Standard,
			Size:      ImageSiae[0],
			WaterMark: true,
			UserID:    getRandomString(6, 128),
		},
	}
}

func MustNewImageModel(imc ImageModelConfig) *ImageModel {
	config := MustNewConfig()
	return &ImageModel{
		Config:           config,
		URL:              config.URL + ApiConfig.Image,
		ImageModelConfig: zaitype.ImageModelConfig(imc),
	}
}

func (im *ImageModel) toString() (s string, err error) {
	request := im.ImageModelConfig
	if request.UserID == "" {
		request.UserID = getRandomString(6, 128)
	}
	if request.Quality == "" {
		request.Quality = ImageQuality.Standard
	}
	if request.Size == "" {
		return "", fmt.Errorf("图片尺寸不能为空")
	}
	bytes, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	s = string(bytes)
	return s, nil
}

func (im *ImageModel) Chat(content string) (ImageResponse, error) {
	im.Prompt = content
	url := im.URL
	s, err := im.toString()
	if err != nil {
		return ImageResponse{}, fmt.Errorf("[Error]: message: 请求参数序列化失败, %v", err)
	}
	payload := strings.NewReader(s)

	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", im.ApiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return ImageResponse{}, fmt.Errorf("[Error]: message: 远程服务器错误, %v", err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var response ImageResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return ImageResponse{}, fmt.Errorf("[Error]: message: 响应获取失败, %v", err)
	}
	if response.Error.Code != "" {
		e := response.Error
		return response, fmt.Errorf("[Error]: code: %s, message: %s\n", e.Code, e.Message)
	}

	return response, nil
}
