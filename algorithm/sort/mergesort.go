package sort

func mergeAsc(data []int) []int {
	if len(data) == 1 {
		return data
	}
	d1 := mergeAsc(data[:len(data)/2])
	d2 := mergeAsc(data[len(data)/2:])

	tmp := []int{}
	index := 0
	for i, j := 0, 0; i < len(d1) || j < len(d2); {
		if i == len(d1) {
			tmp = append(tmp, d2[j:]...)
			break
		}
		if j == len(d2) {
			tmp = append(tmp, d1[i:]...)
			break
		}
		if d1[i] < d2[j] {
			tmp = append(tmp, d1[i])
			i++
		} else {
			tmp = append(tmp, d2[j])
			j++
		}
		index++
	}
	return tmp
}
