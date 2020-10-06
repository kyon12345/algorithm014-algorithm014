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

func New() HashSet {
	set := HashSet{list: []string{}}
	return set
}

func (set *HashSet) Add (s string) bool {
	if set.Contain(s) {
		return false
	} else {
		set.list = append(set.list, s)
	}

	return true
}

func (set *HashSet) Contain(s string) bool {
	for _, v := range set.list {
		if v == s {
			return true
		}
	}
	return false
}

func isValidSudoku(board [][]byte) bool {
	set := New()

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val := string(board[i][j])

			if val != "." {
			  if !set.Add(val + "row" + string(i)) || !set.Add(val + "col" + string(j)) || !set.Add(val + "box" + string(i/3) + string(j / 3)) {
					return false
				} 
			}
		}
	}
	
	
	return true
}
// @lc 	code=end

