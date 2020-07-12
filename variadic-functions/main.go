package main

// Here's a function that will take an arbitrary number
// of `ints` as arguments.
func sum(nums ...int) int {
	//fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	return total
	//fmt.Println(total)
}

func main() {

	// Variadic functions can be called in the usual way
	// with individual arguments.
	sum(1, 2)
	sum(1, 2, 3)

	// If you already have multiple args in a slice,
	// apply them to a variadic function using
	// `func(slice...)` like this.
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}//原文出自【易百教程】，商业转载请联系作者获得授权，非商业请保留原文链接：https://www.yiibai.com/go/golang-variadic-functions.html


