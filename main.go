package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"google.golang.org/api/option"

	"cloud.google.com/go/vision/v2/apiv1/visionpb"
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

	// Detects labels in the image file
	fmt.Println("Detect Labels")
	hint, err := client.CropHints(ctx, image, nil)
	if err != nil {
		log.Fatalf("Failed to detect crop hints: %v", err)
	}
	fmt.Println(hint)

	// Detects landmarks in the image file
	fmt.Println("Detect Landmarks")
	landmark, err := client.DetectLandmarks(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect landmarks: %v", err)
	}
	fmt.Println(landmark)

	// Detects image properties in the image file
	fmt.Println("Detect Image Properties")
	prep, err := client.DetectImageProperties(ctx, image, nil)
	if err != nil {
		log.Fatalf("Failed to detect image properties: %v", err)
	}
	fmt.Println(prep)

	// Detects faces in the image
	fmt.Println("Detect Faces")
	faces, err := client.DetectFaces(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect faces: %v", err)
	}

	fmt.Printf("Found %d face(s) in %s:\n", len(faces), fileName)
	for _, face := range faces {
		fmt.Println(face)
	}

	// Detects labels in the image file
	fmt.Println("Detect Labels")
	labels, err := client.DetectLabels(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect labels: %v", err)
	}

	fmt.Printf("Found %d labels in %s:\n", len(labels), fileName)

	// Detects Text in the image file
	fmt.Println("Detect Text")
	texts, err := client.DetectTexts(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect texts: %v", err)
	}

	for _, text := range texts {
		fmt.Print(text.Description)
	}

	// Product Search
	fmt.Println("Product Search")
	product, err := client.ProductSearch(ctx, image, nil)
	if err != nil {
		log.Fatalf("Failed to detect product search: %v", err)
	}

	fmt.Println(product)

	// Web entities
	fmt.Println("Web Entities")
	web, err := client.DetectWeb(ctx, image, nil)
	if err != nil {
		log.Fatalf("Failed to detect web: %v", err)
	}

	fmt.Println(web)

	// clop hints
	fmt.Println("Crop Hints")
	crop, err := client.CropHints(ctx, image, nil)
	if err != nil {
		log.Fatalf("Failed to detect crop hints: %v", err)
	}

	fmt.Println(crop)

	// Safe Search
	fmt.Println("Safe Search")
	safe, err := client.DetectSafeSearch(ctx, image, nil)
	if err != nil {
		log.Fatalf("Failed to detect safe search: %v", err)
	}

	fmt.Println(safe)

	// Landmark
	fmt.Println("Landmark")
	land, err := client.DetectLandmarks(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect landmarks: %v", err)
	}

	fmt.Println(land)

	// Logo
	fmt.Println("Logo")
	logo, err := client.DetectLogos(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect logos: %v", err)
	}

	fmt.Println(logo)
}

func createProduct(ctx context.Context, client *vision.ProductSearchClient, projectID, location, productID, productDisplayName string) error {
	// [START vision_product_search_create_product]
	req := &visionpb.CreateProductRequest{
		Parent: fmt.Sprintf("projects/%s/locations/%s", projectID, location),
		Product: &visionpb.Product{
			DisplayName:     productDisplayName,
			Description:     "Product description",
			ProductCategory: "apparel",
		},
		ProductId: productID,
	}

	resp, err := client.CreateProduct(ctx, req)
	if err != nil {
		return err
	}

	fmt.Println("Product name:", resp.Name)
	// [END vision_product_search_create_product]
	return nil
}
