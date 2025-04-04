package forger

import (
	"context"
	"fmt"
)

type Plan struct {
	Name    string            `json:"name"`
	Image   string            `json:"image"`
	Command []string          `json:"command"`
	Env     map[string]string `json:"env"`
}

type Forger interface {
	Forge(context.Context, Plan) ([]byte, error)
}

func GetForgerForTarget(target string) (Forger, error) {
	f, ok := registered[target]
	if !ok {
		return nil, fmt.Errorf("%q is not a known target", target)
	}
	return f(), nil
}

type registeredFunc func() Forger

var registered map[string]registeredFunc = map[string]registeredFunc{}

func registerForger(name string, f registeredFunc) error {
	if _, ok := registered[name]; ok {
		return fmt.Errorf("%s forger is already registered", name)
	}

	registered[name] = f
	return nil
}
