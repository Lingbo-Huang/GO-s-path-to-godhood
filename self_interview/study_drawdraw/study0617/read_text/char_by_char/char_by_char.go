package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func charByChar(file string) error {
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
			for _, x := range line {
				fmt.Println(string(x))
			}
			break
		} else if err != nil {
			fmt.Printf("error reading file %s\n", err)
			break
		}
		for _, x := range line {
			fmt.Println(string(x))
		}
	}
	return nil
}

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		fmt.Printf("usage:byWord <file1> [<file2> ... ]\n")
		return
	}
	for _, file := range flag.Args() {
		err := charByChar(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}
