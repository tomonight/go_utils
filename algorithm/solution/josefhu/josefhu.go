package josefhu

type circleList struct {
	root *human
	len  int
}

type human struct {
	pre, next *human
	no        int
}

//约瑟夫问题。n个人，从k开始报数,数到m的人出列，一直到所有人出列
func josephuLink(m, k, n int) []int {
	first := &human{no: 1}
	pre := first
	var root *human
	if k == 1 {
		root = first
	}
	for i := 2; i <= n; i++ {
		man := &human{no: i}
		man.pre = pre
		man.pre.next = man
		pre = man
		if i == n {
			man.next = first
			first.pre = man
		}
		if i == k {
			root = man
		}
	}

	list := &circleList{root: root, len: n}

	rl := make([]int, 0)
	i := 1
	tmp := list.root
	for true {
		i++
		tmp = tmp.next
		if list.len == 1 {
			rl = append(rl, tmp.no)
			break
		}
		if i%m == 0 {
			if list.len == 2 {
				tmp.pre = tmp.next
				tmp.next = tmp.pre
				tmp.next.pre = tmp
				tmp.next.next = tmp
			} else {
				tmp.pre.next = tmp.next
				tmp.next.pre = tmp.pre
			}
			rl = append(rl, tmp.no)
			list.len--
		}
	}
	return rl
}
