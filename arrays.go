package main

import "fmt"
/**
	数组是值类型，因此函数中不会改变外面数组的内容
 */
func printArray(arr [3]int){
	arr[0]=100
	for i,v :=range arr{
		fmt.Println(i,v)
	}
}
/**
	可以使用指针来改变数组中的值
 */
 func printArray02(arr *[3]int){
 	arr[0]=100
	for i,v:=range arr{
		fmt.Println(i,v)
	}
 }


/**
	 go语言数组是值类型
	[10]int,[20]int是不同的类型
	调用函数传递数组是会拷贝一份数组的值进行值传递
 */
func main() {
	//定义数组的几种方式
	var arr1 [3]int
	arr2:=[3]int{4,5,6}
	arr3:=[...]int{1,2,3}
	var grad [4][5]int

	fmt.Println(arr1)
	fmt.Println(arr2)
	fmt.Println(arr3)
	fmt.Println(grad)

	//遍历数组
	for i:=0;i<len(arr2);i++ {
		fmt.Println(arr2[i])
	}
	//使用range遍历
	//在任何地方，可以使用"_"来省略变量
	for i,v :=range arr2{
		fmt.Println(i,v)
	}
	fmt.Println("==========================")
	printArray(arr2)
	fmt.Println(arr2)
	fmt.Println("**************************")
	printArray02(&arr2)
	fmt.Println(arr2)

	}
