package main

import (
	"errors"
)

type Container struct {
	classes map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		classes: make(map[string]interface{}),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	if name == "" || constructor == nil {
		return
	}

	constructor, ok := constructor.(func() interface{})
	if ok != true {
		return
	}

	c.classes[name] = constructor
}

func (c *Container) Resolve(name string) (interface{}, error) {
	constructor, isFound := c.classes[name]
	if isFound == false {
		return nil, errors.New("Class is not found")
	}

	return constructor.(func() interface{})(), nil
}
