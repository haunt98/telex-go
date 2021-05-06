package telex

var allowedPlusRunes = map[rune]struct{}{
	'a': {},
	'd': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'y': {},
}

func ConvertText(text string) string {
	runes := []rune(text)
	viChars := make([]viChar, 0, len(runes))
	viCharIndex := -1

	// Keep track on rune to plus
	lastAllowedPlusRune := -1

	for _, r := range runes {
		if !isChar(r) {
			viChars = append(viChars, viChar{
				main: r,
			})
			viCharIndex++
			lastAllowedPlusRune = -1
			continue
		}

		// Try to plus r
		if lastAllowedPlusRune != -1 {
			if ok := viChars[lastAllowedPlusRune].plus(r); ok {
				continue
			}
		}

		// Can not plus r

		viChars = append(viChars, viChar{
			main: r,
		})
		viCharIndex++

		if _, ok := allowedPlusRunes[r]; ok {
			lastAllowedPlusRune = viCharIndex
		}

	}

	resultRunes := make([]rune, 0, len(viChars))
	for _, c := range viChars {
		resultRunes = append(resultRunes, c.toRune())
	}

	return string(resultRunes)
}

func isChar(r rune) bool {
	return (r >= 'A' && r <= 'Z') || (r >= 'a' && r <= 'z')
}
