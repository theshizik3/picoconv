package core

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

func GetFileType(filename string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	buffer := make([]byte, 512)
	_, err = file.Read(buffer)
	if err != nil && err.Error() != "EOF" {
		return "", err
	}

	f := http.DetectContentType(buffer)
	f = filepath.Base(f)

	if f == "octet-stream" {
		return "", fmt.Errorf("could not recognize the file type")
	}

	return f, nil
}
