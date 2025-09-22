package env

import (
	"errors"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

var isLoadedEnv = false
var ENV = map[string]string{}

func Init(env ...map[string]string) error {
	if isLoadedEnv {
		return errors.New("env already initialized")
	}

	if len(env) > 0 {
		isLoadedEnv = true
		ENV = env[0]
		return nil
	}

	// 로딩 순위
	// 1. 현재 디렉토리의 .파일명.env
	// 2. 현재 디렉토리의 .config.env
	// 3. 상위 디렉토리의 .config.env

	exPaths := strings.Split(os.Args[0], "/")
	fileName := exPaths[len(exPaths)-1]

	path, err := os.Getwd()
	if err != nil {
		return err
	}

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
			isLoadedEnv = true
			return nil
		}
	}

	return errors.New("not found ." + fileName + ".env or .config.env")
}

func Get(key string) (string, error) {
	if !isLoadedEnv {
		return "", errors.New("env not initialized")
	}

	if result, ok := ENV[key]; ok {
		return result, nil
	}

	result := os.Getenv(key)
	if result == "" {
		return "", errors.New("no env " + key)
	}

	return result, nil
}

func GetInt(key string) (int, error) {
	result, err := Get(key)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(result)
}

func GetFloat(key string) (float64, error) {
	result, err := Get(key)
	if err != nil {
		return 0, nil
	}
	return strconv.ParseFloat(result, 64)
}
