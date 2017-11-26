package tools

//ContainsRune checks if given slice contains particular rune
func ContainsRune(runes []rune, rn rune) bool {
	for _, r := range runes {
		if r == rn {
			return true
		}
	}
	return false
}

//ContainsString checks if given slice contains particular string
func ContainsString(strs []string, str string) bool {
	for _, s := range strs {
		if s == str {
			return true
		}
	}
	return false
}
