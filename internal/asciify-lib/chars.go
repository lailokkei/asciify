package asciify

import (
	"fmt"
	"image"
)

func reverseSet(set []rune) []rune {
	reversed := []rune{}
	for i := len(set) - 1; i >= 0; i-- {
		reversed = append(reversed, set[i])
	}
	return reversed
}

type CharSet interface {
	tileToChar(image.Image, tile) string
}

func getCharSet(charSetName string) CharSet {
	if charSetName == "braille" {
		return BrailleSet{}
	}

	sets := map[string][]rune{
		"standard": []rune(" .'" + "`" + "^\",:;Il!i><~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ1OZmwqpdbkhao*#MW&8%B@$"),
		"detailed": []rune(" `.-':_,^=;><+!rc*/z?sLTv)J7(|Fi{C}fI31tlu[neoZ5Yxjya]2ESwqkP6h9d4VpOGbUAKXHm8RD#$Bg0MNWQ%&@"),
		"simple":   []rune(" .:-=+*#%@"),
		"blocks":   []rune(" ░▒▓█"),
		"binary":   []rune(" #"),
	}

	return GradientSet{sets[charSetName]}
}

type GradientSet struct {
	charGradient []rune
}

func (g GradientSet) tileToChar(img image.Image, tile tile) string {

	c := colorToGray(sampleMid(img, tile))
	return grayToChar(c, g.charGradient)
}

type BrailleSet struct {
}

func (b BrailleSet) tileToChar(img image.Image, charTile tile) string {
	//https://en.wikipedia.org/wiki/Braille_Patterns
	//all wierd numbers have just been converted from hex to dec
	brailleOffset := 10240
	positionValues := [8]int{1, 8, 2, 16, 4, 32, 64, 128}
	brailleWidth := 2
	brailleHeight := 4
	dotWidth := charTile.width / brailleWidth
	dotHeight := charTile.height / brailleHeight

	var state byte

	for y := 0; y < brailleHeight; y++ {
		for x := 0; x < brailleWidth; x++ {
			dotTile := tile{charTile.x + x*dotWidth, charTile.y + y*dotHeight, dotHeight, dotHeight}
			if colorToGray(sampleMid(img, dotTile)).Y >= 127 {
				state += byte(positionValues[y*brailleWidth+x])
			}
		}
	}

	return fmt.Sprintf("%c", brailleOffset+int(state))
}
