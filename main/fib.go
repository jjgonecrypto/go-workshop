package main

import "os"
import "fmt"
import "strconv"

func main() {
    until, _ := strconv.ParseInt(os.Args[1], 10, 32)
    fib(until)
}

func printArgs(arg string) {
    fmt.Println(arg)
}

func fib(until int64) {
    nums := []int64{0, 1}
    if (until > 2) {
        nums = inner(0, 1, until, nums)
    }
    for _, val := range nums {
        fmt.Println(val)
    }
}

func inner(last, i, until int64, memo []int64) []int64 {
    memo = append(memo, last + i)
    if (int64(len(memo)) < until) {
        return inner(i, last + i, until, memo)
    } else {
        return memo
    }
}
