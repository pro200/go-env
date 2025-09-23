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

var GlobalEnv *Env

func New() *Env {
	// 로딩 순위
	// ./.파일명.env
	// ./.config.env
	// ../.config.env

	execPath, err := os.Executable()
	if err != nil {
		panic(err)
	}
	fileName := filepath.Base(execPath)

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
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
			GlobalEnv = &Env{loaded: true}
			return GlobalEnv
		}
	}

	panic("not found ." + fileName + ".env or .config.env")
}

func GetEnv() (*Env, error) {
	if !GlobalEnv.loaded {
		return nil, errors.New("env not loaded")
	}

	return GlobalEnv, nil
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
