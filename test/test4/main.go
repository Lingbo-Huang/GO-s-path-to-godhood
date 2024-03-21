package main

/*
变量逃逸分析
*/
// Person /*
type Person struct {
	Name string
	Age  int
}

func createPerson(name string, age int) *Person {
	p := new(Person) // 局部变量 p 逃逸到堆上
	p.Name = name
	p.Age = age
	return p
}

func main() {
	person := createPerson("Alice", 30)
	_ = person
}
