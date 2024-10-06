package sublist

func Sublist(l1, l2 []int) Relation {
	switch {
	case equalSlices(l1, l2):
		return RelationEqual
	case len(l1) < len(l2) && isSublist(l1, l2):
		return RelationSublist
	case isSublist(l2, l1):
		return RelationSuperlist
	default:
		return RelationUnequal
	}
}

func equalSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func isSublist[T comparable](a, b []T) bool {
	if len(a) == 0 {
		return true
	}
	for i := 0; i <= len(b)-len(a); i++ {
		if equalSlices(a, b[i:i+len(a)]) {
			return true
		}
	}
	return false
}
