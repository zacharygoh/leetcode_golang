func numRescueBoats(people []int, limit int) (result int) {
    sort.Ints(people)
    i, j := 0, len(people)-1
    for i <= j {
        result++
        if people[i] + people[j] <= limit {
            j--
            i++
        }else{
            j--
        }
    }
    return
}

/* 
	Example 1:

	Input: people = [1,2], limit = 3
	Output: 1
	Explanation: 1 boat (1, 2)

	Example 2:

	Input: people = [3,2,2,1], limit = 3
	Output: 3
	Explanation: 3 boats (1, 2), (2) and (3)

	Example 3:

	Input: people = [3,5,3,4], limit = 5
	Output: 4
	Explanation: 4 boats (3), (3), (4), (5)
*/