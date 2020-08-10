func lemonadeChange(bills []int) bool {
    changes := map[int]int{
        5: 0,
        10: 0,
    }
    
    for _, value := range bills {
        if value == 5 {
            changes[value]++
            continue
        }else if value == 10 {
            if changes[5] > 0 {
                changes[value]++
                changes[5]--
                continue
            }else{
                return false
            }
        }else {
            if changes[10] > 0 && changes[5] > 0 {
                changes[10]--
                changes[5]--
                continue
            }else if changes[5] > 2 {
                changes[5] -= 3
                continue
            }else{
                return false
            }
        }
    }
    return true
}

/* 
Example 1:

Input: [5,5,5,10,20]
Output: true
Explanation: 
From the first 3 customers, we collect three $5 bills in order.
From the fourth customer, we collect a $10 bill and give back a $5.
From the fifth customer, we give a $10 bill and a $5 bill.
Since all customers got correct change, we output true.

Example 2:

Input: [5,5,10]
Output: true

Example 3:

Input: [10,10]
Output: false

Example 4:

Input: [5,5,10,10,20]
Output: false
Explanation: 
From the first two customers in order, we collect two $5 bills.
For the next two customers in order, we collect a $10 bill and give back a $5 bill.
For the last customer, we can't give change of $15 back because we only have two $10 bills.
Since not every customer received correct change, the answer is false.
*/

