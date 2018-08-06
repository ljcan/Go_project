package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s :="YES我爱你！"
	fmt.Printf("%s\n",s)
	fmt.Printf("%d", len(s))
	fmt.Println()
	//rune相当于go语言的char
	for i,ch:= range []rune(s){
		fmt.Printf("(%d %c)",i,ch)
	}
	fmt.Println()
	//获取字符数
	fmt.Println(utf8.RuneCountInString(s))

	
}
