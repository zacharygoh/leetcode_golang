func strStr(text string, pattern string) (result int) {
    if len(pattern) > len(text) {
        return -1
    }
    if len(pattern) == 0 && len(text) == 0 {
        return 0
    }
    if len(text) > 0 && len(pattern) == 0 {
        return 0
    }
    if len(text) == 0 && len(pattern) > 0 {
        return -1
    }
	matches := match(pattern, text)
    if len(matches) > 0 {
	    return matches[0]
    }else{
        return -1
    }
}

func match(pattern, text string) (matches []int) {
	m, n, aligned_at := len(text), len(pattern), 0
    rightMostIndexes := preprocessForBadCharacterShift(pattern)
    fmt.Print(rightMostIndexes)
    texts, patterns := strings.Split(text, ""), strings.Split(pattern, "")
	for aligned_at + (n-1) < m {
		for indexInPattern := n-1; indexInPattern >= 0; indexInPattern-- {
            indexInText := aligned_at + indexInPattern
			x := texts[indexInText]
			y := patterns[indexInPattern]
            if indexInText >= m {
                break
            }
			if x != y {
				r := rightMostIndexes[x]
				if r < 0 {
					aligned_at = indexInText + 1
				}else{
					shift := indexInText - (aligned_at + r)
					if shift > 0 {
						aligned_at += shift
					}else{
						aligned_at += 1
					}
					break
				}
			}else if indexInPattern == 0 {
                matches = append(matches, aligned_at)
                aligned_at++
            }
		}
	}
	return
}

func preprocessForBadCharacterShift(pattern string) map[string]int {
    mapped := make(map[string]int)
	for index, value := range strings.Split(pattern, "") {
        mapped[value] = index
	}
	return mapped
}

/* 
Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2

Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
*/