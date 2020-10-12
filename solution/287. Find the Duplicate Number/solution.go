/* Floyd Cycle */

func findDuplicate(nums []int) int {
    slow, fast := nums[0], nums[nums[0]]
    for slow != fast {
        slow = nums[slow]
        fast = nums[nums[fast]]
    }
    fast = 0
    for slow != fast {
        slow = nums[slow]
        fast = nums[fast]
    }
    return slow
}

/* 
	Follow-ups:

	How can we prove that at least one duplicate number must exist in nums? 
	Can you solve the problem without modifying the array nums?
	Can you solve the problem using only constant, O(1) extra space?
	Can you solve the problem with runtime complexity less than O(n2)?
	

	Example 1:

	Input: nums = [1,3,4,2,2]
	Output: 2

	Example 2:

	Input: nums = [3,1,3,4,2]
	Output: 3

	Example 3:

	Input: nums = [1,1]
	Output: 1

	Example 4:

	Input: nums = [1,1,2]
	Output: 1
*/