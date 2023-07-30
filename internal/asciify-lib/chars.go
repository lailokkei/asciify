package asciify

func reverseSet(set []rune) []rune {
	reversed := []rune{}
	for i := len(set) - 1; i >= 0; i-- {
		reversed = append(reversed, set[i])
	}
	return reversed
}

func getCharSet(charSetName string, invert bool) []rune {
	sets := map[string][]rune{
		"standard": []rune(" .'" + "`" + "^\",:;Il!i><~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ1OZmwqpdbkhao*#MW&8%B@$"),
		"detailed": []rune(" `.-':_,^=;><+!rc*/z?sLTv)J7(|Fi{C}fI31tlu[neoZ5Yxjya]2ESwqkP6h9d4VpOGbUAKXHm8RD#$Bg0MNWQ%&@"),
		"simple":   []rune(" .:-=+*#%@"),
		"blocks":   []rune(" ░▒▓█"),
		"binary":   []rune(" #"),

		"blocks-binary":    []rune("░█"),
		"helvetica-blocks": []rune("┌░▒▓█"),
		"blocks-vertical":  []rune("▏▎▍▌▋▊▉█"),
	}

	charSet := sets[charSetName]

	if invert {
		charSet = reverseSet(charSet)
	}

	return charSet
}
