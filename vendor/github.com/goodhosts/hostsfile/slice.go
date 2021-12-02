package hostsfile

func itemInSlice(item string, list []string) bool {
	for _, i := range list {
		if i == item {
			return true
		}
	}

	return false
}

func removeFromSlice(s string, slice []string) []string {
	pos := findPositionInSlice(s, slice)
	for pos > -1 {
		slice = append(slice[:pos], slice[pos+1:]...)
		pos = findPositionInSlice(s, slice)
	}
	return slice
}

func findPositionInSlice(s string, slice []string) int {
	for index, v := range slice {
		if v == s {
			return index
		}
	}
	return -1
}
