package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"server/global"
	"server/model/other"
	"server/utils"
)

type GaodeService struct {
}

// 根据IP地址获取地理位置信息
func (gaodeService *GaodeService) GetLocationByIP(ip string) (other.IPResponse, error) {
	data := other.IPResponse{}
	key := global.Config.Gaode.Key
	urlStr := "https://restapi.amap.com/v3/ip"
	method := "GET"
	params := map[string]string{
		"key": key,
		"ip":  ip,
	}
	res, err := utils.HttpRequest(urlStr, method, nil, params, nil)
	if err != nil {
		return data, err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return data, fmt.Errorf("request failed with status code: %d", res.StatusCode)
	}

	byteData, err := io.ReadAll(res.Body)
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(byteData, &data)
	if err != nil {
		return data, err
	}

	return data, nil
}
