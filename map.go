package main

import "fmt"

func main() {
	//创建map
	fmt.Println("creating map")
	m1:=map[string]string{
		"name":"shinelon",
		"age":"19",
		"sex":"male",
	}
	fmt.Println(m1)    //map中的元素无序

	//使用make函数来创建map[key]value
	m2:=make(map[string]string)
	fmt.Println(m2)

	//创建空的map
	var m3 map[int]string
	fmt.Println(m3)

	//遍历map，可以使用range来遍历map
	fmt.Println("Traversing map")
	for k,v:=range m1 {
		fmt.Println(k,v)
	}

	//获取map中的一个元素
	fmt.Println("getting map element")
	name:=m1["name"]
	fmt.Println("name=",name)
	names:=m1["names"]       //go语言获取一个不存在的key值会返回一个空字符串
	fmt.Println(names)
	if names,ok:=m1["names"];ok{
		fmt.Println(names)
	}else{
		fmt.Println("names isn't exist!")
	}
	//map的key
	//map使用了哈希表，必须可以比较相等
	//除了slice，map，function等内建类型外其他类型都可作为key
	//Struct类型不包含上述字段，也可以作为key

}
