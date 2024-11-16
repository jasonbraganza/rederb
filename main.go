package main

import (
	"fmt"
	"github.com/dhowden/tag"
	"github.com/joho/godotenv"
	//"github.com/dhowden/tag"
	"log"
	"os"
)

// The original Python file was just an imperative set of instructions
// Since I cannot replicate that here, I hope to just build a main function with calls to
// what I think are obvious disparate functions

// I need a type to hold all my metadata. copied and adapted from the tag website
// MetadataInfo holds all possible metadata fields
type MetadataInfo struct {
	Format      tag.Format
	FileType    tag.FileType
	Title       string
	Album       string
	Artist      string
	AlbumArtist string
	Composer    string
	Genre       string
	Year        int
	Track       int
	TotalTracks int
	Disc        int
	TotalDiscs  int
	Lyrics      string
	Comment     string
	Picture     *PictureInfo
}

type PictureInfo struct {
	MIMEType    string
	Description string
	Size        int
}

// This brings everything together
func main() {
	// First things first. Load up any variables you need from environment files.
	// Put in a comma seperated list of env files to load
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment file(s)")
	}
	my_name := os.Getenv("NAME")
	artist1 := os.Getenv("ARTIST1")
	fmt.Printf("Testing env loading: \n")
	fmt.Printf("My name is %s and my favourite modern musician is %s\n", my_name, artist1)

	//trying to load metadata from files
	if len(os.Args) != 2 {
		fmt.Println("Usage: program <audio-file-path>")
		os.Exit(1)
	}

	filePath := os.Args[1]

	fileMetadata, err := getAudioFileMetadata(filePath)
	if err != nil {
		fmt.Printf("Error extracting metadata: %v\n", err)
		os.Exit(1)
	}

	prettyPrint(fileMetadata)

}

func getAudioFileMetadata(filePath string) (*MetadataInfo, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %w", err)
	}
	defer file.Close()

	// Read metadata from file
	m, err := tag.ReadFrom(file)
	if err != nil {
		return nil, fmt.Errorf("error reading metadata: %w", err)
	}

	// Extract track and disc numbers
	track, totalTracks := m.Track()
	disc, totalDiscs := m.Disc()

	// Create metadata info struct
	info := &MetadataInfo{
		Format:      m.Format(),
		FileType:    m.FileType(),
		Title:       m.Title(),
		Album:       m.Album(),
		Artist:      m.Artist(),
		AlbumArtist: m.AlbumArtist(),
		Composer:    m.Composer(),
		Genre:       m.Genre(),
		Year:        m.Year(),
		Track:       track,
		TotalTracks: totalTracks,
		Disc:        disc,
		TotalDiscs:  totalDiscs,
		Lyrics:      m.Lyrics(),
		Comment:     m.Comment(),
	}

	// Extract picture info if available
	if pic := m.Picture(); pic != nil {
		info.Picture = &PictureInfo{
			MIMEType:    pic.MIMEType,
			Description: pic.Description,
			Size:        len(pic.Data),
		}
	}

	return info, nil
}

func prettyPrint(info *MetadataInfo) {
	fmt.Println("=== Audio File Metadata ===")
	fmt.Printf("Format: %v\n", info.Format)
	fmt.Printf("File Type: %v\n", info.FileType)

	// Basic metadata
	if info.Title != "" {
		fmt.Printf("Title: %s\n", info.Title)
	}
	if info.Album != "" {
		fmt.Printf("Album: %s\n", info.Album)
	}
	if info.Artist != "" {
		fmt.Printf("Artist: %s\n", info.Artist)
	}
	if info.AlbumArtist != "" {
		fmt.Printf("Album Artist: %s\n", info.AlbumArtist)
	}
	if info.Composer != "" {
		fmt.Printf("Composer: %s\n", info.Composer)
	}
	if info.Genre != "" {
		fmt.Printf("Genre: %s\n", info.Genre)
	}
	if info.Year != 0 {
		fmt.Printf("Year: %d\n", info.Year)
	}

	// Track information
	if info.Track != 0 {
		if info.TotalTracks != 0 {
			fmt.Printf("Track: %d/%d\n", info.Track, info.TotalTracks)
		} else {
			fmt.Printf("Track: %d\n", info.Track)
		}
	}

	// Disc information
	if info.Disc != 0 {
		if info.TotalDiscs != 0 {
			fmt.Printf("Disc: %d/%d\n", info.Disc, info.TotalDiscs)
		} else {
			fmt.Printf("Disc: %d\n", info.Disc)
		}
	}

	// Additional metadata
	if info.Lyrics != "" {
		fmt.Printf("Lyrics: %s\n", info.Lyrics)
	}
	if info.Comment != "" {
		fmt.Printf("Comment: %s\n", info.Comment)
	}

	// Picture information
	if info.Picture != nil {
		fmt.Println("\n=== Picture Information ===")
		fmt.Printf("MIME Type: %s\n", info.Picture.MIMEType)
		fmt.Printf("Description: %s\n", info.Picture.Description)
		fmt.Printf("Size: %d bytes\n", info.Picture.Size)
	}
}
