package sort

//冒泡排序

type Bubble struct {
}

func (b *Bubble) Asc(data []int) {
	flag := false
	for i := 0; i < len(data); i++ {
		if flag {
			break
		}
		flag = true
		for j := i; j < len(data); j++ {
			if data[i] > data[j] {
				tmp := data[i]
				data[i] = data[j]
				data[j] = tmp
				flag = false
			}
		}

	}
}

func bubbleAsc(data []int) {
	flag := false
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				flag = true
				a := data[j]
				data[j] = data[j+1]
				data[j+1] = a
			}
		}
		if flag == false {
			break
		}
		flag = false
	}
}
