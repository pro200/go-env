package main

import (
	"fmt"

	"github.com/pro200/go-env"
	"github.com/pro200/go-utils"
)

func main() {
	configFile := ".config.env"

	data, err := env.New(configFile)
	if err != nil {
		if err.Error() != env.ERROR_NOT_FOUND {
			utils.Exit("env load error: ", err)
		}

		// env not found, make new env with default values
		defaults := map[string]string{
			"NAME":  "pro200",
			"EMAIL": "pro200@gmail.com",
			"AGE":   "123",
		}
		data, err = env.MakeEnv(defaults, configFile)
		if err != nil {
			utils.Exit("env make error: ", err)
		}
	}

	fmt.Println("env values:", data.Get("NAME"), data.Get("EMAIL"), data.Get("AGE"))
}
