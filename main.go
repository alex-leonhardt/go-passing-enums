package main

import (
	"fmt"

	"github.com/alex-leonhardt/go-passing-enums/pkg/config"
)

func main() {
	cfg, err := config.New(config.Env)
	fmt.Printf("%#v %#v\n", cfg, err)
	cfg, err = config.New(config.File)
	fmt.Printf("%#v %#v\n", cfg, err)
	cfg, err = config.New(42)
	fmt.Printf("%#v %#v\n", cfg, err)
}
