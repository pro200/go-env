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

var GlobalEnv *Env

// Key: 비어 있으면 Val이 주석 및 출력문으로 처리, Val도 비어있을경우 공백줄 추가
type Default struct {
	Key string
	Val string
}

func New(path ...string) (*Env, error) {
	// from path
	if len(path) > 0 {
		fullPath, err := filepath.Abs(path[0])
		fmt.Println(fullPath)
		if err != nil {
			return nil, err
		}

		if err := godotenv.Load(fullPath); err == nil {
			GlobalEnv = &Env{loaded: true}
			return GlobalEnv, nil
		} else {
			return nil, errors.New(ERROR_NOT_FOUND)
		}
	}

	// 로딩 순위
	// ./.파일명.env
	// ./.config.env
	// ../.config.env
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
			GlobalEnv = &Env{loaded: true}
			return GlobalEnv, nil
		}
	}

	return nil, errors.New(ERROR_NOT_FOUND)
}

// .env 로드 실패시 값을 받아 새로 생성
func Create(defaults []Default, path ...string) (*Env, error) {
	if GlobalEnv != nil && GlobalEnv.loaded {
		return nil, errors.New("env already loaded")
	}

	// 입력 받기
	fmt.Println("Create a new .config.env file.")
	fmt.Println()

	var lines []string
	for i := range defaults {
		key := strings.ToUpper(defaults[i].Key)
		value := defaults[i].Val

		if key == "" && value == "" {
			lines = append(lines, "")
			continue
		}
		if key == "" {
			fmt.Println("# " + value)
			lines = append(lines, "# "+value)
			continue
		}

		var input string
		fmt.Printf("%s [%s]:", key, value)
		fmt.Scanln(&input)

		// 기본값 사용
		if input == "" {
			input = value
		}

		lines = append(lines, key+": "+input)
	}

	// 저장 여부 확인
	var confirm string
	fmt.Print("\nWould you like to save it?  [y|N]")
	fmt.Scanln(&confirm)

	if strings.ToLower(confirm) != "y" {
		return nil, errors.New("Creating a new env file was canceled.")
	}

	// 파일 저장
	var fullPath string
	var err error

	if len(path) > 0 {
		fullPath, err = filepath.Abs(path[0])
		if err != nil {
			return nil, err
		}
	} else {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		paths := strings.Split(wd, string(os.PathSeparator))
		fullPath = filepath.Join(strings.Join(paths, string(os.PathSeparator)), ".config.env")
	}

	// write file to fullPath
	if err := os.WriteFile(fullPath, []byte(strings.Join(lines, "\n")), 0644); err != nil {
		return nil, err
	}

	if err := godotenv.Load(fullPath); err == nil {
		GlobalEnv = &Env{loaded: true}
		return GlobalEnv, nil
	}

	return nil, errors.New("Failed to save the env file.")
}

// 전역 Env 반환
func Load() (*Env, error) {
	if GlobalEnv == nil || !GlobalEnv.loaded {
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
