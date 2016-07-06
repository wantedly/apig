package main

import (
	"os"
)

func fileExists(dir string) bool {
	_, err := os.Stat(dir)
	return err == nil
}

func mkdir(dir string) error {
	return os.Mkdir(dir, os.ModePerm)
}
