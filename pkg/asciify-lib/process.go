package asciify

import (
	"image"
)

type Options struct {
	CharSetName  string
	Invert       bool
	ScaleWidth   int
	SampleMethod string
}

type tile struct {
	x      int
	y      int
	width  int
	height int
}

func NewOptions() Options {
	return Options{
		"standard",
		false,
		20,
		"mid",
	}
}

func ImageToText(img image.Image, options Options) (string, error) {
	var textImage string
	charSet, err := getCharSet(options.CharSetName)

	if err != nil {
		return textImage, err
	}

	if options.ScaleWidth <= 0 {
		return textImage, nil
	}

	tileWidth := img.Bounds().Max.X / options.ScaleWidth

	if tileWidth <= 0 {
		tileWidth = 1
	}

	tileHeight := tileWidth * 2

	sampleMethod, err := getSampleFunc(options.SampleMethod)

	if err != nil {
		return textImage, err
	}

	for y := img.Bounds().Min.Y; y+tileHeight <= img.Bounds().Max.Y; y += tileHeight {
		for x := img.Bounds().Min.X; x+tileWidth <= img.Bounds().Max.X; x += tileWidth {
			tile := tile{x, y, tileWidth, tileHeight}
			textImage += charSet.tileToChar(img, tile, options.Invert, sampleMethod)
		}
		textImage += "\n"
	}

	return textImage, nil
}
