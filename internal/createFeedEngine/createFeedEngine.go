package createFeedEngine

import (
	"fmt"
	"github.com/dhowden/tag"
	"github.com/gorilla/feeds"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"rederb/internal/rederbStructures"
	"sort"
	"strings"
	"time"
)

/* ------------------------------------------------------------------------- */

/* ------------------------------------------------------------------------- */

func CreateFeed(rawUrl string, rawPath string) {
	//rawPath := path
	//rawUrl := url
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

		// Build a map/dict of audiofile objects using tag
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

		// Setup a feed
		// get author/email from Viper. Dumping into a struct to prevent repeated Viper calls
		feedAuthorDetails := rederbStructures.FeedMeta{
			AuthorName:  viper.GetString("author_name"),
			AuthorEmail: viper.GetString("author_email"),
		}
		// Grab one audio file to pull metadata for the feed’s main details
		feedMetaDetails := dictOfAudioEntriesWithTags[sortedAudioDictKeys[0]]

		// Grab cover image from metadata and write to file
		imgData := feedMetaDetails.Picture().Data
		coverArt := filepath.Join(fullPath, "cover.jpg")
		err = os.WriteFile(coverArt, imgData, 0644)
		if err != nil {
			fmt.Println("There is no cover art")
		}
		coverArtUrl := fmt.Sprint(PodcastUrl + "/cover.jpg")

		// Dummy time object that I’ll keep incrementing by a second to add new entries
		samayMayaHai := time.Now()

		// Instantiate a new feed and start pulling details in
		feed := &feeds.Feed{
			Title:       feedMetaDetails.Album(),
			Link:        &feeds.Link{Href: PodcastUrl},
			Description: feedMetaDetails.Lyrics(),
			Author:      &feeds.Author{Name: feedAuthorDetails.AuthorName, Email: feedAuthorDetails.AuthorEmail},
			Created:     samayMayaHai,
			Image:       &feeds.Image{Url: coverArtUrl},
		}

		// Now add episodes to the feed by looping through our keys list and audiofile dict
		counter := time.Minute
		for _, audioEntry := range sortedAudioDictKeys {
			audioData := dictOfAudioEntriesWithTags[audioEntry]
			feedEntry := &feeds.Item{
				Title:       audioData.Title(),
				Link:        &feeds.Link{Href: PodcastUrl},
				Description: audioData.Lyrics(),
				Created:     samayMayaHai.Add(counter),
			}
			feed.Items = append(feed.Items, feedEntry)
			counter += counter + 5
		}

		// Write the feed to a file and call it a day
		feedFilePath := filepath.Join(fullPath, "feed.xml")
		feedFile, err := os.Create(feedFilePath)
		if err != nil {
			fmt.Println("Could not create feed.xml")
			os.Exit(1)
		}
		defer feedFile.Close()
		feed.WriteRss(feedFile)

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

	return rawAudioFileObjectMap
}

/* ------------------------------------------------------------------------- */
