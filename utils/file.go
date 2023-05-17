package utils

import (
	"os"
	"strings"
)

func FileIsExist(path string) bool {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func GetHomeDir() string {
	home, _ := os.UserHomeDir()
	return home
}

func GetFilePath() string {
	return strings.Join([]string{GetHomeDir(), ".memo.json"}, "/")
}
