func groupThePeople(groupSizes []int) (result [][]int) {
    // tracker for the number and temporary insert into array of the number's index
    tracker := make(map[int][]int)
    for index, value := range groupSizes {
        // get the number index, and insert into the array
        tracker[value] = append(tracker[value], index)
        // if the tracker's number of each index maxed out
        if len(tracker[value]) == value {
            // insert into the two dimensional array
            result = append(result, tracker[value])
            // reset the value of the array
            tracker[value] = []int{}
        }
    }
    return
}

/* 
	Example 1:

	Input: groupSizes = [3,3,3,3,3,1,3]
	Output: [[5],[0,1,2],[3,4,6]]
	Explanation: 
	The first group is [5]. The size is 1, and groupSizes[5] = 1.
	The second group is [0,1,2]. The size is 3, and groupSizes[0] = groupSizes[1] = groupSizes[2] = 3.
	The third group is [3,4,6]. The size is 3, and groupSizes[3] = groupSizes[4] = groupSizes[6] = 3.
	Other possible solutions are [[2,1,6],[5],[0,4,3]] and [[5],[0,6,2],[4,3,1]].
	
	Example 2:

	Input: groupSizes = [2,1,3,3,3,2]
	Output: [[1],[0,5],[2,3,4]]
*/