package env_test

import (
	"testing"

	"github.com/pro200/go-env"
)

func TestEnv(t *testing.T) {
	// .config.env
	// STRING: hello
	// INT:    1234
	// FLOAT:  12.34

	data, err := env.New()
	if err != nil {
		t.Error(err)
	}

	strVal := data.Get("STRING")
	intVal := data.GetInt("INT")
	floatVal := data.GetFloat("FLOAT")

	if strVal != "hello" || intVal != 1234 || floatVal != 12.34 {
		t.Error("Wrong result")
	}
}
