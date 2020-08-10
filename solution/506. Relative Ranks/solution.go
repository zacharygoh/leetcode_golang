func findRelativeRanks(nums []int) (result []string) {
    score_map := make(map[int]string)
    var ss []int
    for _, value := range nums {
        ss = append(ss, value)
    }
    sort.Ints(ss)
    max_score := len(ss)-1
    ii := 1
    for i := len(ss)-1; i > -1; i-- {
        if i == max_score {
            score_map[ss[i]] = "Gold Medal"
        }else if i == max_score - 1 {
            score_map[ss[i]] = "Silver Medal"
        }else if i == max_score - 2 {
            score_map[ss[i]] = "Bronze Medal"
        }else {
            score_map[ss[i]] = strconv.Itoa(ii)
        }
        ii++
    } 
    
    for i := 0; i < len(nums) ; i++ {
        result = append(result, score_map[nums[i]])
    }
    return
}

/* 
Example 1:
Input: [5, 4, 3, 2, 1]
Output: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
Explanation: The first three athletes got the top three highest scores, so they got "Gold Medal", "Silver Medal" and "Bronze Medal". 
For the left two athletes, you just need to output their relative ranks according to their scores.
*/