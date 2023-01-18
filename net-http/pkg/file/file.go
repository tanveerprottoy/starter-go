package file

import (
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
)

func GetPWD() (string, error) {
	return os.Getwd()
}

func ReadFile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func SaveFile(multipartFile multipart.File, rootDir string, fileName string) (string, error) {
	path := filepath.Join(".", rootDir)
	_ = os.MkdirAll(path, os.ModePerm)
	fullPath := path + "/" + fileName
	file, err := os.OpenFile(fullPath, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return "", err
	}
	defer file.Close()
	// Copy the file to the destination path
	_, err = io.Copy(file, multipartFile)
	if err != nil {
		return "", err
	}
	return fullPath, nil
}
