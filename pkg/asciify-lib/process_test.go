package asciify

import (
	"image"
	"image/color"
	"log"
	"testing"
)

func loadImage() image.Image {
	img, err := DecodeImageFile("../../images/amogus.png")
	if err != nil {
		log.Fatal(err)
	}
	return img
}

var value color.Color

func BenchmarkSampleMid(b *testing.B) {
	img := loadImage()

	for n := 0; n < b.N; n++ {
		value = sampleMid(img, tile{0, 0, img.Bounds().Max.X, img.Bounds().Max.Y})
	}
}

func BenchmarkSampleMean(b *testing.B) {
	img := loadImage()

	for n := 0; n < b.N; n++ {
		value = sampleMean(img, tile{0, 0, img.Bounds().Max.X, img.Bounds().Max.Y})
	}
}
