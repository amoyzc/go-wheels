package string_set

// SetToSlice convert map set into slice
func SetToSlice(set map[string]struct{}) []string {
	keys := make([]string, len(set))
	i := 0
	for k := range set {
		keys[i] = k
		i++
	}
	return keys
}

// SliceToSet convert slice into map set
func SliceToSet(slice []string) map[string]struct{} {
	m := make(map[string]struct{}, len(slice))
	for _, k := range slice {
		m[k] = struct{}{}
	}
	return m
}
