package first_bad_version

/*
假设你有 n 个版本 [1, 2, ..., n]，你想找出导致之后所有版本出错的第一个错误的版本。
你可以通过调用 bool isBadVersion(version) 接口来判断版本号 version 是否在单元测试中出错。
实现一个函数来查找第一个错误的版本。
你应该尽量减少对调用 API 的次数。
*/

func firstBadVersion(n int) int {
	start := 0
	end := n
	for start+1 < end {
		mid := start + (end-start)/2
		if isBadVersion(mid) {
			// 找第一个，如果满足条件，则end往左走
			end = mid
		} else if isBadVersion(mid) == false {
			// 不满足，start往右走
			start = mid
		}
	}
	if isBadVersion(start) {
		// 找第一个，最后怕判断start位置满不满足条件，满足就返回
		return start
	}
	return end
}

func isBadVersion(n int) bool {
	if n == 3 {
		return true
	}
	return false
}
