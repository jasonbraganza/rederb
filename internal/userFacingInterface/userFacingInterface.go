package userFacingInterface

import (
	"errors"
	"fmt"
	"github.com/manifoldco/promptui"
	"os"
)

type promptContent struct {
	errorMsg string
	label    string
}

func promptGetInput(pc promptContent) string {
	validate := func(input string) error {
		if len(input) <= 0 {
			return errors.New("pc.errorMsg")
		}
		return nil
	}

	templates := &promptui.PromptTemplates{
		Prompt:  "{{ . }} ",
		Valid:   "{{ . | green }} ",
		Invalid: "{{ . | red }} ",
		Success: "{{ . | green }} ",
	}

	prompt := promptui.Prompt{
		Label:     pc.label,
		Templates: templates,
		Validate:  validate,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Prompt result %v\n", result)

	return result
}

func promptGetSelect(pc promptContent) string {
	items := []string{"animal", "food", "person", "object"}
	index := -1
	var result string
	var err error

	for index < 0 {
		prompt := promptui.SelectWithAdd{
			Label:    pc.label,
			Items:    items,
			AddLabel: "Other",
		}

		index, result, err = prompt.Run()

		if index == -1 {
			items = append(items, result)
		}
	}

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Input: %s\n", result)

	return result
}

func CreateNewNote() {
	wordPromptContent := promptContent{
		"Please provide a word.",
		"What word would you like to make a note of?",
	}
	word := promptGetInput(wordPromptContent)

	//definitionPromptContent := promptContent{
	//	"Please provide a definition.",
	//	fmt.Sprintf("What is the definition of the %s?", word),
	//}
	//definition := promptGetInput(definitionPromptContent)
	//
	//categoryPromptContent := promptContent{
	//	"Please provide a category.",
	//	fmt.Sprintf("What category does %s belong to?", word),
	//}
	//category := promptGetSelect(categoryPromptContent)

	fmt.Println(word)
}
