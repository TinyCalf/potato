package config

import (
	"os"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"potato/piface"
)

// Compnent 配置文件组件
type Compnent struct {
	piface.BaseCompnent
	Env string
	Config map[string](map[string]([]Peer))
}

// Peer ..
type Peer struct {
	PeerID string `json:"peerid"`
	Host string `json:"host"`
	RemotePort int `json:"remotePort"`
	ClientPort int `json:"clientPort"`
}

// NewCompnent 新建一个配置模块
func NewCompnent() ICompnent {
	return &Compnent{Env:"dev"}
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


// GetName 获取组件名称 实现Compnent
func (c *Compnent) GetName() string {
	return "Config"
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
	config := make(map[string](map[string]([]Peer)))
	err = json.Unmarshal(data, &config)
	if err != nil {
		panic(err)
	}
	c.Config = config
}

// SetEnv 设置运行环境
func (c *Compnent) SetEnv(env string) {
	c.Env = env
}
