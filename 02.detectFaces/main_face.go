package main

import (
	"context"
	"fmt"
	"log"
	"os"

	vision "cloud.google.com/go/vision/apiv1"
	"cloud.google.com/go/vision/v2/apiv1/visionpb"
	"google.golang.org/api/option"

	"github.com/fogleman/gg"
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

	// Detects faces in the image
	faces, err := client.DetectFaces(ctx, image, nil, 10)
	if err != nil {
		log.Fatalf("Failed to detect faces: %v", err)
	}

	// 画像の読み込み
	img, _ := gg.LoadImage(fileName)
	dc := gg.NewContextForImage(img)

	if len(faces) == 0 {
		fmt.Println("No faces found.")
		return
	} else {
		fmt.Printf("%d人が見つかりました in %s:\n", len(faces), fileName)
		for _, face := range faces {
			// 顔の位置を取得
			x := float64(face.BoundingPoly.Vertices[0].X)
			y := float64(face.BoundingPoly.Vertices[0].Y)
			w := float64(face.BoundingPoly.Vertices[2].X - face.BoundingPoly.Vertices[0].X)
			h := float64(face.BoundingPoly.Vertices[2].Y - face.BoundingPoly.Vertices[0].Y)

			// 顔の位置に四角を描画
			dc.DrawRectangle(x, y, w, h)
			dc.SetRGB(1, 0, 0)
			dc.SetLineWidth(2)
			dc.Stroke()

			// 顔のランドマークを取得
			for _, landmark := range face.Landmarks {
				// 左目の位置を取得
				if landmark.Type == visionpb.FaceAnnotation_Landmark_LEFT_EYE {
					xle := float64(landmark.Position.X)
					yle := float64(landmark.Position.Y)

					// 目の位置に四角を描画
					dc.DrawCircle(xle, yle, 5)
					dc.SetRGB(0, 1, 0)
					dc.SetLineWidth(2)
					dc.Stroke()
				}

				// 右目の位置を取得
				if landmark.Type == visionpb.FaceAnnotation_Landmark_RIGHT_EYE {
					xre := float64(landmark.Position.X)
					yre := float64(landmark.Position.Y)

					// 目の位置に四角を描画
					dc.DrawCircle(xre, yre, 5)
					dc.SetRGB(0, 0, 1)
					dc.SetLineWidth(2)
					dc.Stroke()
				}

				// 鼻の位置を取得
				if landmark.Type == visionpb.FaceAnnotation_Landmark_NOSE_TIP {
					xnt := float64(landmark.Position.X)
					ynt := float64(landmark.Position.Y)

					// 鼻の位置に四角を描画
					dc.DrawCircle(xnt, ynt, 5)
					dc.SetRGB(2, 0, 0)
					dc.SetLineWidth(2)
					dc.Stroke()
				}

				// 口の位置を取得
				if landmark.Type == visionpb.FaceAnnotation_Landmark_MOUTH_CENTER {
					xmc := float64(landmark.Position.X)
					ymc := float64(landmark.Position.Y)

					// 口の位置に四角を描画
					dc.DrawCircle(xmc, ymc, 5)
					dc.SetRGB(0, 2, 0)
					dc.SetLineWidth(2)
					dc.Stroke()
				}
			}

			// 顔の位置に
			dc.DrawString(fmt.Sprintf("emotions: %s", face.JoyLikelihood), x, y-10)
		}
	}
	// 画像の保存
	dc.SavePNG("images/out/" + photo + ".png")
}
