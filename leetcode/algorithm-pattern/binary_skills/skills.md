# 二进制的技巧
## 交换两个变量的值
```
// 利用异或
a = a ^ b
b = a ^ b
a = a ^ b
```
## 检查一个数是否是2的幂
```
func isPowerOfTwo(n int) bool {
    return n > 0 && (n & (n-1)) == 0
}
```
## 计算一个整数的二进制表示中1的个数
```
func hammingWeight(n uint32) int {
    count := 0
    for n != 0 {
        // 每次将最低位的1变成0
        n = n & (n - 1)
        count++
    }
    return count
}
```
## 获取整数的绝对值
```
// 第一位是符号位，负数首位是1，正数首位是0
// 任何一个负数加上-1（全1的数）等于这个数减1
// 正数的补码和二进制表示相同
// 负数的表示是二进制补码，除了符号位（也就是绝对值）取反加1
func abs(x int) int {
    // mask: 负数全1(mask为-1)，正数全0
    mask := x >> 31
    // 实现按位取反再加1：先减1（加全1）再和全1异或
    return (x+mask)^mask
}
```
## 检查两个数的符号位是否相同
```
func hasSameSign(a, b int) bool {
    return (a ^ b) >= 0
}
```
## 找到整数的二进制表示里最低位的1
```
// -n就是对n取反加1，变负数
func lowestOneBit(n int) int {
    return n & -n
}
```
## 移除最低位的1
```
a = n & (n-1)
```

