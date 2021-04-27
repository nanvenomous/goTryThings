package symdiff

func SymmetricDifference(subtractFrom []string, toSubtract []string) []string {
	diff := []string{}
	for _, elem := range subtractFrom {
		if !ElementInArray(elem, toSubtract) {
			diff = append(diff, elem)
		}
	}
	return diff
}

func ElementInArray(elem string, arr []string) bool {
	for _, subelem := range arr {
		if elem == subelem {
			return true
		}
	}
	return false
}
