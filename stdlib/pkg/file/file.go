package file

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
)

func GetPWD() (string, error) {
	return os.Getwd()
}

func FilePathWalkDir(root string) ([]string, error) {
    var files []string
    err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if !info.IsDir() {
            files = append(files, path)
        }
        return nil
    })
    return files, err
}

func IOReadDir(root string) ([]string, error) {
    var files []string
    fileInfo, err := ioutil.ReadDir(root)
    if err != nil {
        return files, err
    }

    for _, file := range fileInfo {
        files = append(files, file.Name())
    }
    return files, nil
}

func OSReadDir(root string) ([]string, error) {
    var files []string
    f, err := os.Open(root)
    if err != nil {
        return files, err
    }
    fileInfo, err := f.Readdir(-1)
    f.Close()
    if err != nil {
        return files, err
    }

    for _, file := range fileInfo {
        files = append(files, file.Name())
    }
    return files, nil
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
