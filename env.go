package env

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var isLoaded bool

type Env struct{}

func New() (*Env, error) {
	// 로딩 순위
	// 1. 현재 디렉토리의 .파일명.env
	// 2. 현재 디렉토리의 .config.env
	// 3. 상위 디렉토리의 .config.env

	exPaths := strings.Split(os.Args[0], "/")
	fileName := exPaths[len(exPaths)-1]

	path, _ := os.Getwd()
	paths := strings.Split(path, "/")

	envFiles := []string{
		strings.Join(paths, "/") + "/." + fileName + ".env",
		strings.Join(paths, "/") + "/.config.env",
	}

	if len(paths) > 1 {
		envFiles = append(envFiles, strings.Join(paths[:len(paths)-1], "/")+"/.config.env")
	}

	for _, file := range envFiles {
		if err := godotenv.Load(file); err == nil {
			isLoaded = true
			env := Env{}
			return &env, nil
		}
	}

	return &Env{}, errors.New("not found ." + fileName + ".env or .config.env")
}

func GetEnv() (*Env, error) {
	if !isLoaded {
		return nil, errors.New("env not loaded. use env.New()")
	}

	return &Env{}, nil
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
