func lengthOfLIS(nums []int) int {
    var piles [][]int
    for _, n := range nums {
        if len(piles) == 0 {
            piles = append(piles, []int{n})
        }
        
        j := sort.Search(len(piles), func(k int) bool {
            return piles[k][len(piles[k])-1] >= n
        })
        
        if j < len(piles) {
            piles[j] = append(piles[j], n)
        } else {
            piles = append(piles, []int{n})
        }
    }
    return len(piles)
}

/* 
	Example 1:
	Input: nums = [10,9,2,5,3,7,101,18]
	Output: 4
	Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
	
	Example 2:
	Input: nums = [0,1,0,3,2,3]
	Output: 4

	Example 3:
	Input: nums = [7,7,7,7,7,7,7]
	Output: 1
*/