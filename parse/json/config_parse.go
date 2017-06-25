package json

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// 配置文件
type Config struct {
	Strs []ConfigItem `json:"strings"`
	Dims []ConfigItem `json:"dimens"`
}

type ConfigItem struct {
	NameSpace string   `json:"ns"`
	Items     []string `json:"items"`
}

func ParseJson(configPath string) ([]ConfigItem, []ConfigItem) {
	if configPath == "" {
		return nil, nil
	}
	var config Config
	bytes, err := ioutil.ReadFile(configPath)
	if err != nil {
		log.Fatalf("file %s read faild : %s", configPath, err.Error())
	}
	err = json.Unmarshal(bytes, &config)
	if err != nil {
		log.Fatalf("parse file %s to json faild : %s", configPath, err.Error())
	}
	return config.Strs, config.Dims
}
