package phrase_types

func Match(want, given string) bool {
	if want == Any || given == Any || want == given {
		return true
	}
	return false
}
