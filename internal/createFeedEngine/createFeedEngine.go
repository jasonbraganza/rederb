package createFeedEngine

import (
	"fmt"
	"github.com/dhowden/tag"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

/* ------------------------------------------------------------------------- */

func CreateFeed(url string, path string) {
	rawPath := path
	rawUrl := url
	var fullPath string
	var err error
	// Set of extensions to validate against lower down the page
	validExtensions := map[string]bool{
		".mp3": true,
		".m4a": true,
		".m4b": true,
	}
	// to hold a list (string slice) of audio files later
	var listOfAudioEntries []string

	// a dictionary of audio files with metadata. The keys are the track numbers
	var dictOfAudioEntriesWithTags map[int]tag.Metadata

	// If current folder, get absolute path via cwd, or just abs path
	if rawPath == "./" {
		fullPath, _ = os.Getwd()
	} else {
		fullPath, err = filepath.Abs(rawPath)
		if err != nil {
			fmt.Println("Not a valid path")
			os.Exit(1)
		}
	}

	// Check if path is valid
	fullPathStat, err := os.Stat(fullPath)
	if err != nil {
		fmt.Println("Not a valid path")
		os.Exit(1)
	}

	// Build a list of audio files, after reading the path. if invalid path, exit
	// also call the build url function
	if fullPathStat.IsDir() {
		PodcastUrl := processPathAndCreateURL(fullPath, rawUrl)
		fmt.Println(PodcastUrl)

		listOfDirEntries, _ := os.ReadDir(fullPath)
		for _, file := range listOfDirEntries {
			if !file.IsDir() {
				ext := filepath.Ext(file.Name())
				if validExtensions[ext] {
					listOfAudioEntries = append(listOfAudioEntries, file.Name())
				}
			}
		}

		// Build a list of audiofile objects using tag
		dictOfAudioEntriesWithTags = buildADictOfAudioFilesWithTags(fullPath, listOfAudioEntries)

		// Sort the keys of the dict we just received.
		// The keys are track numbers. We put the keys into a list and sort them
		var sortedAudioDictKeys []int
		for key, _ := range dictOfAudioEntriesWithTags {
			sortedAudioDictKeys = append(sortedAudioDictKeys, key)
		}
		sort.Ints(sortedAudioDictKeys)

		// Now we can range over the list and call the values in the dict by each key
		for _, audioDictObject := range sortedAudioDictKeys {
			fmt.Println(audioDictObject, dictOfAudioEntriesWithTags[audioDictObject].Title())
			fmt.Println("-------------------------------------------------------------------")
		}

	} else {
		fmt.Println("Not a directory path, have you given a filename?")
		os.Exit(1)
	}

}

/* ------------------------------------------------------------------------- */

// Take the base url, get the tip of the folder path and join them together
// to get the url that will be used
func processPathAndCreateURL(fullPath string, rawUrl string) string {
	var podUrl string
	workingPath := strings.Split(fullPath, "/")
	lastBitOfPath := workingPath[len(workingPath)-1]
	podUrl = fmt.Sprint(rawUrl, lastBitOfPath)
	return podUrl
}

/* ------------------------------------------------------------------------- */

// Build a dictionary (map) of audio files with metadata
func buildADictOfAudioFilesWithTags(path string, fileNameList []string) map[int]tag.Metadata {
	rawAudioFileObjectMap := make(map[int]tag.Metadata) // Temporary unsorted dict
	for _, fileName := range fileNameList {
		fileNameWithPath := filepath.Join(path, fileName)
		fileOpened, _ := os.Open(fileNameWithPath)
		defer fileOpened.Close()
		audioFileObject, err := tag.ReadFrom(fileOpened)
		if err != nil {
			log.Fatal(err)
		}
		key, _ := audioFileObject.Track()
		rawAudioFileObjectMap[key] = audioFileObject
	}

	// Sort the keys this function to move to where I actually want to build the feed later
	// experimented with it here
	//keysToSort := make([]int, 0, len(rawAudioFileObjectMap))
	//for key := range rawAudioFileObjectMap {
	//	keysToSort = append(keysToSort, key)
	//}
	//sort.Ints(keysToSort)
	//// Iterate over the sorted keys
	//for _, k := range keys {
	//	fmt.Println(k, ":", m[k])
	//}

	return rawAudioFileObjectMap
}
