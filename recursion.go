package main

import "fmt"

/*通过递归实现阶乘*/

func factorial(n uint64) (result uint64) {
	if n > 0 {
		result = n * factorial(n-1)
		return result
	}
	return 1
}

/*通过递归实现斐波那契数列*/
func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	i := 15
	fmt.Printf("%d 的阶乘是 %d\n", i, factorial(uint64(i)))

	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t",fibonacci(i))
	}
}
