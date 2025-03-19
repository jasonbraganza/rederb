package createFeedEngine

import (
	"fmt"
	"github.com/dhowden/tag"
	"github.com/gorilla/feeds"
	"github.com/jasonbraganza/rederb/internal/rederbStructures"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
	"time"
)

/* ------------------------------------------------------------------------- */
// Custom struct, because I don’t want to bother with extending tag metadata
// This will come in handy, because I need filenames (for the url),
// filesizes for enclosures, mimetypes, as well as metadata to create the feed

type xtag struct {
	audioTags tag.Metadata
	fileName  string
	fileSize  string
	mimeType  string
}

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
	// a dictionary of audio files with metadata and fileName. The keys are the track numbers
	var dictOfAudioEntriesWithTags map[int]xtag

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
			fmt.Println(audioDictObject, dictOfAudioEntriesWithTags[audioDictObject].audioTags.Title())
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
		imgData := feedMetaDetails.audioTags.Picture().Data
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
			Title:       feedMetaDetails.audioTags.Album(),
			Link:        &feeds.Link{Href: PodcastUrl},
			Description: feedMetaDetails.audioTags.Lyrics(),
			Author:      &feeds.Author{Name: feedAuthorDetails.AuthorName, Email: feedAuthorDetails.AuthorEmail},
			Created:     samayMayaHai,
			Image:       &feeds.Image{Url: coverArtUrl},
		}

		// Now add episodes to the feed by looping through our keys list and audiofile dict
		counter := time.Second
		for _, audioEntry := range sortedAudioDictKeys {
			audioData := dictOfAudioEntriesWithTags[audioEntry]
			feedEntry := &feeds.Item{
				Title: audioData.audioTags.Title(),
				Link:  &feeds.Link{Href: fmt.Sprintf(PodcastUrl + "/" + audioData.fileName)},
				Enclosure: &feeds.Enclosure{
					Url:    PodcastUrl + "/" + audioData.fileName,
					Length: audioData.fileSize,
					Type:   audioData.mimeType,
				},
				Description: audioData.audioTags.Lyrics(),
				Created:     samayMayaHai.Add(counter),
				Id:          feeds.NewUUID().String(),
			}
			feed.Items = append(feed.Items, feedEntry)
			counter += counter + 10
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
		fmt.Println("Not a directory path, have you given a fileName?")
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

// Build a dictionary (map) of audio files with metadata and fileName
func buildADictOfAudioFilesWithTags(path string, fileNameList []string) map[int]xtag {
	rawAudioFileObjectMap := make(map[int]xtag) // Temporary unsorted dict
	for _, fileName := range fileNameList {
		var err error
		var audioFileObject xtag

		fileNameWithPath := filepath.Join(path, fileName)
		fileStat, _ := os.Stat(fileNameWithPath)
		audioFileObject.fileSize = strconv.FormatInt(fileStat.Size(), 10)
		fileExt := filepath.Ext(fileName)
		if fileExt == ".mp3" {
			audioFileObject.mimeType = "audio/mpeg"
		} else {
			audioFileObject.mimeType = "audio/x-m4a"
		}
		fileOpened, _ := os.Open(fileNameWithPath)
		defer fileOpened.Close()
		audioFileObject.fileName = fileName
		audioFileObject.audioTags, err = tag.ReadFrom(fileOpened)
		if err != nil {
			log.Fatal(err)
		}
		key, _ := audioFileObject.audioTags.Track()
		rawAudioFileObjectMap[key] = audioFileObject
	}

	return rawAudioFileObjectMap
}
