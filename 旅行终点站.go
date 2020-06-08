package main

func destCity(paths [][]string) string {
	var ans string
	var k_v = make(map[string]string)
	for i := 0; i < len(paths); i++ {
		k_v[paths[i][0]] = paths[i][1]
	}

	for _, v := range k_v {
		if _, ok := k_v[v]; !ok { // k_v[v] 判断v是否在map里面，存在为true，不存在为false
			ans = v
		}
	}
	return ans
}

func main() {
	var value = [][]string{{"1", "2"}, {"2", "3"}}
	destCity(value)
}
