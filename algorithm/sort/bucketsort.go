package sort

//桶排序

func bucketAsc(data []int) {
	bucket := map[int][]int{}
	flag := true
	level := 1
	for flag {
		flag = false
		for i := 0; i < len(data); i++ {
			tmp := data[i] % (level * 10) / level
			if tmp != 0 {
				flag = true
			}
			if _, ok := bucket[tmp]; !ok {
				bucket[tmp] = []int{}
			}
			bucket[tmp] = append(bucket[tmp], data[i])
		}
		index := 0
		for i := 0; i < 10; i++ {
			if v, ok := bucket[i]; ok {
				for _, w := range v {
					data[index] = w
					index++
				}
			}
		}
		for k, _ := range bucket {
			bucket[k] = nil
		}
		level *= 10
	}

}
