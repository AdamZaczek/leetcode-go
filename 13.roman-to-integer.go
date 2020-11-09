/*
 * @lc app=leetcode id=13 lang=golang
 *
 * [13] Roman to Integer
 */

// @lc code=start
var m = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000,
}

func romanToInt(s string) int {
	var r int
	for i := 0; i < len(s); i++ {
		current := m[string(s[i])]

		// is last item, or is same as prev
		if i == len(s)-1 || current >= m[string(s[i+1])] {
			r += current
		} else {
			r -= current
		}
	}
	return r
}

// @lc code=end

