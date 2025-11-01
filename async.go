package zaigosdk

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pykelysia/zaigosdk/zaitype"
)

type (
	GetResponse zaitype.GetResponse
)

func CheckResult(id string) (GetResponse, error) {
	config := MustNewConfig()
	url := config.URL + ApiConfig.AsyncGet + "/" + id

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", config.ApiKey))
	req.Header.Add("Content-Type", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return GetResponse{}, fmt.Errorf("[Error]: message: 远程服务器错误, %v", err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	var response GetResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return GetResponse{}, fmt.Errorf("[Error]: message: 响应获取失败, %v", err)
	}
	if response.Error.Code != "" {
		e := response.Error
		return GetResponse{}, fmt.Errorf("[Error]: code: %s, message: %s\n", e.Code, e.Message)
	}

	return response, nil
}
