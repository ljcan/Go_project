package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"os"
	"bufio"
)
//******************go语言变量******************
var (
	ff=12
	gg=45
)

/**
	1.变量定义使用var来定义，先写变量名后，后写变量类型，也可以省略变量类型让编译器自动判断
	2.可以在一行定义多个变量，中间使用逗号隔开
	3,go语言变量定义之后就必须使用，否则编译器就会报错
 */
 func initialVariable(){
 	var a,b int = 1,2
 	var c=3
 	fmt.Printf("%d %d %d",a,b,c)
 }
/**
	1.变量可以使用:=来定义，这样就省略了var关键字，不过这样的方法只能在 方法中定义
	2.变量可以在函数内定义，也可以在方法外定义，go语言没有其他语言所描述的全局变量，在函数外定义的变量是包内变量
	3.go语言还可以使用var集合来创建多个变量
 */
 func shortVariable(){
 	a := 2
 	var (
 		b=4
 		c=6
 		d=8
	)
 	fmt.Printf("%d %d %d %d %d %d",a,b,c,d,ff,gg)
 }

//***************go语言常量*****************************
/**
	go语言常量带括号定义可以当做枚举类型来使用
 */
func enums()  {
	//const(
	//	a=0
	//	b=1
	//	c=2
	//)
	const (
		a = iota	//iota可以实现自动增长，从0开始
		_		//跳过一个值
		b
		c
	)
	const (			//iota还可以参加运算
		aa=4*iota
		bb
		cc
		dd
	)
	println(a ,b ,c)	//输出0 2 3
	println(aa,bb,cc)	//输出0 4 8
}
/**
	go语言变量：
		1.和scala类似，其变量类型定义在变量名之后，不过之间用空格隔开（scala使用：隔开）
		2.编译器可以还自动推断变量的类型
		3.没有char类型，只有rune
		4.原生支持复数类型
 */
 func variables(){
 	var a int=4
 	var b="dssf44"		//可以自动推断类型
 	fmt.Println(a)
 	fmt.Println(b)
 }

 //****************go语言分支语句************************
 /**
 	if条件语句不需要括号
  */
 func branchIf(){

 		//第二种写法
 	if contents,error:=ioutil.ReadFile("abc.txt");error!=nil{
		fmt.Println(error)
	}else{
		fmt.Println("%s\n",contents)
	}
 		//第一种写法
 	//if error!=nil {
 	//	fmt.Println(error)
	//}else{
	//	fmt.Println("%s\n",contents)
	//}
 }
/**
	switch语句测试
	switch语句后面可以不加任何变量
 */
 func testSwitch(scores int) string {
 	g:=""
	 switch {
	 case scores <60:
	 	g="E"
	 case scores <70:
	 	g="D"
	 case scores <80:
	 	g="C"
	 case scores <90:
	 	g="B"
	 case scores <100:
	 	g="A"
	 default:
		 panic(fmt.Sprintf("Wrong Scores: %d",scores))

	 }
	 return  g
 }

 //###################for循环语句#######################
 /**
 	将一个十进制数转换为二进制数
  */
 func convertToBin(n int) string {
 	result := ""
 	for ;n>0;n/=2 {
 		lsb:=n%2
 		result=strconv.Itoa(lsb)+result
	}
	return result
 }

 /**
 	for循环也可以省略初始语句和结束表达式
 	即可以当做for循环来使用
  */
  func while_for(){
  	file,err:=os.Open("abc.txt")
  	if err!=nil {
  		fmt.Println(err)
	}
	scanner:=bufio.NewScanner(file)
  	for scanner.Scan(){
  		fmt.Println(scanner.Text())
	}
  }
  /**
  	for循环也可以什么表达式也不用写用来写一个死循环
   */
   func forever(){
   	for{
   		fmt.Println("shinelon")
	}
   }


func main() {
	//const name  ="shinelon"
	//print(name)
	//enums()
	//variables()
	//branchIf()
 	//fmt.Println(
 	//	testSwitch(86),
	//	testSwitch(72),
	//	testSwitch(56),
	//	testSwitch(0),
	//	testSwitch(105),
 	//	)
 	//initialVariable()
 	//shortVariable()
 	//fmt.Printf(convertToBin(13))
	//while_for()
  	forever()
   }

