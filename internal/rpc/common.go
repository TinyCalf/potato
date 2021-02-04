package rpc

// Method 是rpc方法的实现 由用户定义
type Method func([]byte) ([]byte, error)

// Options 是需要的参数集合
type Options struct {
	Port int
	Services ServiceMap
}