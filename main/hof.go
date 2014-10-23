package main

import "fmt"

func main() {
	loops(15, twos)
}

func loops(c int, worker func(int, int) int) {
	for i := 0; i < c; i++ {
		fmt.Println(worker(i, i))
	}
}

func twos(a int, b int) int {
	return  a * b
}
