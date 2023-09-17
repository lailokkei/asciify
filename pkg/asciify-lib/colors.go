package asciify

import (
	"image/color"
	"math"
)

type contrast interface {
	calculate(value float64) float64
}

type none struct {
}

func (self none) calculate(value float64) float64 {
	return value
}

type stretch struct {
}

func (self stretch) calculate(value float64) float64 {
	return math.Pow(value, 2) / 255
}

func invertValue(pixel color.Gray) color.Gray {
	const maxValue uint8 = 255
	return color.Gray{maxValue - pixel.Y}
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
