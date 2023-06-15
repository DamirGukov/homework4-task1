package main

import "fmt"

func searchText(textSaver []string) {
	var search string
	found := false

	for {
		fmt.Println("Find your text:")
		fmt.Scan(&search)

		for _, value := range textSaver {
			if value != "" && value == search {
				fmt.Println(value)
				found = true
			}
		}

		if search == "end" {
			break
		}
	}
	if !found {
		fmt.Println("No matches found.")
	}
}

func main() {
	textSaver := make([]string, 0)
	var text string

	for {
		fmt.Println("Enter text (or 'continue' to continue): ")
		fmt.Scan(&text)

		if text == "continue" {
			break
		}

		textSaver = append(textSaver, text)
	}

	fmt.Println("Save text:", textSaver)

	searchText(textSaver)
}
