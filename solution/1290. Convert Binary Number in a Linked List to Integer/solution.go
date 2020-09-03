/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func getDecimalValue(head *ListNode) (result int) {
    test := ""
    counter := -1
    for head != nil {
        test += strconv.Itoa(head.Val)
        head = head.Next
        counter++
    }
    for _, value := range strings.Split(test, "") {
        n, _ := strconv.Atoi(value)
        result += n * int(math.Pow(2.0, float64(counter)))
        counter--
    }
    return
}

/* 
Example 1:


Input: head = [1,0,1]
Output: 5
Explanation: (101) in base 2 = (5) in base 10

Example 2:

Input: head = [0]
Output: 0

Example 3:

Input: head = [1]
Output: 1

Example 4:

Input: head = [1,0,0,1,0,0,1,1,1,0,0,0,0,0,0]
Output: 18880
*/