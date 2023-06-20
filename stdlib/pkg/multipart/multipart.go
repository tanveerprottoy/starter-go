package multipart

import (
	"net/http"

	"github.com/tanveerprottoy/starter-go/stdlib/pkg/file"
	httpPkg "github.com/tanveerprottoy/starter-go/stdlib/pkg/httppkg"
)

func ParseMultipartForm(r *http.Request) (*http.Request, error) {
	// left shift 32 << 20 which results in 32*2^20 = 33554432
	// x << y, results in x*2^y
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		return r, err
	}
	return r, nil
}

func HandleFiles(r *http.Request, keys []string, rootDir string) ([]string, error) {
	var paths []string
	r, err := ParseMultipartForm(r)
	if err != nil {
		return paths, err
	}
	for _, k := range keys {
		// Retrieve the file from form data
		f, header, err := httpPkg.GetFile(r, k)
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
