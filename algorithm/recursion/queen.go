package recursion

import "fmt"

func queen(a [8]int, cur int) {
	if cur == len(a) {
		fmt.Print(a)
		fmt.Println()
		return
	}
	for i := 0; i < len(a); i++ {
		a[cur] = i
		flag := true
		for j := 0; j < cur; j++ {
			ab := i - a[j]
			temp := 0
			if ab > 0 {
				temp = ab
			} else {
				temp = -ab
			}
			if a[j] == i || temp == cur-j {
				flag = false
				break
			}
		}
		if flag {
			queen(a, cur+1)
		}
	}
}

func princess(q [8]int, n int, kk *int) {
	if n == 8 {
		fmt.Println(q)
		*kk++
		return
	}
	for i := 0; i < 8; i++ {
		q[n] = i
		if notConflict(q, n, i) {
			princess(q, n+1, kk)
		}
	}
}

func notConflict(q [8]int, n, m int) bool {
	if n == 0 {
		return true
	}
	for j := 0; j < n; j++ {
		if q[j] == m {
			return false
		}
		if q[j] == m-(n-j) || q[j] == m+(n-j) {
			return false
		}
	}
	return true
}
