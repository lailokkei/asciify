package asciify

import (
	"image/color"
	"math"
)

func squash(value float64) float64 {
	return math.Pow(value, 3) / 255 / 255
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
