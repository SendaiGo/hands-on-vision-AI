package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"
)

func main() {
	photo := os.Args[1]

	// Vision API client is created with the default application credentials.
	// See https://cloud.google.com/docs/authentication/production
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile("key.json"))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// The name of the image file to annotate
	// fileName := "images/pic01-7.jpg"
	fileName := "images/" + photo

	// Reads the image file into memory
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("Failed to read file: %v", err)
	}

	// Creates a new image
	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Fatalf("Failed to create image: %v", err)
	}

	// Detect Safe Search
	fmt.Println("Detect Safe Search")
	safeSearch, err := client.DetectSafeSearch(ctx, image, nil)
	if err != nil {
		log.Fatalf("Failed to detect safe search: %v", err)
	}
	// fmt.Println(safeSearch)
	fmt.Println("Safe Search properties:")
	fmt.Println("Adult:", safeSearch.Adult)
	fmt.Println("Medical:", safeSearch.Medical)
	fmt.Println("Racy:", safeSearch.Racy)
	fmt.Println("Spoofed:", safeSearch.Spoof)
	fmt.Println("Violence:", safeSearch.Violence)
}
