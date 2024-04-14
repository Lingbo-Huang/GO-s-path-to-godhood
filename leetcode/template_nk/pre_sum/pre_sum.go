package main

// 前缀和很简单，往往输入数据会很变态，所以不能用fmt.Scan()和fmt.Println()
// 有好几种优化的输入输出

// 用封装好的bufio.NewReader(os.Stdin)和bufio.NewWriter(os.Stdout)

/*
var (
	n, q, l, r int
	in         = bufio.NewReader(os.Stdin)
	out        = bufio.NewWriter(os.Stdout)
)

func main() {
	fmt.Fscan(in, &n, &q)
	a := make([]int, n)
	ap := make([]int64, n+1)
	ap[0] = 0

	for i := 0; i < n; i++ {
		fmt.Fscan(in, &a[i])
		ap[i+1] = ap[i] + int64(a[i])
	}

	for j := 0; j < q; j++ {
		fmt.Fscan(in, &l, &r)
		fmt.Fprintln(out, ap[r]-ap[l-1])
		out.Flush()
	}
}
*/

// 用适合于整行读取的组合方法
/*
sc := bufio.NewScanner(os.Stdin)
bs := make([]byte, 20000 * 1024) //设置缓冲区的最大读取
readLine = func() (res string) {
	sc.Scan() //读一行
	l := strings.Split(sc.Text(), " ")
	var res string
	for _, s := range l {
		res += s
	}
	return
}
out = bufio.NewWriter(os.Stdout)

scanner.Buffer(bs, len(bs)) //设置缓冲区的最大读取
cur := readLine()
fmt.Fprint(out, cur)
out.Flush()
*/

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	n, p, l, r int
	sc         = bufio.NewScanner(os.Stdin) //按行扫描器
	out        = bufio.NewWriter(os.Stdout) //文件输出流（要用fmt.Fprint(out, ...)）
	bs         = make([]byte, 20000*1024)   //设置缓冲区最大读取
	readLine   = func() (res []int) {       //把读取一行的操作封装成一个匿名函数
		sc.Scan()                             //扫描器读取一行
		strs := strings.Split(sc.Text(), " ") //将读取的字符串分割成切片
		res = make([]int, len(strs))          //这一句不能遗漏，返回值是切片类型，必须要有初始化
		for i, s := range strs {              //将切片中的每个元素转换为int类型，再存入返回值切片里
			x, _ := strconv.Atoi(s)
			res[i] = x
		}
		return
	}
)

func main() {
	//最后一下给Flush出来
	defer out.Flush()
	sc.Buffer(bs, len(bs))         //设置缓冲区读取最大数量
	cur1 := readLine()             //读第一行
	ap := make([]int64, cur1[0]+1) //前缀和数组（切片）
	cur2 := readLine()             //读第二行
	for i := range cur2 {
		ap[i+1] = ap[i] + int64(cur2[i]) //求前缀和
	}
	for i := 0; i < cur1[1]-1; i++ { //求要求的区间内的数值和
		cur := readLine()
		fmt.Fprintln(out, ap[cur[1]]-ap[cur[0]-1])
	}
	//最后一组单独写是为了防止最后多个换行
	cur := readLine()
	fmt.Fprint(out, ap[cur[1]]-ap[cur[0]-1])
}
