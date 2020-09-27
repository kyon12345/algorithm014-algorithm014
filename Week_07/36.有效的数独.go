/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */
package main


// @lc code=start
// HashSet custom set
type HashSet struct {
	list []string
}

// New HashSet
func New() HashSet {
	set := HashSet{list: make([]string, 0)}
	return set
}


// Return true if val not in the set
func (set *HashSet) add(val string) bool {
	var flag bool = true
	if set.Contains(val) {
		flag = false
	}
	set.list = append(set.list, val)
	return flag
}

func (set *HashSet) String() string {
	return string(len(set.list))
}



// Contains check if val is in list
func (set *HashSet) Contains(val string) bool {
	for _, value := range set.list {
		if value == val {
			return true
		}
	}
	return false
}


func isValidSudoku(board [][]byte) bool {
	set := New()

	for i:=0; i<9; i++ {
		for j:=0; j<9; j++ {
			val := string(board[i][j])
			if val != "." {
				if !set.add(val + "row" + string(i)) || !set.add(val + "col" + string(j)) || !set.add(val + "box" + string(i/3) + string(j/3)) {
					return false
				}

			}
		}
	}
	
	return true
}
// @lc 	code=end

