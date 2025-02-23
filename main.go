package main

import (
	"bufio"
	"fmt"
	"github.com/spf13/viper"
	"os/user"
	"path"
	"path/filepath"
	"strings"

	//"github.com/dhowden/tag"
	"log"
	"os"
)

// The original Python file was just an imperative set of instructions
// Since I cannot replicate that here, I hope to just build a main function with calls to
// what I think are obvious disparate functions

// I need an interface that mimics the tag package interface to hold all my feed metadata.
// so i can reuse it across functions

//func getAudioFiles(path string) []string error {
//
//}

// This main function brings everything together
func main() {
	// read file and get metadata
	// set up a feed
	// process episodes and get them all in a row
	// write feed

	// First things first. Load up any variables you need from environment files.
	// Put in a comma seperated list of env files to load
	// figure out where go wants to put in env files

	// 1. Load in my env, using Viper.
	// find it in the rederb folder in the home/.config directory
	userHome, _ := user.Current()
	viperConfigPath := filepath.Join(userHome.HomeDir, ".config", "rederb")
	viper.AddConfigPath(viperConfigPath)
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error loading environment file(s)")
	}

	// 2. Setup the environment.
	baseUrl := viper.GetString("BASE_URL")
	AuthorName := viper.Get("AUTHOR_NAME")
	fmt.Printf("Feed author name: %v\n", AuthorName)
	subPath := os.Getenv("SUB_PATH")
	if strings.TrimSpace(subPath) == "" {
		subPath = viper.GetString("SUB_PATH")
	}
	baseFeedUrl := ""
	if subPath != "" {
		baseFeedUrl = fmt.Sprintf("%s/%s", baseUrl, subPath)
	} else {
		baseFeedUrl = baseUrl
	}
	fmt.Printf("Feed URL is: %s\n", baseFeedUrl)

	// 3. Get an audiobook folder and look for files in them

	readInPath := bufio.NewReader(os.Stdin)
	fmt.Println("Enter path to book directory: ")
	bookFolderPath, _ := readInPath.ReadString('\n')
	bookFolderName := path.Base(bookFolderPath)
	feedUrl := fmt.Sprintf("%s/%s", baseFeedUrl, bookFolderName)
	fmt.Println(feedUrl)
	sliceOfAudioFiles, err := getListofAudioFiles(bookFolderPath)
	//fmt.Println(sliceOfAudioFiles, sliceOfAudioFiles[0], len(sliceOfAudioFiles))
	fmt.Printf("%v", sliceOfAudioFiles)

	// Build a slice of filenames from the given directory

	// Load basic feed metadata from the first file
	// currently just getting all the metadata from the first file
	// and writing it to a file in the same folder as the audiobook

	//filePath := os.Args[1]
	//
	//fileMetadata, err := getAudioFileMetadata(filePath)
	//if err != nil {
	//	fmt.Printf("Error extracting metadata: %v\n", err)
	//	os.Exit(1)
	//}

	//prettyPrint(fileMetadata)

	// Get cover art from the file
	//if fileMetadata.Picture != nil {
	//	err := extractCoverArt(fileMetadata, filePath)
	//	if err != nil {
	//		fmt.Printf("Error extracting cover art: %v\n", err)
	//	}
	//} else {
	//	fmt.Printf("No cover art found in the Metadata\n")
	//}

}

func getListofAudioFiles(folderPath string) ([]string, error) {
	var audioFiles []string
	//audioExtensions := map[string]bool{
	//	".mp3": true,
	//	".m4a": true,
	//	".m4b": true,
	//}

	files, err := os.ReadDir(folderPath)
	if err != nil {
		return nil, fmt.Errorf("Error reading directory: %w", err)
	}
	fmt.Printf("Reading %d files\n", len(files))
	for _, file := range files {
		//ext := strings.ToLower(filepath.Ext(file.Name()))
		//if audioExtensions[ext] {
		//	audioFiles = append(audioFiles, file.Name())
		//}
		fmt.Println(file.Name())
		audioFiles = append(audioFiles, file.Name())
	}
	fmt.Printf("%v", audioFiles)
	return audioFiles, nil
}

//func getAudioFileMetadata(filePath string) (*MetadataInfo, error) {
//	file, err := os.Open(filePath)
//	if err != nil {
//		return nil, fmt.Errorf("Error opening file: %w", err)
//	}
//	defer file.Close()

//func extractCoverArt(info *MetadataInfo, audioPath string) error {
//
//}
