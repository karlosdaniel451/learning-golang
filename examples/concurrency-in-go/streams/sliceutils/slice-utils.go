package sliceutils

// Report whether the slice `a` is equals to the the slice `b`, that is,
// if `a` it would be genereated from copy(a, b).
func SliceEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
