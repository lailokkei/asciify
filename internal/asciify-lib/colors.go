package asciify

import (
	"image"
	"image/color"
	"math"
	"sort"
)

func squash(value float64) float64 {
	return math.Pow(value, 3) / 255 / 255
}

func colorToGray(pixel color.Color) color.Gray {
	return color.GrayModel.Convert(pixel).(color.Gray)
}

func grayToChar(c color.Gray, characterSet []rune) string {
	levels := float64(len(characterSet))
	level := float64(c.Y) / 255.0 * levels
	if level == levels {
		level--
	}
	return string(characterSet[int(level)])
}

func sampleMid(img image.Image, tile tile) color.Color {
	return img.At(tile.x+tile.width/2, tile.y+tile.height/2)
}

func sampleMean(img image.Image, tile tile) color.Color {
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
