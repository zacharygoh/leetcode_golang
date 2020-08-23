func dayOfTheWeek(day int, month int, year int) string {
	//     c, y := 0, 0
	//     years := strings.Split(strconv.Itoa(year), "")
	//     cc, yy := years[0] + years[1], years[2] + years[3]
		
	//     ccc, _ := strconv.Atoi(cc)
	//     yyy, _ := strconv.Atoi(yy)
	//     c = ccc
	//     y = yyy
	// y, d, m := year, day, month
	// w := (y + y / 4 - y / 100 + y / 400 + m+d) % 7
	// switch w {
	//     case 0:
	//         return "Sunday"
	//     case 1: 
	//         return "Monday"
	//     case 2:
	//         return "Tuesday"
	//     case 3: 
	//         return "Wednesday"
	//     case 4:
	//         return "Thursday"
	//     case 5:
	//         return "Friday"
	//     case 6: 
	//         return "Saturday"
	// }
	// return ""
	m := strconv.Itoa(month)
	if len(m) == 1 {
		m = "0" + m
	}
	d := strconv.Itoa(day)
	if len(d) == 1{
		d = "0" + d
	}
	date := m + "-" + d + "-" + strconv.Itoa(year)//  "01-25-2019"
	t, err := time.Parse("01-02-2006", date)
	if err != nil {
		panic(err)
	}
	return t.Weekday().String()
}

/* 
	dow(m,d,y){y-=m<3;return(y+y/4-y/100+y/400+"-bed=pen+mad."[m]+d)%7;}
	(k + (2.6m - .2) - 2C + Y + [Y/4] + [C/4])mod7
	k is day (1 to 31)
	m is month (1 = March, .., 10 = December, 11 = Jan, 12 = Feb)
	C is century (1987 has C = 19)
	Y is year (1987 has Y = 87 except Y = 86 for Jan & Feb)
	W is weekday (0 = Sunday, .., 6 = Saturday)
*/

/* 
Example 1:

Input: day = 31, month = 8, year = 2019
Output: "Saturday"

Example 2:

Input: day = 18, month = 7, year = 1999
Output: "Sunday"

Example 3:

Input: day = 15, month = 8, year = 1993
Output: "Sunday"
*/