package env

import (
	"errors"
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
		} else {
			return nil, errors.New(ERROR_NOT_FOUND)
		}
	}

	// 로딩 순위 - ./.파일명.env -> ./.config.env -> ../.config.env
	execPath, err := os.Executable()
	if err != nil {
		return nil, err
	}
	fileName := filepath.Base(execPath)

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

func (e *Env) Get(key string) string {
	return os.Getenv(key)
}

func (e *Env) GetInt(key string) int {
	result := e.Get(key)
	data, _ := strconv.Atoi(result)
	return data
}

func (e *Env) GetFloat(key string) float64 {
	result := e.Get(key)
	data, _ := strconv.ParseFloat(result, 64)
	return data
}

func (e *Env) GetBool(key string) bool {
	result := e.Get(key)
	data, _ := strconv.ParseBool(result)
	return data
}
