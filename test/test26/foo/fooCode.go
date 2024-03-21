package foo

import "fmt"

// 自定义错误类型：1.结构体 2.实现Error方法
type errNotFound struct {
	file string
}

func (e errNotFound) Error() string {
	return fmt.Sprintf("file %q not found", e.file)
}

// 把判断错误类型的逻辑提取出来作为public的函数，
// 便于使用该错误接口的人能检查代码逻辑
func IsNotFoundError(err error) bool {
	_, ok := err.(errNotFound)
	return ok
}

func Open(file string) error {
	return errNotFound{file: file}
}
