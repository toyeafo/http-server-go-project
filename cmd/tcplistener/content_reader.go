package main

import (
	"fmt"
	"io"
	"log"
	"strings"
)

func getLinesChannel(f io.ReadCloser) <-chan string {
	line := make(chan string)

	go func() {
		defer close(line)
		defer f.Close()
		lineContents := ""
		for {
			textBytes := make([]byte, 8)
			n, err := f.Read(textBytes)

			if err != nil {
				if lineContents != "" {
					line <- lineContents
				}
				if err.Error() == "EOF" {
					break
				}
				log.Fatalf("error reading file: %v", err)
				return
			}

			parts := strings.Split(string(textBytes[:n]), "\n")

			for x := range parts[:len(parts)-1] {
				line <- fmt.Sprintf("%s%s", lineContents, parts[x])
				lineContents = ""
			}
			lineContents += parts[len(parts)-1]
		}
	}()
	return line
}
