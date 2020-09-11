func findRestaurant(list1 []string, list2 []string) (result []string) {
    common_map := make(map[string]int)
    min := 0 
    first_time := true
    for i1, first := range list1 {
        for i2, second := range list2 {
            if first == second {
                common_map[first] = i1 + i2
                if first_time {
                    min = i1 + i2
                    first_time = false
                }
                if min > i1 + i2 {
                    min = i1 + i2
                }
            }  
        }
    }
    
    for index, value := range common_map {
        if value == min {
            result = append(result, index)
        }
    }
    return
}

/* 
	Example 1:
	Input:
	["Shogun", "Tapioca Express", "Burger King", "KFC"]
	["Piatti", "The Grill at Torrey Pines", "Hungry Hunter Steakhouse", "Shogun"]
	Output: ["Shogun"]
	Explanation: The only restaurant they both like is "Shogun".

	Example 2:
	Input:
	["Shogun", "Tapioca Express", "Burger King", "KFC"]
	["KFC", "Shogun", "Burger King"]
	Output: ["Shogun"]
	Explanation: The restaurant they both like and have the least index sum is "Shogun" with index sum 1 (0+1).
*/