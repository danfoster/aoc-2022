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
	new := []string{}
	for _, v := range old {
		if v != key {
			new = append(new, v)
		}
	}
	return new
}
