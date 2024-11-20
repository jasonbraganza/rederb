package main

import (
	"fmt"
	"github.com/spf13/viper"
	"os/user"
	"path/filepath"
	"strings"

	//"github.com/dhowden/tag"
	"log"
	"os"
)

// The original Python file was just an imperative set of instructions
// Since I cannot replicate that here, I hope to just build a main function with calls to
// what I think are obvious disparate functions

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
	BaseUrl := viper.GetString("BASE_URL")
	fmt.Printf("BASE_URL = %s\n", BaseUrl)
	AuthorName := viper.Get("AUTHOR_NAME")
	fmt.Printf("Feed author name: %v\n", AuthorName)
	SubPath := os.Getenv("SUB_PATH")
	if strings.TrimSpace(SubPath) == "" {
		SubPath = viper.GetString("SUB_PATH")
	}
	FeedUrl := ""
	if SubPath != "" {
		FeedUrl = fmt.Sprintf("%s/%s", BaseUrl, SubPath)
	} else {
		FeedUrl = BaseUrl
	}
	fmt.Printf("Feed URL is: %s\n", FeedUrl)

	// 3. trying to load metadata from files
	if len(os.Args) != 2 {
		fmt.Println("Usage: program <audio-file-path>")
		os.Exit(1)
	}

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

//func getAudioFileMetadata(filePath string) (*MetadataInfo, error) {
//	file, err := os.Open(filePath)
//	if err != nil {
//		return nil, fmt.Errorf("Error opening file: %w", err)
//	}
//	defer file.Close()

//func extractCoverArt(info *MetadataInfo, audioPath string) error {
//
//}
