func fib(n int) (result int) {
    /* DP bottom up solution */
    if n == 0 {
        return 0
    }
    fibo := []int{0, 1}
    for i := 2; i <= n; i++ {
        res := fibo[i-1] + fibo[i-2]
        fibo = append(fibo, res)
    }
    return fibo[len(fibo)-1]
    /* Recursion + memoization */
    if _, found := memo[n]; found {
        return memo[n]
    }
    if n == 1 || n == 2 {
        result = 1
    }else{
        result = fib(n-1) + fib(n-2)
    }
    memo[n] = result
    return 
}
var memo = map[int]int{}

/* 
	Example 1:

	Input: 2
	Output: 1
	Explanation: F(2) = F(1) + F(0) = 1 + 0 = 1.

	Example 2:

	Input: 3
	Output: 2
	Explanation: F(3) = F(2) + F(1) = 1 + 1 = 2.

	Example 3:

	Input: 4
	Output: 3
	Explanation: F(4) = F(3) + F(2) = 2 + 1 = 3.
*/