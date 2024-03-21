package sublist

// Relation type is defined in relations.go file.

func Sublist(l1, l2 []int) Relation {
	if isEqual(l1, l2) {
		return RelationEqual
	} else if isSubList(l1, l2) {
		return RelationSublist
	} else if isSuperList(l1, l2) {
		return RelationSuperlist
	} else {
		return RelationUnequal
	}
}

func isEqual(l1, l2 []int) bool {
	if len(l1) != len(l2) {
		return false
	}

	for i := range l1 {
		if l1[i] != l2[i] {
			return false
		}
	}

	return true
}

func isSubList(l1, l2 []int) bool {

	equals := false
	smaller := len(l1)
	bigger := len(l2)
	if smaller > bigger {
		return false
	}
	if smaller == 0 {
		return true
	}

	for i := 0; i <= bigger-smaller; i++ {
		for j := 0; j < smaller; j++ {
			if l1[j] == l2[j+i] {
				equals = true
			} else {
				equals = false
				break
			}
		}
		if equals {
			return true
		}
	}

	return false
}

func isSuperList(l1, l2 []int) bool {
	return isSubList(l2, l1)
}
