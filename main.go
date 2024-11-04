package main

import (
	"bufio"
	"fmt"
	"os"
)

type Project struct {
	name string
	path string
}

func main() {
	fmt.Println("Lets set up a new project!")
	fmt.Println("What is the name of your project?")
	reader := bufio.NewReader(os.Stdin)
	for {
		response, err := reader.ReadBytes('\n')
		if err != nil || response == nil {
			fmt.Println(err)
			continue
		}

	}
	setUpDir(response)
}

func (p Project) setUpDir() error {
	err := os.MkdirAll(p.path, 0750)
	if err != nil {
		return err
	} else {
		return nil
	}
}
