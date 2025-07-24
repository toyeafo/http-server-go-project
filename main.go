package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalln(err)
	}

	results := getLinesChannel(f)
	for val := range results {
		fmt.Printf("read: %s\n", val)
	}
}
