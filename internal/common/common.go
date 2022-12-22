package common

func Check(e error) {
	if e != nil {
		panic(e)
	}

}

func CopyMap(a *map[string]bool) *map[string]bool {
	b := make(map[string]bool)
	for k, v := range *a {
		b[k] = v
	}
	return &b
}

func RemoveFromStringSlice(old []string, key string) []string {
	new := make([]string, len(old)-1)
	var k int
	var v string
	for k, v = range old {
		if v == key {
			break
		}
	}
	// new = append(old[:k], old[k+1:]...)
	copy(new, old[:k])
	copy(new[k:], old[k+1:])
	// fmt.Printf("%v - %s = %v\n", old, key, new)
	return new
}
