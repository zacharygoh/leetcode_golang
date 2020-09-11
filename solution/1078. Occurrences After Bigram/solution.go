func findOcurrences(text string, first string, second string) (result []string) {
    texts := strings.Split(text, " ")
    for i := 0 ; i < len(texts)-2; i++ {
        if texts[i] == first && texts[i+1] == second {
            result = append(result, texts[i+2])
        }
    }
    return
}

/* 
    Example 1:

    Input: text = "alice is a good girl she is a good student", first = "a", second = "good"
    Output: ["girl","student"]

    Example 2:

    Input: text = "we will we will rock you", first = "we", second = "will"
    Output: ["we","rock"]
*/