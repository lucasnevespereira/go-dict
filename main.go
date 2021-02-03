package main

import (
	"fmt"
	"go-dict/dictionary"
)

func main() {
	d, err := dictionary.New("./badger")
	if err != nil {
		fmt.Printf("Dictionary Error: %v\n", err)
	}
	defer d.Close()
}
