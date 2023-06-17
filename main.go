package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func searchText(textSaver []string) {
	var search string
	found := false

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Find your text:")
		scanner.Scan()
		search = scanner.Text()

		for _, value := range textSaver {
			if strings.Contains(value, search) {
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

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("Enter text (or 'continue' to continue): ")
		scanner.Scan()
		text = scanner.Text()

		if text == "continue" {
			break
		}

		textSaver = append(textSaver, text)
	}

	fmt.Println("Save text:", textSaver)

	searchText(textSaver)
}
