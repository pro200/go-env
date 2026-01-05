package env

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Env struct {
	loaded bool
}

func NewEnv(path ...string) (*Env, error) {
	// 이미 로드된 경우
	if envPath := os.Getenv("ENV_PATH"); envPath != "" {
		return &Env{loaded: true}, nil
	}

	// from path
	if len(path) > 0 {
		fullPath, err := filepath.Abs(path[0])
		if err != nil {
			return nil, err
		}

		if err := godotenv.Load(fullPath); err == nil {
			return &Env{loaded: true}, nil
		}

		return nil, errors.New(ERROR_NOT_FOUND)
	}

	// 로딩 순위 - ./.파일명.env -> ./.config.env -> ../.config.env
	execPath, err := os.Executable()
	if err != nil {
		return nil, err
	}

	fileName := filepath.Base(execPath)
	if strings.HasPrefix(fileName, "___go_") {
		return nil, fmt.Errorf("this file name \"%s\" not supported", fileName)
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	paths := strings.Split(wd, string(os.PathSeparator))

	envFiles := []string{
		filepath.Join(wd, "."+fileName+".env"),
		filepath.Join(wd, ".config.env"),
	}

	if len(paths) > 1 {
		envFiles = append(envFiles,
			filepath.Join(strings.Join(paths[:len(paths)-1], string(os.PathSeparator)), ".config.env"),
		)
	}

	for _, file := range envFiles {
		if err := godotenv.Load(file); err == nil {
			os.Setenv("ENV_PATH", file)
			return &Env{loaded: true}, nil
		}
	}

	return nil, errors.New(ERROR_NOT_FOUND)
}

func (e *Env) Get(key string, defaultVal ...string) string {
	val := os.Getenv(key)
	if val == "" && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return val
}

func (e *Env) GetInt(key string, defaultVal ...int) int {
	result := e.Get(key)
	if result == "" {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}
	data, err := strconv.Atoi(result)
	if err != nil && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return data
}

func (e *Env) GetInt64(key string, defaultVal ...int64) int64 {
	result := e.Get(key)
	if result == "" {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}
	data, err := strconv.ParseInt(result, 10, 64)
	if err != nil && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return data
}

func (e *Env) GetFloat(key string, defaultVal ...float64) float64 {
	result := e.Get(key)
	if result == "" && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	data, err := strconv.ParseFloat(result, 64)
	if err != nil && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return data
}

func (e *Env) GetBool(key string, defaultVal ...bool) bool {
	result := e.Get(key)
	if result == "" {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return false
	}
	data, err := strconv.ParseBool(result)
	if err != nil && len(defaultVal) > 0 {
		return defaultVal[0]
	}
	return data
}
