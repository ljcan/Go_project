package main

import "fmt"
/**
	go语言指针不同于C语言，指针不能进行运算
 */
func testPoint(){
	var a int=2
	var pa *int=&a
	*pa=3
	fmt.Println(a)
	fmt.Println(*pa)
}

func main() {
	testPoint()
}
