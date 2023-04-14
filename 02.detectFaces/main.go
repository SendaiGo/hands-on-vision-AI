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
	client, err := vision.NewImageAnnotatorClient(ctx, option.WithCredentialsFile("healthy-garage-383507-b65e6912ff13.json"))
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

	// Detects faces in the image
	faces, err := client.DetectFaces(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect faces: %v", err)
	}

	if len(faces) == 0 {
		fmt.Println("No faces found.")
		return
	} else {
		fmt.Printf("%d人が見つかりました in %s:\n", len(faces), fileName)
		for _, face := range faces {
			fmt.Println(face)
		}
	}
}
