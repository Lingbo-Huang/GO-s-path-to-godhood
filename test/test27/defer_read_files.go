package main

import "os"

type File struct {
	path    string
	content string
}

// 第一种写法会导致大量的文件句柄在延迟调用堆栈中未被及时释放
func writeManyFiles(files []File) error {
	for _, file := range files {
		f, err := os.Open(file.path)
		if err != nil {
			return err
		}
		defer f.Close()

		_, err = f.WriteString(file.content)
		if err != nil {
			return err
		}

		err = f.Sync()
		if err != nil {
			return err
		}
	}

	return nil
}

// 在大量资源的遍历循环中嵌套一个匿名函数来处理每一个资源，这里面的defer会在本次循环内就释放句柄
func writeManyFiles2(files []File) error {
	for _, file := range files {
		if err := func() error {
			f, err := os.Open(file.path)
			if err != nil {
				return err
			}
			defer f.Close() // 将在此循环步步尾执行

			_, err = f.WriteString(file.content)
			if err != nil {
				return err
			}

			return f.Sync()
		}(); err != nil {
			return err
		}
	}

	return nil
}
