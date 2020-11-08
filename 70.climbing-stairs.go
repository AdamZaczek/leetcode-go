/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
var cache = make(map[int]int)

func count(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	}
	if _, ok := cache[n]; ok {
		return cache[n]
	}
	val := count(n-1) + count(n-2)
	cache[n] = val
	return val
}

func climbStairs(n int) int {
	if n > 45 || n < 1 {
		panic("Number must be between 1 & 45")
	}
	return count(n)
}

// @lc code=end

