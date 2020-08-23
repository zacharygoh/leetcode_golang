func lastStoneWeight(stones []int) int {
    if len(stones) == 0 {
        return 0
    }
    if len(stones) == 1 {
        return stones[0]
    }
    sort.Ints(stones)
    w1, w2 := stones[len(stones)-1], stones[len(stones)-2]
    newStones := stones[0: len(stones)-2]
    
    if w1 - w2 != 0 {
        newStones = append(newStones, w1 - w2)
    }
    return lastStoneWeight(newStones)
}

/* 
Example 1:

Input: [2,7,4,1,8,1]
Output: 1
Explanation: 
We combine 7 and 8 to get 1 so the array converts to [2,4,1,1,1] then,
we combine 2 and 4 to get 2 so the array converts to [2,1,1,1] then,
we combine 2 and 1 to get 1 so the array converts to [1,1,1] then,
we combine 1 and 1 to get 0 so the array converts to [1] then that's the value of last stone.
*/