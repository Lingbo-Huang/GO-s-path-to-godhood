package main

import "fmt"

type MyWriter interface {
	Write(string) error
}

type MyCloser interface {
	Close() error
}

type writerCloser struct {
	MyWriter //interface也是一个类型
}

type fileWriter struct {
	filepath string
}

type databaseWriter struct {
	db   string
	host string
	port string
}

//func (wc *writerCloser) Write(string) error {
//	fmt.Println("writecloser write")
//	return nil
//}
//
//func (wc *writerCloser) Close() error {
//	fmt.Println("close")
//	return nil
//}

func (fw *fileWriter) Write(string) error {
	fmt.Println("write string to file ")
	return nil
}

func (dw *databaseWriter) Write(string) error {
	fmt.Println("write something to db ")
	return nil
}

func main() {
	//var mw MyWriter = &writerCloser{}
	//var mc MyCloser = &writerCloser{}
	//mw.Write("body")
	//mc.Close()

	var mw1 MyWriter = &writerCloser{
		&fileWriter{},
	}
	mw1.Write("h")
}
