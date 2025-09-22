package env_test

import (
	"testing"

	"github.com/pro200/go-env"
)

func TestEnv(t *testing.T) {
	err := env.Init(map[string]string{
		"STRING": "hello",
		"INT":    "1234",
		"FLOAT":  "12.34",
	})

	if err != nil {
		t.Error(err)
	}

	strVal, err := env.Get("STRING")
	if err != nil {
		t.Error(err)
	}

	intVal, err := env.GetInt("INT")
	if err != nil {
		t.Error(err)
	}

	floatVal, err := env.GetFloat("FLOAT")
	if err != nil {
		t.Error(err)
	}

	if strVal != "hello" || intVal != 1234 || floatVal != 12.34 {
		t.Error("Wrong result")
	}
}
