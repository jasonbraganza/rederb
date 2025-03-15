package createFeedEngine

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFeed(url string, path string) {
	rawPath := path
	rawUrl := url
	fmt.Println(rawUrl)
	var fullPath string

	if rawPath == "./" {
		fullPath, _ = os.Getwd()
	} else {
		fullPath, _ = filepath.Abs(rawPath)
	}
	fmt.Println(fullPath)
}
