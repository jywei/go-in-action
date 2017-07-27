package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/jywei/go-in-action/chapter3/words"
)

func main() {
	filname := os.Args[1]

	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("There was an error opening the file:", err)
		return
	}

	text := string(contents)

	count := words.CountWords(text)
	fmt.Printf("There are %d words in your text.", count)
}
