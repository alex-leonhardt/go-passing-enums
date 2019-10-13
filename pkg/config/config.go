package config

import (
	"errors"

	"github.com/alex-leonhardt/go-passing-enums/pkg/config/env"
	"github.com/alex-leonhardt/go-passing-enums/pkg/config/file"
)

// Configurer describes a config provider
type Configurer interface {
	Get(string) string
	Set(string) bool
	Del(string) bool
}

// T is used to select the type of config to return
type T uint

// Constants for T
const (
	Env T = iota
	File
	Unknown
)

// New takes a config.T type and returns an Configurer
func New(t T) (Configurer, error) {
	switch t {
	case Env:
		return env.New(), nil
	case File:
		return file.New(), nil
	default:
		return nil, errors.New("eh?")
	}
}
