package file

import (
	"io/ioutil"
	"os"
)

func GetPWD() (string, error) {
	pwd, err := os.Getwd()
	return pwd, err
}

func ReadFile(pwd string, path string) ([]byte, error) {
	b, err := ioutil.ReadFile(pwd + path)
	return b, err
}
