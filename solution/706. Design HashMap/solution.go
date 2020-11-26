const _size = 100000

type ListNode struct {
    Key int
    Val int
    Next *ListNode
}

type MyHashMap struct {
    table [_size]*ListNode
}


/** Initialize your data structure here. */
func Constructor() MyHashMap {
    return MyHashMap{}
}


/** value will always be non-negative. */
func (m *MyHashMap) Put(key int, value int)  {
    var prev *ListNode
    hash := key % _size
    curr := m.table[hash]
    node := ListNode{
        Key: key,
        Val: value,
    }
    
    if curr == nil {
        m.table[hash] = &node
        return
    }
    for curr != nil {
        if curr.Key == key {
            curr.Val = value
            return
        }
        prev = curr
        curr = curr.Next
    }
     prev.Next = &node
}


/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (m *MyHashMap) Get(key int) int {
    hash := key % _size
    curr := m.table[hash]
    for curr != nil {
        if curr.Key == key {
            return curr.Val
        }
        curr = curr.Next
    }
    return -1
}


/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (m *MyHashMap) Remove(key int)  {
    var prev *ListNode
    
    hash := key % _size
    curr := m.table[hash]
    if curr == nil {
        return
    }
    
    if curr.Key == key {
        m.table[hash] = curr.Next
        return
    }
    
    for curr != nil {
        if curr.Key == key {
            if prev == nil {
                m.table[hash] = nil
                return
            }
            prev.Next = curr.Next
            return
        }
        prev = curr
        curr = curr.Next
    }
}


/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */

 /* 
	Example:

	MyHashMap hashMap = new MyHashMap();
	hashMap.put(1, 1);          
	hashMap.put(2, 2);         
	hashMap.get(1);            // returns 1
	hashMap.get(3);            // returns -1 (not found)
	hashMap.put(2, 1);          // update the existing value
	hashMap.get(2);            // returns 1 
	hashMap.remove(2);          // remove the mapping for 2
	hashMap.get(2);            // returns -1 (not found) 
 */