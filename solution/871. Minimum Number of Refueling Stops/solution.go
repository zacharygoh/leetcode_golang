/* Min heap solution */
func minRefuelStops(target int, startFuel int, stations [][]int) int {
    total_stations := len(stations)
    pq := &IntHeap{}
    ans, index := 0, 0
    curr := startFuel
    for curr < target {
        for index < total_stations && stations[index][0] <= curr {
            Push(pq, -stations[index][1])
            index+=1
        }
        if pq.Len() == 0 {
            return -1
        }
        curr += -Pop(pq).(int)
        ans += 1
    }
    return ans
}

type Interface interface {
	sort.Interface
	Push(x interface{}) // add x as element Len()
	Pop() interface{}   // remove and return element Len() - 1.
}
func Init(h Interface) {
	// heapify
	n := h.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(h, i, n)
	}
}
func Push(h Interface, x interface{}) {
	h.Push(x)
	// Min-Heapify
	up(h, h.Len()-1)
}
func up(h Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}
func down(h Interface, i, n int) {
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && !h.Less(j1, j2) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
}
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}
func Remove(h Interface, i int) interface{} {
	n := h.Len() - 1
	if n != i {
		h.Swap(i, n)
		down(h, i, n)
		up(h, i)
	}
	return h.Pop()
}
type IntHeap []int

func (s IntHeap) Len() int           { return len(s) }
func (s IntHeap) Less(i, j int) bool { return s[i] < s[j] }
func (s IntHeap) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func (s *IntHeap) Push(val interface{}) {
	*s = append(*s, val.(int))
}
func (s *IntHeap) Pop() interface{} {
	osl := *s
	n := len(osl)
	ns := osl[n-1]
	*s = osl[0 : n-1]
	return ns
}


/* Knapsack solution */
func minRefuelStops(target int, startFuel int, stations [][]int) int {
    total_stations := len(stations)
    dp := make([][]int, total_stations+1)
    for i, _ := range dp {
        dp[i] = make([]int, total_stations+1)
    }
    fmt.Println(dp)
    for i := 0; i < total_stations+1; i++ {
        dp[i][0] = startFuel
    }
    fmt.Println(dp)
    for i := 1; i < total_stations+1; i++ {
        for j := 1; j < i+1; j++ {
            dp[i][j] = dp[i-1][j]
            if dp[i-1][j-1] >= stations[i-1][0] {
                dp[i][j] = Max(dp[i][j], dp[i-1][j-1]+stations[i-1][1])
            }
        }
    }
    
    fmt.Println(dp)
    for i := 0; i < total_stations+1; i++ {
        if dp[total_stations][i] >= target {
            return i
        }
    }
    return -1
}

func Max(x, y int) int {
    if x < y {
        return y
    }
    return x
}

/* 
	Example 1:

	Input: target = 1, startFuel = 1, stations = []
	Output: 0
	Explanation: We can reach the target without refueling.

	Example 2:

	Input: target = 100, startFuel = 1, stations = [[10,100]]
	Output: -1
	Explanation: We can't reach the target (or even the first gas station).

	Example 3:

	Input: target = 100, startFuel = 10, stations = [[10,60],[20,30],[30,30],[60,40]]
	Output: 2
	Explanation: 
	We start with 10 liters of fuel.
	We drive to position 10, expending 10 liters of fuel.  We refuel from 0 liters to 60 liters of gas.
	Then, we drive from position 10 to position 60 (expending 50 liters of fuel),
	and refuel from 10 liters to 50 liters of gas.  We then drive to and reach the target.
	We made 2 refueling stops along the way, so we return 2.
*/