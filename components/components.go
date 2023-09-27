package components

import (
	"os/exec"
	"vimana/config"
)

type Component interface {
	InitializeConfig() error
	GetStartCmd() *exec.Cmd
}

type ComponentManager struct {
	ComponentType config.ComponentType
	Component
}

func NewComponentManager(componentType config.ComponentType, root string, nodeType string) *ComponentManager {
	var component Component

	switch componentType {
	case config.Celestia:
		component = NewCelestiaComponent(root, ".vimana/celestia", nodeType)
	case config.Avail:
		component = NewAvailComponent(root, ".vimana/avail", nodeType)
	// case config.Berachain:
	// 	component = berachain.NewBerachainComponent(home)
	default:
		panic("Unknown component type")
	}

	return &ComponentManager{
		ComponentType: componentType,
		Component:     component,
	}
}
