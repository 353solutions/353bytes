package main

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func fileSha1(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha1.New()
	_, err = io.Copy(hash, file)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", hash.Sum(nil)), nil
}

func main() {
	hash, err := fileSha1("sha1.go")
	fmt.Println(hash, err)
}
