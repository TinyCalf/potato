package rpc

// Service 代表一组服务
type Service struct {
	Host string // ip
	Port int
	NameSpace string // 所属命名空间
	ServiceName string // 服务名称
	Methods map[string]Method //方法
}

// NewService 创建一个新的服务
func NewService(namespace, serviceName string) *Service {
	return &Service{
		NameSpace: namespace,
		ServiceName: serviceName,
		Methods: make(map[string]Method),
	}
}

// AddMethod 给Service添加方法
func (s *Service) AddMethod(name string, method Method) {
	s.Methods[name] = method
}

// ServiceMap 一组Service的集合 
// ServiceMap -> namespace -> sid -> service
type ServiceMap map[string]map[string]*Service

// NewServiceMap ..
func NewServiceMap() ServiceMap {
	return make(map[string]map[string]*Service)
}

// Add 设置节点信息
func (s ServiceMap) Add(service *Service) {
	namespaces := s[service.NameSpace]
	if namespaces == nil {
		s[service.NameSpace] = make(map[string]*Service)
		namespaces = s[service.NameSpace]
	}
	namespaces[service.ServiceName] = service
}