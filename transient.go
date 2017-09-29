package main

import (
	"fmt"
	"sync"
)

type transientContainer struct {
	ThreadLocker *sync.RWMutex
	Mapper       map[string]interface{}
}

func newTransientContainer() *transientContainer {
	container := &transientContainer{
		ThreadLocker: &sync.RWMutex{},
		Mapper:       make(map[string]interface{}),
	}
	return container
}

func (c *transientContainer) Register(name string, obj interface{}) error {
	defer c.ThreadLocker.Unlock()
	c.ThreadLocker.Lock()

	if c.Exists(name) {
		return fmt.Errorf("'%s' is already in IoC container.", name)
	}
	c.Mapper[name] = obj
	return nil
}

func (c *transientContainer) Resolve(name string) (interface{}, error) {
	return c.Mapper[name], nil
}

func (c *transientContainer) Release(name string) error {
	delete(c.Mapper, name)
	return nil
}

func (c *transientContainer) Exists(name string) bool {
	if _, exists := c.Mapper[name]; exists {
		return true
	}
	return false
}
