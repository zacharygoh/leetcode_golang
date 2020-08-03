func fizzBuzz(n int) (result []string) {
    for i := 1; i <= n; i++ {
        if i % 3 == 0 && i % 5 != 0 {
            result = append(result, "Fizz")
        }else if i % 5 == 0 && i % 3 != 0 {
            result = append(result, "Buzz")
        }else if i % 3 == 0 && i % 5 == 0 {
            result = append(result, "FizzBuzz")
        }else{
            result = append(result, strconv.Itoa(i))
        }
    }
    return
}

/* 
Example:

n = 15,

Return:
[
    "1",
    "2",
    "Fizz",
    "4",
    "Buzz",
    "Fizz",
    "7",
    "8",
    "Fizz",
    "Buzz",
    "11",
    "Fizz",
    "13",
    "14",
    "FizzBuzz"
]
*/