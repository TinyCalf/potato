package config

import (
	"os"
	"io/ioutil"
	"fmt"
	"encoding/json"
)

// Compnent 配置文件组件
type Compnent struct {
	Config map[string](map[string]([]Process))
}

// Process ..
type Process struct {
	PeerID string `json:"peerid"`
	Host string `json:"host"`
	RemotePort int `json:"remotePort"`
	ClientPort int `json:"clientPort"`
}

func pathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// Load 加载文件
func (c *Compnent) Load(filePath string) {
	if ok, err := pathExists(filePath); !ok || err != nil {
		panic(fmt.Sprintf("file path %s not exists", filePath))
	}

	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	//将json数据解析到struct中
	config := make(map[string](map[string]([]Process)))
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	c.Config = config
}
