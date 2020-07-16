func defangIPaddr(address string) string {
    return strings.Replace(address, ".", "[.]", -1)
}

/* 
Example 1:

Input: address = "1.1.1.1"
Output: "1[.]1[.]1[.]1"

Example 2:

Input: address = "255.100.50.0"
Output: "255[.]100[.]50[.]0" 
*/