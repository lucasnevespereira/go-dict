package main

import (
	"fmt"
	"go-dict/dictionary"
)

func main() {
	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()

	d.Add("python", "An interpreted language")
	words, entries, _ := d.List()
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary Error: %v\n", err)
	}
}
