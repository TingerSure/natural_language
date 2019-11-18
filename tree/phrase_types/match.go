package phrase_types

func Match(want, given string) bool {
	if want == Any || want == given {
		return true
	}
	return false
}
