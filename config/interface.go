package config

import (
	"potato/piface"
)

// ICompnent ..
type ICompnent interface {
	piface.ICompnent
	Load(filePath string) //从文件读取配置
	SetEnv(env string)	//设置当前环境，不同环境不同配置
	SetLocalPeer(appname string, peerid string) //指定本地节点
	GetLocalPeerInfo() *PeerInfo	//获取本地节点信息
	GetAllPeerInfo() map[string]*PeerInfo //获取所有远程节点
}