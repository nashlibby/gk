package gk

// slice是否包含字符串
func SliceContains(s []string, search string) bool {
	for _, a := range s {
		if a == search {
			return true
		}
	}
	return false
}
