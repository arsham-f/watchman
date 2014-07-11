package main

import (
	"fmt"
	. "gopkgs.com/magick.v1"
)

var (
	imageA                = &Image{}
	imageB                = &Image{}
	NullImage             = Image{}
	images    chan *Image = make(chan *Image, 2)
)

/*
	Blocks until there is an image available
*/
func NextImage() *Image {
	return <-images
}

/*
	Continuously keep the buffer full
*/
func DownloadImages() {

	for {

		nextFile := PopFromQueue()
		data, err := Get(nextFile)

		if HandleError(err, "Downloading file") {
			continue
		}

		go Del(nextFile)
		go AddToQueue(data)

	}
}

func AddToQueue(data []byte) {

	image, err := DecodeData(data)

	if HandleError(err, "Decoding image") {
		return
	}

	images <- image
}

/*
	Rotate images calculate mean difference
*/
func NextImageAndCompare() float64 {
	var s Stopwatch

	*imageA = *imageB
	imageB = NextImage()

	if *imageA == NullImage || *imageB == NullImage {
		return 0
	}

	s.Start()
	diff, err := imageA.Compare(imageB)
	s.Stop("Compare")

	if HandleError(err, "Comparing two images") {
		return 0
	}

	fmt.Printf("%#v", diff)
	return diff.MeanPerPixel
}
