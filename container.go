package main

import "fmt"

// Struct of the IOC container
type IOC struct {
	TransientContainer  *transientContainer
	SingletonsContainer *singletonsContainer
}

func (ioc *IOC) Register(name string, obj interface{}, lifeTime int) error {
	switch lifeTime {
	case 0:
		return ioc.TransientContainer.Register(name, obj)
	default:
		return ioc.SingletonsContainer.Register(name, obj)
	}

}

func (ioc *IOC) Resolve(name string) (interface{}, error) {
	if ioc.TransientContainer.Exists(name) {
		return ioc.TransientContainer.Resolve(name)
	}
	if ioc.SingletonsContainer.Exists(name) {
		return ioc.SingletonsContainer.Resolve(name)
	}

	return nil, fmt.Errorf("'%s' is not in IoC container.", name)
}

func (ioc *IOC) Release(name string) error {
	if ioc.TransientContainer.Exists(name) {
		return ioc.TransientContainer.Release(name)
	}
	if ioc.SingletonsContainer.Exists(name) {
		return ioc.SingletonsContainer.Release(name)
	}
	return fmt.Errorf("'%s' is not in IoC container.", name)
}

// Returns new IOC Container Instance
func NewIOC() *IOC {
	return &IOC{
		TransientContainer:  newTransientContainer(),
		SingletonsContainer: newSingletonsContainer(),
	}
}
