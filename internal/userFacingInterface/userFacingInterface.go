package userFacingInterface

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/spf13/viper"
	"rederb/internal/rederbStructures"
)

func ShowFeedCategories() {
	feedMetaData := rederbStructures.FeedMeta{
		AuthorName:  viper.GetString("author_name"),
		AuthorEmail: viper.GetString("author_email"),
		BaseUrl:     viper.GetString("base_url"),
		FeedUrl:     viper.GetString("feed_url"),
		SubUrlSlice: viper.GetStringSlice("sub_url"),
	}

	feedCategories := feedMetaData.SubUrlSlice
	garFeedCategories(feedCategories)

}

func garFeedCategories(categoryList []string) {
	// Get, Add, Remove Feed Categories
	index := -1
	var result string
	var err error
	items := categoryList

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    "Select a Feed category",
			Items:    items,
			AddLabel: "Add another category",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	fmt.Printf("You chose %s\n", result)
}
