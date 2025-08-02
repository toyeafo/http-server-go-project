package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	// f, err := os.Open("messages.txt")
	f, err := net.Listen("tcp", ":42069")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	for {
		fmt.Println("Connection has been established")
		conn, err := f.Accept()
		if err != nil {
			log.Fatal(err)
		}
		results := getLinesChannel(conn)
		for val := range results {
			fmt.Printf("%s\n", val)
		}
		fmt.Println("Connection has been closed.")
	}

}
