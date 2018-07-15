package main

import "fmt"

func div(a int,b int) (resoult int,err string){
	if b==0 {
		err="Denominator error!!!"
		return 0,err
	}else{
		return a/b,""
	}
}
/**
   函数作为参数传递给函数，函数式编程
 */
 func sum(op func(a int,b int) int,a int,b int) int{
 	return op(a,b)
 }
 /**
 	可变参数
  */
  func sum02(numbers ...int) int{
  	s:=0
	  for i:=range numbers {
		  s+=numbers[i]
	  }
	  return s
  }
  /**
  GO语言大多数情况下是值传递
   */
   func swap(a,b int){
   	 a,b=b,a
   }
   /**
   	可以使用指针来实现引用传递
    */
    func swap02(a,b *int){
    	*a,*b=*b,*a
	}
	/**
		也可使用返回两个值来实现变量的值的交换
	 */
	 func swap03(a,b int) (int ,int){
	 	return b,a
	 }

func main() {
	resoult,err:=div(5,5)
	fmt.Printf("%d %s",resoult,err)
	//如果不想全部获取函数的返回值，只想获取其中的一个，可以使用_代替，
	// 不然go语法严格，定义了的变量必须使用，不然编译不通过
	resoult,_=div(8,2)
	fmt.Println(resoult)
	s:=sum(func(a int, b int) int {
		return a+b
	},3,4)
	fmt.Println(s)

	fmt.Println(sum02(1,2,3,4,5))

	a,b:=3,4
	swap(a,b)
	fmt.Println(a,b)		//值传递不会改变变量的值
	//swap02(&a,&b)
	fmt.Println(a,b)
	a,b=swap03(a,b)
	fmt.Println(a,b)
}
