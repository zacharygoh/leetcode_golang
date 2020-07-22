func distributeCandies(candies int, num_people int) []int {
    people_candies := make([]int, num_people)
    candy := 1
    for candies > 0 {
        for i := 0; i < num_people; i++ {
            if candy > candies {
                if candies < 0 {
                    break
                }
                people_candies[i] += candies
            }else{
                people_candies[i] += candy
            }
            candies -= candy
            candy++
        }
    }
    return people_candies
}

/* 
Example 1:

Input: candies = 7, num_people = 4
Output: [1,2,3,1]
Explanation:
On the first turn, ans[0] += 1, and the array is [1,0,0,0].
On the second turn, ans[1] += 2, and the array is [1,2,0,0].
On the third turn, ans[2] += 3, and the array is [1,2,3,0].
On the fourth turn, ans[3] += 1 (because there is only one candy left), and the final array is [1,2,3,1].

Example 2:

Input: candies = 10, num_people = 3
Output: [5,2,3]
Explanation: 
On the first turn, ans[0] += 1, and the array is [1,0,0].
On the second turn, ans[1] += 2, and the array is [1,2,0].
On the third turn, ans[2] += 3, and the array is [1,2,3].
On the fourth turn, ans[0] += 4, and the final array is [5,2,3].
*/