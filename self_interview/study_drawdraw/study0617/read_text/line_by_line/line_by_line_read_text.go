package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func lineByLine(file string) error {
	var err error
	f, err := os.Open(file)
	if err != nil {
		return err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			// 最后读到文件末尾，是没有换行的，需要用 Println 加上换行
			fmt.Println(line)
			break
		} else if err != nil {
			fmt.Printf("error reading file %s\n", err)
			break
		}
		// ReadString 把分割符也读入了, 所以这里 line 自带换行
		fmt.Print(line)
	}
	return nil
}

func main() {
	/*
		os.Args[0] 通常是程序的名称或路径。
		os.Args[1] 及其后的元素是传递给程序的实际命令行参数。
		flag.Parse parses the command-line flags from os.Args[1:].
	*/
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage: byLine <file1> [<file2> .. ]\n")
		return
	}
	for _, file := range flag.Args() {
		err := lineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}

}
