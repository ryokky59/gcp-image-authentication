package main

import (
	vision "cloud.google.com/go/vision/apiv1"
	"context"
	"log"
	"os"
)

func main() {
	ctx := context.Background()
	client, err := vision.NewImageAnnotatorClient(ctx)
	if err != nil {
		log.Println("クライアントの作成に失敗しました: ", err)
		return
	}

	filename := "bassy.png"
	file, err := os.Open(filename)
	if err != nil {
		log.Println("画像の読み込みに失敗しました: ", err)
		return
	}

	image, err := vision.NewImageFromReader(file)
	if err != nil {
		log.Println("imageの作成に失敗しました: ", err)
		return
	}

	annotetions, err := client.DetectFaces(ctx, image, nil, 10)
	if err != nil {
		log.Println("検出に失敗しました: ", err)
		return
	}

	if len(annotetions) == 0 {
		log.Println("検出できませんでした")
	} else {
		for _, annotation := range annotetions {
			log.Println("怒っている: ", annotation.GetAngerLikelihood())
			log.Println("楽しんでいる: ", annotation.GetJoyLikelihood())
			log.Println("悲しんでいる: ", annotation.GetSorrowLikelihood())
			log.Println("驚いている: ", annotation.GetSurpriseLikelihood())
		}
	}
}
