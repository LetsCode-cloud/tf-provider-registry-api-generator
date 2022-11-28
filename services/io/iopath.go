package io

import (
	"log"
	"os"
	"path/filepath"
)

func GetAbsPath(relPath string) string {
	abs, err := filepath.Abs(relPath)
	if err != nil {
		log.Fatalln(err)
	}
	return abs
}

func DirExists(absPath string) bool {
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return false
	}
	return true
}

func RemoveRecursivelyDir(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		log.Fatalln(err)
	}
}

func CreateDir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
}
