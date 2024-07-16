package sort_slice

import (
	"sort"
	"testing"
)

func TestGroupAnagrams(t *testing.T) {
	// 测试集：strs 输入，expectAns 预期结果
	var dataset = []struct {
		input     []string
		expectAns [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	// 对测试集里的每一条输入传入函数计算结果，与预期结果比对
	for _, value := range dataset {
		// 计算结果
		re := GroupAnagrams(value.input)
		// 遍历答案[][]string类型，比对

		// 按一定顺序输出的就要逐个比对
		for i := 0; i < len(re); i++ {
			for j := 0; j < len(re[i]); j++ {
				if re[i][j] != value.expectAns[i][j] {
					// t.Errorf遇错不停， t.Fatalf遇错即停
					t.Errorf("expect:%v, actual:%v\n", value.expectAns, re)
				}
			}
		}

		// 对输出顺序没要求的，要排序后比对
		for i := 0; i < len(re); i++ {
			// 将字符串排序
			sRe := re[i]
			sort.Slice(sRe, func(i, j int) bool {
				return sRe[i] < sRe[j]
			})
			sEx := value.expectAns[i]
			sort.Slice(sEx, func(i, j int) bool {
				return sEx[i] < sEx[j]
			})
			// 将字符串排序后，进行比对
			for j := 0; j < len(re[i]); j++ {
				if sRe[j] != sEx[j] {
					t.Errorf("expect:%v, actual:%v\n", value.expectAns, re)
				}
			}
		}
		// go test -v
	}
}
