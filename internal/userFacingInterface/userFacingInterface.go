package userFacingInterface

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
	"rederb/internal/rederbStructures"
)

/*
This set of functions constructs a url for the feed
*/

func GetNewPodcastFeedRawUrl() string {

	// construct a feed structure getting values from config file
	feedMetaData := rederbStructures.FeedMeta{
		AuthorName:  viper.GetString("author_name"),
		AuthorEmail: viper.GetString("author_email"),
		BaseUrl:     viper.GetString("base_url"),
		FeedUrl:     viper.GetString("feed_url"),
		SubUrlSlice: viper.GetStringSlice("sub_url"),
	}

	// Get category from user and construct a url
	feedCategory := showFeedCategories(&feedMetaData)
	// Set URL to baseURL
	url := feedMetaData.BaseUrl
	if feedCategory == "none" {
		return url
	} else {
		url = fmt.Sprint(feedMetaData.BaseUrl, "/", feedCategory, "/")
		return url
	}
}

/* ---------------------------------------------------------------------- */

/*
This set of functions, shows a list of feed categories
and returns the selected one. If a new category is added,
it offers to save it to the configuration file.
*/

func showFeedCategories(s *rederbStructures.FeedMeta) string {
	// Main function that calls the rest in this section
	// and returns the category

	feedCategories := s.SubUrlSlice

	// Get length of current category list to track changes
	originalCategoryListLength := len(feedCategories)

	selectedCategory := gaFeedCategories(&feedCategories)

	// Use Viper to write changes to config file, if user says so
	if len(feedCategories) != originalCategoryListLength {
		writeConfigToFile(&feedCategories)
	}

	// Go back to the caller
	return selectedCategory
}

func gaFeedCategories(categoryList *[]string) string {
	// Get, Add Feed Categories
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    "Select a Feed category",
			Items:    *categoryList,
			AddLabel: "Add another category",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			*categoryList = append(*categoryList, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
	}

	return result

}

func writeConfigToFile(categoryList *[]string) {
	// Get Viper to write a changed config to disk
	fmt.Println(*categoryList)

	prompt := promptui.Prompt{
		Label:     "The feed category list has changed. Write changes to config?",
		IsConfirm: true,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Config file unchanged\n")
		return
	}

	if result == "y" {
		viper.Set("sub_url", *categoryList)
		viper.WriteConfig()
	}
}
