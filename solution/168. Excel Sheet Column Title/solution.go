func convertToTitle(n int) (result string) {
    n--
    if n < 0 {
        return
    }
    var Letters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
    result = Letters[n % 26]
    n = n / 26
    for n > 0 {
        n -= 1
        result = Letters[n % 26] + result
        n /= 26
    }
    return 
}

/* 
    Example 1:

    Input: 1
    Output: "A"

    Example 2:

    Input: 28
    Output: "AB"

    Example 3:

    Input: 701
    Output: "ZY"
*/