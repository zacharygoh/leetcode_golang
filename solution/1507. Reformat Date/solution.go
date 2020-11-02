func reformatDate(date string) (result string) {
    dates := strings.Split(date, " ")
    result += dates[2] + "-"
    result += getMonth(dates[1]) + "-"
    result += justDay(dates[0])
    
    return
}

func addZero(day string) (result string) {
    if len(day) == 1 {
        result = "0" + day
    }else{
        return day
    }
    return
}

func justDay(day string) string {
    if strings.Contains(day, "th") {
        day = strings.TrimSuffix(day, "th")
    }else if strings. Contains(day, "st") {
        day = strings.TrimSuffix(day, "st")
    }else if strings. Contains(day, "rd") {
        day = strings.TrimSuffix(day, "rd")
    }else{
        day = strings.TrimSuffix(day, "nd")
    }
    day = addZero(day)
    return day
}

func getMonth(month string) (result string) {
    var month_map = map[string]int{"Jan": 1, "Feb": 2, "Mar": 3, "Apr": 4, "May": 5, "Jun": 6, "Jul": 7, "Aug": 8, "Sep": 9, "Oct": 10, "Nov": 11, "Dec": 12,}
    result = addZero(strconv.Itoa(month_map[month]))
    return 
}

/* 
	Example 1:

	Input: date = "20th Oct 2052"
	Output: "2052-10-20"

	Example 2:

	Input: date = "6th Jun 1933"
	Output: "1933-06-06"

	Example 3:

	Input: date = "26th May 1960"
	Output: "1960-05-26"
*/