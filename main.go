package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image/color"
)

func main() {
	webcam, err := gocv.OpenVideoCapture(0)
	if err != nil {
		fmt.Println("Error opening webcam:" + err.Error())
		return
	}
	defer webcam.Close()
	window := gocv.NewWindow("Detector")
	defer window.Close()
	classifier := gocv.NewCascadeClassifier()
	defer classifier.Close()
	classifier.Load("haarcascade_frontalface_default.xml")
	for {
		img := gocv.NewMat()
		if ok := webcam.Read(&img); !ok {
			fmt.Println("Cannot read from webcam")
			return
		}
		if img.Empty() {
			continue
		}
		myFace := classifier.DetectMultiScale(img)
		for _, r := range myFace {
			gocv.Rectangle(&img, r, color.RGBA{G: 255}, 2)
		}
		window.IMShow(img)
		if window.WaitKey(1) >= 0 {
			break
		}
	}
}
