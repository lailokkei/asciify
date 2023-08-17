package asciify

import (
	"fmt"
	"image"
	"image/color"
	"sort"
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
	return Options{"standard", false, 20, ""}
}

	total := 0
	for y := tile.y; y < tile.y+tile.height; y++ {
		for x := tile.x; x < tile.x+tile.width; x++ {
			total += int(colorToGray(img.At(x, y)).Y)
		}
	}
	return color.Gray{uint8(total / (tile.width * tile.height))}
}

func sampleMedian(img image.Image, tile tile) color.Color {
	values := []uint8{}
	for y := tile.y; y < tile.y+tile.height; y++ {
		for x := tile.x; x < tile.x+tile.width; x++ {
			values = append(values, uint8(colorToGray(img.At(x, y)).Y))
		}
	}
	less := func(i int, j int) bool {
		if i < j {
			return true
		}
		return false
	}
	sort.Slice(values, less)
	return color.Gray{values[len(values)/2]}
}

func sampleTopLeft(img image.Image, tile tile) color.Color {
	return colorToGray(img.At(tile.x, tile.y))
}

func ImageToText(img image.Image, options Options) string {
	charSet := getCharSet(options.CharSetName, options.Invert)
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
			// c := colorToGray(img.At(x, y))
			c := colorToGray(sampleMid(img, tile{x, y, tileWidth, tileHeight}))
			textImage += grayToChar(c, charSet)
		}
		textImage += "\n"
	}

	return textImage
}
