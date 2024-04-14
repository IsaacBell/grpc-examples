package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	entries, err := os.ReadDir("D:/my-courses/protobuf-grpc-go")
	if err != nil {
		log.Fatal(err)
	}

	for _, e := range entries {
		if e.IsDir() {
			fmt.Println(e.Name())
		}
	}
}
