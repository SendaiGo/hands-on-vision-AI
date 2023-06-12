package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"github.com/fogleman/gg"
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

	// clop hints
	fmt.Println("Crop Hints")
	crop, err := client.CropHints(ctx, image, nil)
	if err != nil {
		log.Fatalf("Failed to detect crop hints: %v", err)
	}

	// 画像の読み込み
	img, _ := gg.LoadImage(fileName)
	dc := gg.NewContextForImage(img)

	fmt.Println(crop)
	for _, hint := range crop.CropHints {
		for _, vertex := range hint.BoundingPoly.Vertices {
			x := float64(vertex.X)
			y := float64(vertex.Y)
			w := float64(vertex.X - vertex.X)
			h := float64(vertex.Y - vertex.Y)

			// 顔の位置に四角を描画
			dc.DrawRectangle(x, y, w, h)
			dc.SetRGB(1, 0, 0)
			dc.SetLineWidth(2)
			dc.Stroke()
		}
	}

	dc.SavePNG("images/" + photo + "_crop.png")
}
