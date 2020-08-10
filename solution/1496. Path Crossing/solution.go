func isPathCrossing(path string) bool {
    // N ^, E >, S v, W <
    path_map := make(map[string]bool)
    coordinates := []string{"0", "0"}
    coordinate := strings.Join(coordinates, ",")
    path_map[coordinate] = true
    for _, value := range strings.Split(path, "") {
        switch value {
            case "N", "S":
                coor, _ := strconv.Atoi(coordinates[0])
                if value == "N" {
                    coor+=1
                }else{
                    coor-=1
                }
                coordinates[0] = strconv.Itoa(coor)
                coordinate = strings.Join(coordinates, ",")
                if path_map[coordinate] {
                    return true
                }
                path_map[coordinate] = true
                break
            case "W", "E":
                coor, _ := strconv.Atoi(coordinates[1])
                if value == "W" {
                    coor-=1
                }else{
                    coor+=1
                }
                coordinates[1] = strconv.Itoa(coor)
                coordinate = strings.Join(coordinates, ",")
                if path_map[coordinate] {
                    return true
                }
                path_map[coordinate] = true
                break
        }
    }
    return false
}

/* 
Example 1:



Input: path = "NES"
Output: false 
Explanation: Notice that the path doesn't cross any point more than once.

Example 2:



Input: path = "NESWW"
Output: true
Explanation: Notice that the path visits the origin twice.
*/

