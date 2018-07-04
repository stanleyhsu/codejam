#include "foo.h"

int myadd(int a, int b) {
    return a + b;
}

int fact(int n) {
    if (0 == n) return 1;
    return n * fact(n - 1);
}

long long int fib(int n) {
    if (n <= 1) return n;
    return fib(n - 1) + fib(n - 2);
}

const int MAX_N = 100;
static long long int memo[MAX_N + 1] = {0};
long long int fibm(int n) {
    if (n <= 1) return n;
    if (memo[n] != 0) return memo[n];
    return memo[n] = fib(n - 1) + fib(n - 2);
}

int Foo::add(int a, int b) {
    return a + b;
}