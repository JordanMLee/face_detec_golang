package main

import (
	"gocv.io/x/gocv"
	"log"
)

func main() {

	// set device 0 to use built-in webcam
	webcam, err := gocv.VideoCaptureDevice(0)
	if err != nil {
		log.Fatalf("unable to initialize web camera %v", err)
	}

	// close webcam when complete
	defer webcam.Close()


	img := gocv.NewMat() // open new image object
	defer img.Close()

	// initial webcam window
	window := gocv.NewWindow("face detection in Go")
	defer window.Close()

	// Read takes as input a pointer to the matrix object
	for {
		if ok := webcam.Read(&img); !ok || img.Empty() {
			log.Printf("unable to read from web cam")
			continue
		}

		window.IMShow(img)
		window.WaitKey(500)
	}
}