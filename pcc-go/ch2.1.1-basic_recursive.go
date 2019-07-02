package main

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func fib(n int) int {
	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

const MAX_N = 100

var memo [MAX_N]int

func fibm(n int) int {
	if n <= 1 {
		return n
	}
	if memo[n] != 0 {
		return memo[n]
	}
	memo[n] = fibm(n-1) + fibm(n-2)
	return memo[n]
}
