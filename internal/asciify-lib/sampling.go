package asciify

import (
	"image"
	"image/color"
	"sort"
)

type sampleMethod func(image.Image, tile) color.Color

func getSampleFunc(name string) sampleMethod {
	funcs := map[string]sampleMethod{
		"mid":  sampleMid,
		"mean": sampleMean,
		"min":  sampleMin,
		"max":  sampleMax,
	}
	return funcs[name]
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

func sampleMin(img image.Image, tile tile) color.Color {
	var minColor = color.Gray{255}
	for y := tile.y; y < tile.y+tile.height; y++ {
		for x := tile.x; x < tile.x+tile.width; x++ {
			minColor.Y = min(colorToGray(img.At(x, y)).Y, minColor.Y)
		}
	}
	return minColor
}

func sampleMax(img image.Image, tile tile) color.Color {
	var maxColor = color.Gray{255}
	for y := tile.y; y < tile.y+tile.height; y++ {
		for x := tile.x; x < tile.x+tile.width; x++ {
			maxColor.Y = max(colorToGray(img.At(x, y)).Y, maxColor.Y)
		}
	}
	return maxColor
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
