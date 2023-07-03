package multipart

import (
	"fmt"
	"io"
	"log"
	"mime"

	"github.com/gin-gonic/gin"
)

func HandleFiles(ctx *gin.Context, keys []string, rootDir string) ([]string, error) {
	var out []string
	var err error
	multipart, err := ctx.Request.MultipartReader()
	if err != nil {
		log.Println("Failed to create MultipartReader", err)
		return out, err
	}
	for {
		mimePart, err := multipart.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println("Error reading multipart section: %v", err)
			break
		}
		disposition, params, err := mime.ParseMediaType(mimePart.Header.Get("Content-Disposition"))
		if err != nil {
			log.Printf("Invalid Content-Disposition: %v", err)
			break
		}
		fmt.Print(disposition)
		fmt.Print(params)
	}
	return out, err
}
