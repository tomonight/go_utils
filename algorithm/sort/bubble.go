package sort

//冒泡排序

type Bubble struct {
}

func (b *Bubble) Asc(data []int) {

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
