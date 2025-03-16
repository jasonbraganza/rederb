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
	var err error
	//validExtensions := map[string]bool{
	//	".mp3": true,
	//	".m4a": true,
	//	".m4b": true,
	//}

	if rawPath == "./" {
		fullPath, _ = os.Getwd()
	} else {
		fullPath, err = filepath.Abs(rawPath)
		if err != nil {
			fmt.Println("Not a valid path")
			os.Exit(1)
		}
	}
	fullPathStat, err := os.Stat(fullPath)
	if err != nil {
		fmt.Println("Not a valid path")
		os.Exit(1)
	}
	if fullPathStat.IsDir() {
		PodcastUrl := processPathAndCreateURL(fullPath, rawUrl)
		fmt.Println(PodcastUrl)
		listOfFiles, _ := os.ReadDir(fullPath)
		for _, file := range listOfFiles {
			fmt.Println(file.Name())
		}
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
