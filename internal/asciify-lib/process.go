package asciify

import (
	"image"
)

type Options struct {
	CharSetName string
	Invert      bool
	ScaleWidth  int
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
	}
}

func ImageToText(img image.Image, options Options) string {
	charSet := getCharSet(options.CharSetName)
	textImage := ""

	if options.ScaleWidth <= 0 {
		return textImage
	}

	tileWidth := img.Bounds().Max.X / options.ScaleWidth

	if tileWidth <= 0 {
		tileWidth = 1
	}

	tileHeight := tileWidth * 2

	for y := img.Bounds().Min.Y; y+tileHeight <= img.Bounds().Max.Y; y += tileHeight {
		for x := img.Bounds().Min.X; x+tileWidth <= img.Bounds().Max.X; x += tileWidth {
			tile := tile{x, y, tileWidth, tileHeight}
			textImage += charSet.tileToChar(img, tile, options.Invert)
		}
		textImage += "\n"
	}

	return textImage
}
