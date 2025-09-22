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

	strVal, err := data.Get("STRING")
	if err != nil {
		t.Error(err)
	}

	intVal, err := data.GetInt("INT")
	if err != nil {
		t.Error(err)
	}

	floatVal, err := data.GetFloat("FLOAT")
	if err != nil {
		t.Error(err)
	}

	if strVal != "hello" || intVal != 1234 || floatVal != 12.34 {
		t.Error("Wrong result")
	}
}
