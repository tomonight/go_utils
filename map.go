package utils

//AssembleTwoMapString add map2 to map1 map[string]string
func AssembleTwoMapString(map1, map2 map[string]string) {
	if map2 == nil || len(map2) == 0 {
		return
	}
	for k, v := range map2 {
		map1[k] = v
	}
}
