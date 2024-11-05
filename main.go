package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Lets set up a new project!")

	fileDir := setUpDir()

}

func setUpDir() string {
	dirPath := askQuestion("Please provide the file path from the root directory.")
	err := os.MkdirAll(dirPath, 0700)
	if err != nil {
		fmt.Println(err)
		return setUpDir()
	}
	return dirPath
}

// Handle all inputs from the user
func askQuestion(s string) string {
	fmt.Println(s)
	scanner := bufio.NewScanner(os.Stdin)
	for {
		scanner.Scan()
		in := scanner.Text()
		if in != "" {
			return in
		} else {
			fmt.Println("Please enter a valid response")
		}

	}
}

func createFiles() {
	os.
}