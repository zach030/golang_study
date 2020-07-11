package main

import "fmt"

func ifelse(num int) bool {
	// Here's a basic example.
	if num%2 == 0 {
		fmt.Printf("%v is even", num)
		return true
	} else {
		fmt.Printf("%v is odd", num)
		return false
	}
}

// Note that you don't need parentheses around conditions
// in Go, but that the braces are required.//原文出自【易百教程】，商业转载请联系作者获得授权，非商业请保留原文链接：https://www.yiibai.com/go/golang-if-else.html
func main(){
	var num = 6
	ifelse(num)
}
