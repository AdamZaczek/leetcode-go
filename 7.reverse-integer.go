import (
	"math"
)

/*
 * @lc app=leetcode id=7 lang=golang
 *
 * [7] Reverse Integer
 */

// @lc code=start
func reverse(x int) int {
	if x == 0 || x >= math.MaxInt32 || x <= math.MinInt32 {
		return 0
	}
	r := 0
	t := x

	// validate negative
	isNegative := x < 0
	if isNegative {
		t *= -1
	}
	for t > 0 {
		r *= 10
		n := t % 10
		r += n
		t /= 10
	}
	if r >= math.MaxInt32 {
		return 0
	}

	if isNegative {
		return -r
	}

	return r

}

// @lc code=end

