package multipart

import (
	"mime/multipart"
	"net/http"

	"github.com/tanveerprottoy/starter-go/stdlib/pkg/file"
	"github.com/tanveerprottoy/starter-go/stdlib/pkg/httppkg"
)

func ParseMultipartForm(r *http.Request) error {
	// left shift 32 << 20 which results in 32*2^20 = 33554432
	// x << y, results in x*2^y
	return r.ParseMultipartForm(32 << 20)
}

func HandleFiles(r *http.Request, keys []string, rootDir string) ([]string, error) {
	var paths []string
	err := ParseMultipartForm(r)
	if err != nil {
		return paths, err
	}
	for _, k := range keys {
		// Retrieve the file from form data
		f, header, err := httppkg.GetFile(r, k)
		if err != nil {
			return paths, err
		}
		defer f.Close()
		p, err := file.SaveFile(f, rootDir, header.Filename)
		if err != nil {
			return paths, err
		}
		paths = append(paths, p)
	}
	return paths, nil
}

func GetFileContentType(file multipart.File) (string, error) {
	// to sniff the content type only the first
	// 512 bytes are used.
	buf := make([]byte, 512)

	_, err := file.Read(buf)

	if err != nil {
		return "", err
	}

	// the function that actually does the trick
	contentType := http.DetectContentType(buf)

	return contentType, nil
}
