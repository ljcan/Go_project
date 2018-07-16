package main

import "fmt"

//#####################go切片操作###############################

/**
	更新切片
 */
func updateSlice(a []int){	//[]int表示传递切片
	a[0]=100
}
func printSclie(a []int){
	fmt.Printf("%d,%d",len(a),cap(a))
	fmt.Println()
}

func main() {
	//切片是前闭后开区间
	arr1:=[...]int{0,1,2,3,4,5,6}
	arr2:=arr1[2:6]
	arr3:=arr1[2:]
	arr4:=arr1[:6]
	arr5:=arr1[:]
	fmt.Println("arr2=",arr2)
	fmt.Println("arr3=",arr3)
	fmt.Println("arr4=",arr4)
	fmt.Println("arr5=",arr5)

	fmt.Println("update Slice arr2:")
	updateSlice(arr2)			//切片监控着数组的底层，因此它会改变数组的值
	fmt.Println("arr2=",arr2)
	fmt.Println("arr1=",arr1)

	//切片还可以再切片
	fmt.Println("reslice:")
	rarr2:=arr2[3:4]
	fmt.Println("rarr2",rarr2)

	//实际上一个切片底层维护这三个变量，一个是指向切片第一个元素的指针pre，
	//一个是切片的长度len，还有一个是切片底层的数组的从切片头开始到尾部的长度cap
	fmt.Println("extension slice:")
	fmt.Println(arr2)
	fmt.Printf("arr2=%v, len(arr2)=%d, cap(arr2)=%d\n",
		arr2,len(arr2),cap(arr2))
	earr2:=arr2[3:5]
	fmt.Println(earr2)
	fmt.Printf("earr2=%v, len(earr2)=%d, cap(earr2)=%d\n",
		earr2,len(earr2),cap(earr2))

	//当向slice中添加元素的时候，如果slice的元素的长度超过了cap的长度，
	// 那么go语言底层会扩展一个更大的slice来接之前的slice，将元素拷贝到新元素中
	//向sclie中添加元素时使用append，不过该方法中需要一个变量来接返回值
	//当slice中长度超过cap长度时，会扩展一个是之前cap长度二倍的slice来接
	arr2=append(arr2,10)
	fmt.Printf("%v,%d,%d",arr2,len(arr2),cap(arr2))
	fmt.Println()
	arr2=append(arr2,11)
	fmt.Printf("%v,%d,%d",arr2,len(arr2),cap(arr2))
	fmt.Println()
	arr2=append(arr2,12)
	fmt.Printf("%v,%d,%d",arr2,len(arr2),cap(arr2))
	fmt.Println()
	arr2=append(arr2,13)
	fmt.Printf("%v,%d,%d",arr2,len(arr2),cap(arr2))
	fmt.Println()
	arr2=append(arr2,14)
	fmt.Printf("%v,%d,%d",arr2,len(arr2),cap(arr2))
	fmt.Println()
	arr2=append(arr2,15)
	fmt.Printf("%v,%d,%d",arr2,len(arr2),cap(arr2))
	fmt.Println()
	arr2=append(arr2,16)
	fmt.Printf("%v,%d,%d",arr2,len(arr2),cap(arr2))
	fmt.Println()
	//########################创建slice###################3
	//除了之前从数组中读取slice外，还有其他几种创建slice的方法
	fmt.Println("creating slice")
	var slice1 []int    //Zero value for slice is nil
	for i:=0;i<100;i++{
		printSclie(slice1)
		slice1=append(slice1,2*i+1)
	}
	fmt.Println(slice1)

	slice2:=[]int{1,2,3,4,5}
	fmt.Println(slice2)

	//也可使用make函数来指定len和cap来创建slice

	slice3:=make([]int,5)
	fmt.Println(slice3)
	slice4:=make([]int,5,6)
	fmt.Println(slice4)

	//拷贝slice，使用copy函数
	fmt.Println("copy slice")
	slice3[4],slice4[2]=100,100
	copy(slice3,slice4)
	fmt.Println(slice3)

	//删除一个元素
	fmt.Println("delete element")
	fmt.Println(arr2)
	//删除第三个元素
	arr2=append(arr2[:2],arr2[3:]...)  //因为append的第二个参数是可变参数，所以后面需要加三个点
	fmt.Println(arr2)

	//删除第一个元素
	fmt.Println("Popping element from front")
	front:=arr2[0]
	arr2=arr2[1:]
	fmt.Println("front=",front)
	fmt.Println("arr2=",arr2)

	//删除最后一个元素
	fmt.Println("Popping element form back")
	last:=arr2[len(arr2)-1]
	arr2=arr2[:len(arr2)-1]
	fmt.Println("last=",last)
	fmt.Println("arr2=",arr2)


	}
