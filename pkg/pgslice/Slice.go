package pgslice

// DiffSliceItem 高效的找出两个Slice中的不同元素
// 思路：1.比较输入长度大小,2.创建字典遍历大的M,3.循环小的发现不相同的放入,4.循环Map找到为1的结果返回
// 输入两个字符串切片，返回两个字符串切片中不同元素
func DiffSliceItem(s1, s2 []string) []string {
	//s1为更大的值
	maxSlc, minSlc := len(s1), len(s2)
	if maxSlc < minSlc {
		maxSlc, minSlc = minSlc, maxSlc
		s1, s2 = s2, s1
	}
	var strRes []string
	strMap := make(map[string]int, maxSlc)
	for i := 0; i < maxSlc; i++ {
		strMap[s1[i]] = 1
	}
	for i := 0; i < minSlc; i++ {
		if _, ok := strMap[s2[i]]; ok {
			strMap[s2[i]] = 2
		} else {
			strRes = append(strRes, s2[i])
		}
	}
	for key, value := range strMap {
		if value == 1 {
			strRes = append(strRes, key)
		}
	}
	return strRes
}
