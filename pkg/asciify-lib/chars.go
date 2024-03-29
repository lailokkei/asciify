package asciify

import (
	"errors"
	"fmt"
	"image"
	"image/color"
)

type charSet interface {
	tileToChar(image.Image, tile, tileOptions) string
}

func getCharSet(charSetName string) (charSet, error) {
	if charSetName == "braille" {
		return brailleSet{}, nil
	}

	sets := map[string][]rune{
		"standard": []rune(" .'" + "`" + "^\",:;Il!i><~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ1OZmwqpdbkhao*#MW&8%B@$"),
		"detailed": []rune(" `.-':_,^=;><+!rc*/z?sLTv)J7(|Fi{C}fI31tlu[neoZ5Yxjya]2ESwqkP6h9d4VpOGbUAKXHm8RD#$Bg0MNWQ%&@"),
		"simple":   []rune(" .:-=+*#%@"),
		"blocks":   []rune(" ░▒▓█"),
	}

	val, ok := sets[charSetName]

	if !ok {
		return nil, errors.New("Invalid char set name")
	}

	return gradientSet{charGradient: val}, nil

}

type gradientSet struct {
	charGradient []rune
}

func (g gradientSet) tileToChar(img image.Image, tile tile, options tileOptions) string {
	value := colorToGray(options.sample(img, tile))
	value = color.Gray{uint8(options.contrast.calculate(float64(value.Y)))}
	if options.invert {
		value = invertValue(value)
	}
	return grayToChar(value, g.charGradient)
}

type brailleSet struct {
}

func binaryThreshold(value color.Gray) bool {
	const halfValue uint8 = 255 / 2
	if value.Y > halfValue {
		return true
	}
	return false
}

func (b brailleSet) tileToChar(img image.Image, charTile tile, options tileOptions) string {
	//https://en.wikipedia.org/wiki/Braille_Patterns
	//hex values coverted to decimal
	brailleOffset := 10240
	positionValues := [8]int{1, 8, 2, 16, 4, 32, 64, 128}
	brailleWidth := 2
	brailleHeight := 4
	dotWidth := charTile.width / brailleWidth
	dotHeight := charTile.height / brailleHeight

	var state byte

	for y := 0; y < brailleHeight; y++ {
		for x := 0; x < brailleWidth; x++ {
			dotTile := tile{charTile.x + x*dotWidth, charTile.y + y*dotHeight, dotWidth, dotHeight}
			value := colorToGray(options.sample(img, dotTile))
			value = color.Gray{uint8(options.contrast.calculate(float64(value.Y)))}
			if options.invert == true {
				value = invertValue(value)
			}
			if binaryThreshold(value) {
				state += byte(positionValues[y*brailleWidth+x])
			}
		}
	}

	return fmt.Sprintf("%c", brailleOffset+int(state))
}
