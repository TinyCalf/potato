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
	localPeerInfo *PeerInfo
	Config map[string](map[string]([]*PeerInfo))
	peersOfID map[string]*PeerInfo
}

// PeerInfo ..
type PeerInfo struct {
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
	config := make(map[string](map[string]([]*PeerInfo)))
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

// SetLocalPeer 指定本地peer
func (c *Compnent) SetLocalPeer(appname string, peerid string) {
	env := c.Config[c.Env]
	if env == nil {
		panic("local peer not found!")
	}

	app := env[appname]
	if app == nil {
		panic("local peer not found!")
	}

	for _, peerinfo := range app {
		if peerinfo.PeerID == peerid {
			c.localPeerInfo = peerinfo
			return
		}
	}

	panic("local peer not found!")
}

// GetLocalPeerInfo 获取本地节点信息
func (c *Compnent) GetLocalPeerInfo() *PeerInfo {
	if c.localPeerInfo == nil {
		panic("local peer has not been set!")
	}
	return c.localPeerInfo
}

// GetAllPeerInfo 获取所有远程节点
func (c *Compnent) GetAllPeerInfo() map[string]*PeerInfo {
	m := make(map[string]*PeerInfo)

	env := c.Config[c.Env]
	if env == nil {
		panic("local peer not found!")
	}

	for appname := range env {
		app := env[appname]
		for _, peerinfo := range app {
			m[peerinfo.PeerID] = peerinfo
		}
	}

	return m
}
