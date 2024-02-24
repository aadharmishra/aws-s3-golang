package main

import (
	"awsS3Golang/initiate"
	"log"
)

func main() {
	err := initiate.Initialize()
	if err != nil {
		log.Fatalf("application crashed! error: %v\n", err)
	}
}
