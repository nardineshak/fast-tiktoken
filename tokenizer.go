package codec

import "github.com/nardineshak/fast-tiktoken/codec"

func Tokenize(str string) []uint {
	// need to access the c100k_base map
	if len(str) == 0 {
		return []uint{}
	}

	// rune() returns an array that contains
	// the Unicode code points of the string
	runeSlice := []rune(str)
	startIndex := 0
	endIndex := 0
	var result []uint
	// go through the array of Unicode
	// checking if substrings of the string are
	// tokens in the map, if so add them to resulting array
	// otherwise move endIndex by +1
	for pos, _ := range runeSlice {
		r := runeSlice[startIndex:endIndex]
		s := string(r)
		if _, ok := codec.Cl100kBaseVocab[s]; startIndex < endIndex && ok {
			result = append(result, codec.Cl100kBaseVocab[s])
			startIndex = endIndex
			endIndex += 1
		} else {
			endIndex = pos + 1
		}
	}
	// check if you have an ending substring that is a token
	if startIndex < endIndex && endIndex <= len(str) {
		r := runeSlice[startIndex:endIndex]
		s := string(r)
		if _, ok := codec.Cl100kBaseVocab[s]; startIndex < endIndex && ok {
			result = append(result, codec.Cl100kBaseVocab[s])
		}
	}

	return result

}

func EstimateTokens(str string) int {
	// need to access the c100k_base map
	if len(str) == 0 {
		return 0
	}

	// rune() returns an array that contains
	// the Unicode code points of the string
	runeSlice := []rune(str)
	startIndex := 0
	endIndex := 0
	var result int
	// go through the array of Unicode
	// checking if substrings of the string are
	// tokens in the map, if so add them to resulting array
	// otherwise move endIndex by +1
	for pos, _ := range runeSlice {
		r := runeSlice[startIndex:endIndex]
		s := string(r)
		if _, ok := codec.Cl100kBaseVocab[s]; startIndex < endIndex && ok {
			result += 1
			startIndex = endIndex
			endIndex += 1
		} else {
			endIndex = pos + 1
		}
	}
	// check if you have an ending substring that is a token
	if startIndex < endIndex && endIndex <= len(str) {
		r := runeSlice[startIndex:endIndex]
		s := string(r)
		if _, ok := codec.Cl100kBaseVocab[s]; startIndex < endIndex && ok {
			result += 1
		}
	}

	return result
}
