package main

import (
	"fmt"
	"sync"
)

type singletonsContainer struct {
	ThreadLocker *sync.RWMutex
	Mapper       []singletons
}

type singletons struct {
	name string
	val  interface{}
}

func newSingletonsContainer() *singletonsContainer {
	container := &singletonsContainer{
		ThreadLocker: &sync.RWMutex{},
		Mapper:       []singletons{},
	}
	return container
}

func (c *singletonsContainer) Register(name string, obj interface{}) error {
	defer c.ThreadLocker.Unlock()
	c.ThreadLocker.Lock()

	if c.Exists(name) {
		return fmt.Errorf("'%s' is already in IoC container.", name)
	}

	c.Mapper = append(c.Mapper, singletons{name: name, val: obj})
	return nil
}

func (c *singletonsContainer) Resolve(name string) (interface{}, error) {
	for _, s := range c.Mapper {
		if s.name == name {
			return &s.val, nil
		}
	}
	return nil, fmt.Errorf("'%s' is not in IoC container.", name)
}

func (c *singletonsContainer) Release(name string) error {
	for _, s := range c.Mapper {
		if s.name == name {
			s.name = ""
		}
	}
	return nil
}

func (c *singletonsContainer) Exists(name string) bool {
	for _, s := range c.Mapper {
		if s.name == name {
			return true
		}
		return false
	}
	return false
}
