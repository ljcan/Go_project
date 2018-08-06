package main

import "fmt"

/**
	求字符串中最长不重复的子串
 */
func nonRepeat(s string) int {
	start:=0
	lastOccured:=make(map[rune]int)
	maxLength:=0
	for i,ch:= range []rune(s){
		if lastI,ok:=lastOccured[ch];ok&&lastI>start{
			start=lastI+1
		}
		if i-start+1>maxLength{
			maxLength=i-start+1
		}
		lastOccured[ch]=i
	}
	return maxLength
}


func main() {
	length:=nonRepeat("我爱你")
	fmt.Println(length)
}
