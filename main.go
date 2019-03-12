package main

import (
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"log"
)

var (
	faceAlgorithm = "haarcascade_frontalface_default.xml"

	blue = color.RGBA{0, 0, 255, 0}
)

func main() {

	// set device 0 to use built-in webcam
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalf("unable to initialize web camera %v", err)
	}

	// close webcam when complete
	defer webcam.Close()

	img := (gocv.NewMat()) // open new image object
	//var img *gocv.Mat

	defer img.Close()

	// initial webcam window
	window := gocv.NewWindow("face detection in Go")
	defer window.Close()

	// using one classifier algorithm
	classifier := gocv.NewCascadeClassifier()
	classifier.Load(faceAlgorithm)
	defer classifier.Close()

	// Read takes as input a pointer to the matrix object
	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Printf("unable to read from web cam")
			continue
		}

		//frames returned foundface
		rects := classifier.DetectMultiScale(img)

		// draw rectangles around face
		for _, r := range rects {
			// identify region around face
			imgFace := img.Region(r)
			imgFace.Close()

			text := "Who are you?"
			// get the size of the texts
			size := gocv.GetTextSize(text, gocv.FontHersheyPlain, 3, 2)

			// anchor point for rectangle
			pt := image.Pt(r.Min.X+(r.Min.X/2)-(size.X/2), r.Min.Y-2)
			gocv.PutText(&img, text, pt, gocv.FontHersheyPlain, 3, blue, 2)
			gocv.Rectangle(&img, r, blue, 2)

		}

		window.IMShow(img)
		window.WaitKey(500)
	}
}
