package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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
			if line != "" {
				fmt.Printf("read: %s\n", line)
				line = ""
			}
			if err.Error() == "EOF" {
				break
			}
			log.Fatalf("error reading file: %v", err)
			break
		}

		parts := strings.Split(string(textBytes[:n]), "\n")
		line += strings.Join(parts, "")
		if len(parts) < 2 {
			continue
		}

		if len(parts) > 1 {
			for range parts[:len(parts)-1] {
				fmt.Printf("read: %s\n", line)
			}
		}
		line = parts[len(parts)-1]
	}

}
