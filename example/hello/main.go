package main

import (
	"fmt"

	"github.com/pro200/go-env"
)

func main() {
	config, err := env.NewEnv()
	if err != nil {
		fmt.Println()
		panic(err)
	}

	fmt.Println(config)
}
