package env_test

import (
	"testing"

	"github.com/pro200/go-env"
)

func TestEnv(t *testing.T) {
	data, err := env.NewEnv()
	if err != nil {
		t.Error("env load error:", err)
		return
	}

	strVal := data.Get("STRING")
	intVal := data.GetInt("INT")
	floatVal := data.GetFloat("FLOAT")

	if strVal != "hello" || intVal != 1234 || floatVal != 12.34 {
		t.Error("Wrong result")
	}
}
