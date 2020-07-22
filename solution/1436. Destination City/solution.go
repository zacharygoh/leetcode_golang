var destination = make(map[string]string)

func destCity(paths [][]string) string {
    for _, value := range paths {
        destination[value[0]] = value[1]
    }
    return findDestination(paths[0][1])
}

func findDestination(to string) string {
    if _, found := destination[to]; found {
        return findDestination(destination[to])
    }
    return to
}

/* 
Example 1:

Input: paths = [["London","New York"],["New York","Lima"],["Lima","Sao Paulo"]]
Output: "Sao Paulo" 
Explanation: Starting at "London" city you will reach "Sao Paulo" city which is the destination city. Your trip consist of: "London" -> "New York" -> "Lima" -> "Sao Paulo".

Example 2:

Input: paths = [["B","C"],["D","B"],["C","A"]]
Output: "A"
Explanation: All possible trips are: 
"D" -> "B" -> "C" -> "A". 
"B" -> "C" -> "A". 
"C" -> "A". 
"A". 
Clearly the destination city is "A".

Example 3:

Input: paths = [["A","Z"]]
Output: "Z"
*/