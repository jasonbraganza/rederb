package createFeedEngine

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func CreateFeed(url string, path string) {
	rawPath := path
	rawUrl := url

	var fullPath string

	if rawPath == "./" {
		fullPath, _ = os.Getwd()
	} else {
		fullPath, _ = filepath.Abs(rawPath)
	}
	fullPathStat, err := os.Stat(fullPath)
	if err != nil {
		fmt.Println("Not a valid path")
		os.Exit(1)
	}
	if fullPathStat.IsDir() {
		PodcastUrl := processPathAndCreateURL(fullPath, rawUrl)
		fmt.Println(PodcastUrl)
	} else {
		fmt.Println("Not a directory path, have you given a filename?")
	}

}

func processPathAndCreateURL(fullPath string, rawUrl string) string {
	var podUrl string
	workingPath := strings.Split(fullPath, "/")
	lastBitofPath := workingPath[len(workingPath)-1]
	podUrl = fmt.Sprint(rawUrl, lastBitofPath)
	return podUrl
}
