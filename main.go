package main

import (
	"fmt"
	"os"

	"gocv.io/x/gocv"
)

func main() {
	isMirror := true

	deviceID := "0"
	if len(os.Args) >= 2 {
		deviceID = os.Args[1]
	}

	// open webcam
	webcam, err := gocv.OpenVideoCapture(deviceID)
	if err != nil {
		fmt.Printf("error: cannot open device[%v]\n", deviceID)
		return
	}
	defer webcam.Close()

	// open display window
	window := gocv.NewWindow("Webcam Preview")
	defer window.Close()

	img := gocv.NewMat()
	defer img.Close()

	if ok := webcam.Read(&img); !ok {
		fmt.Printf("error: cannot read device[%v]\n", deviceID)
		return
	}

	for {
		if ok := webcam.Read(&img); !ok {
			return
		}
		if img.Empty() {
			continue
		}

		if isMirror {
			gocv.Flip(img, &img, 1)
		}
		window.IMShow(img)

		k := window.WaitKey(10)
		if k >= 0 {
			if k == 109 { // m
				isMirror = !isMirror
			} else {
				break // press any key
			}
		}
	}
}
