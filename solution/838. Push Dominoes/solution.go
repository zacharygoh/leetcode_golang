func pushDominoes(dominoes string) string {
    result := make([]string, 0)
    n := len(dominoes)
    forces := make([]int, n)
    force := 0
    for i, value := range strings.Split(dominoes, "") {
        if value == "R" {
            force = n
        }else if value == "L" {
            force = 0
        }else{
            force = max(force-1, 0)
        }
        forces[i] += force
    }
    force = 0
    dominoess := strings.Split(dominoes, "")
    for i := n-1; i > -1; i-- {
        if dominoess[i] == "L" {
            force = n
        }else if dominoess[i] == "R" {
            force = 0
        }else{
            force = max(force-1, 0)
        }
        forces[i] -= force
    }
    
    for _, value := range forces {
        if value > 0 {
            result = append(result, "R")
        }else if value < 0 {
            result = append(result, "L")
        }else{
            result = append(result, ".")
        }
    }
    
    return strings.Join(result, "")
} 

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

/* 
	Example 1:

	Input: ".L.R...LR..L.."
	Output: "LL.RR.LLRRLL.."
	
	Example 2:

	Input: "RR.L"
	Output: "RR.L"
	Explanation: The first domino expends no additional force on the second domino.
*/