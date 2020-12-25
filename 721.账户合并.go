/*
 * @lc app=leetcode.cn id=721 lang=golang
 *
 * [721] 账户合并
 */
package main

import "sort"

// @lc code=start
func accountsMerge(accounts [][]string) [][]string {
	// Time: O(n)
	// Space: O(n)
	//
	// The emails map goes from an email address to a pair of ints:
	// [0]: an "AUTO_INCREMENT": 1, 2, 3...
	// [1]: index to one of the accounts that it came from (to grab the name)
	//
	// Just imagine it's an `emails` table with an id and a FK to `accounts`!
	// We'll do the Union-Find with its id rather than its email.
	emails := map[string][]int{}
	for i := range accounts {
		for j := 1; j < len(accounts[i]); j++ {
			if _, ok := emails[accounts[i][j]]; !ok {
				emails[accounts[i][j]] = []int{len(emails), i}
			}
		}
	}

	// Time: O(n)
	// Space: O(n)
	//
	// If two emails appear in the same account, we know they belong to the
	// same group so we can `union` them.
	uf := newUnionFind(len(emails))
	for i := range accounts {
		if len(accounts[i]) < 3 {
			continue
		}
		for j := 2; j < len(accounts[i]); j++ {
			uf.union(emails[accounts[i][j-1]][0], emails[accounts[i][j]][0])
		}
	}

	// Time: O(n)
	// Space: O(n)
	//
	// Now we can go through the emails and `find` which group they belong
	// to. We use the `groups` map to go from `group number` to merged's `index`.
	//
	// In here we finally use emails[email][1] to give a name to the merged group.
	groups := map[int]int{}
	merged := [][]string{}
	for email, i := range emails {
		group := uf.find(i[0])
		if _, ok := groups[group]; !ok {
			merged = append(merged, []string{accounts[i[1]][0], email})
			groups[group] = len(merged) - 1
		} else {
			merged[groups[group]] = append(merged[groups[group]], email)
		}
	}

	// Time: O(nlogn)
	// Space: O(1)
	//
	// Emails within each merged group must be sorted.
	for i := range merged {
		sort.Strings(merged[i][1:])
	}

	return merged
}

// Union-Find algorithm
type uf struct{ uf []int }

func (u uf) union(a, b int) {
	i, j := u.find(a), u.find(b)
	if i != j {
		u.uf[i] = j
	}
}

func (u uf) find(i int) int {
	if u.uf[i] == 0 {
		return i
	}
	return u.find(u.uf[i])
}

func newUnionFind(size int) uf {
	return uf{make([]int, size+1)}
}
// @lc code=end
