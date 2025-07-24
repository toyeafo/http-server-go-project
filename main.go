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
	defer f.Close()

	var line string

	for {
		textBytes := make([]byte, 8)
		n, err := f.Read(textBytes)

		if err != nil {
			if err.Error() == "EOF" {
				os.Exit(0)
			}
			log.Fatalf("error reading file: %v", err)
		}
		fmt.Printf("read: %s\n", string(textBytes[:n]))
	}

}
