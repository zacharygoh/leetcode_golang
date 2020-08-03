func arrayPairSum(nums []int) (result int) {
	sort.Ints(nums)
	for i := 0; i < len(nums); i+2 {
		result+=nums[i]
	}
    return
}


/* 
Example 1:
Input: [1,4,3,2]

Output: 4

Explanation: n is 2, and the maximum sum of pairs is 4 = min(1, 2) + min(3, 4).
Note:
n is a positive integer, which is in the range of [1, 10000].
All the integers in the array will be in the range of [-10000, 10000].
*/