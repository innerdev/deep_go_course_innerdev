package main

import (
	"errors"
)

type Container struct {
	classes map[string]interface{}
}

func NewContainer() *Container {
	return &Container{
		classes: make(map[string]interface{}, 0),
	}
}

func (c *Container) RegisterType(name string, constructor interface{}) {
	if name != "" && constructor != nil {
		c.classes[name] = constructor
	}
}

func (c *Container) Resolve(name string) (interface{}, error) {
	pureConstructor, isFound := c.classes[name]
	if isFound == false {
		return nil, errors.New("Class is not found")
	}

	constructor, ok := pureConstructor.(func() interface{})
	if ok != true {
		return nil, errors.New("Class constructor does not follow the contract")
	}

	return constructor(), nil
}
