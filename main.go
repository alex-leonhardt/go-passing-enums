package main

import (
	"fmt"

	"github.com/alex-leonhardt/go-passing-enums/pkg/config"
)

func main() {
	cfg, err := config.New(config.Env)
	fmt.Printf("%#v %#v", cfg, err)
}
