package basic

import "fmt"

func bac(num int, a, b, c string) {

	if num == 1 {
		fmt.Println("第1个盘从" + a + "->" + c)
	} else {
		bac(num-1, a, c, b)
		fmt.Println(fmt.Sprintf("第%d个盘从%v->%v", num, a, c))
		bac(num-1, b, a, c)
	}
}

var ma map[byte][]byte
var ta []string

func init() {
	ma = make(map[byte][]byte)

	ma['0'] = []byte{}
	ma['1'] = []byte{}
	ma['2'] = []byte("abc")
	ma['3'] = []byte("def")
	ma['4'] = []byte("ghi")
	ma['5'] = []byte("jkl")
	ma['6'] = []byte("mno")
	ma['7'] = []byte("pqrs")
	ma['8'] = []byte("tuv")
	ma['9'] = []byte("wxyz")
}
func letterCombinations(digits string) []string {
	// mb := map[string]struct{}{}
	ta = []string{}
	dfs(digits, []string{})
	b := []string{}
	for _, v := range ta {
		b = append(b, string(v))
	}
	return b
}

func dfs(x string, by []string) {
	if len(x) == 0 {
		ta = append(ta, by...)
		return
	}
	a, ok := ma[x[0]]

	if !ok {
		return
	}

	tmp := []string{}
	for _, v := range a {
		flag := true
		for _, v1 := range by {
			flag = false
			tmp = append(tmp, v1+string(v))
		}
		if flag {
			tmp = append(tmp, string(v))
		}
	}
	dfs(x[1:], tmp)
	return
}
