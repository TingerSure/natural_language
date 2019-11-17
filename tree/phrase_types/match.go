package phrase_types

func Match(want, given string) bool {
	if want == Any {
		return true
	}
	if want == given {
		return true
	}
	return false
}
