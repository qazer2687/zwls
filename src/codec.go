package main

func encode(n uint32) string {

	// zero width characters
	chars := []string{"\u200B", "\u200C", "\u200D", "\uFEFF"}

	result := ""
	// calculate base outside of loop and cast to uint32
	base := uint32(len(chars))
	for n > 0 {
		result = chars[n%base] + result
		n = n / base
	}

	return result
}
