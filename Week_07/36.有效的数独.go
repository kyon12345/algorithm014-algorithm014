/*
 * @lc app=leetcode.cn id=36 lang=golang
 *
 * [36] 有效的数独
 */
package main


// @lc code=start
// HashSet custom set
type HashSet struct {
	set map[string]bool
}

func New() *HashSet {
	return &HashSet{make(map[string]bool)}
}

func (this *HashSet) Add(s string) bool {
	_,found := this.set[s]

	this.set[s] = true
	return !found
}

func isValidSudoku(board [][]byte) bool {
	hashS := New()
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] != '.' {
				b := string('(' + board[i][j] + ')')
				if !hashS.Add(string(i) + b) || !hashS.Add(b + string(j)) || !hashS.Add(string(i/3) + b + string(j/3)) {
					return false
				}		  				
			}
		}
	}
	return true
}
// @lc 	code=end

