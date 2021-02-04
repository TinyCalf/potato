package rpc

import (
	"fmt"
	"sync"
)

type dispatcher struct {
	services ServiceMap
	lock sync.RWMutex
}

func newDispatcher(services ServiceMap) *dispatcher {
	return &dispatcher{
		services: services,
	}
}

func (d *dispatcher) reload(services ServiceMap) {
	d.lock.Lock()
	defer d.lock.Unlock()
	d.services = services
}

func (d *dispatcher) findMethod(namespace, sid, methodName string) (Method, error) {
	d.lock.RLock()
	defer d.lock.RUnlock()

	namespaces := d.services[namespace]
	if namespaces == nil {
		return nil, fmt.Errorf("rpc services namespace %s not found", namespace)
	}

	service := namespaces[sid]
	if service == nil {
		return nil, fmt.Errorf("rpc service %s not found", sid)
	}

	method := service.Methods[methodName]
	if method == nil {
		return nil, fmt.Errorf("rpc method %s not found", methodName)
	}

	return method, nil
}

func (d *dispatcher) route(msg *message) ([]byte, error) {
	method, err := d.findMethod(msg.namespace, msg.serviceName, msg.methodName)
	if err != nil {
		return nil, err
	}
	return method(msg.data)
}