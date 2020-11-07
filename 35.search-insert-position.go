/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	for i, n := range nums {
		if n >= target {
			return i
		}
	}
	return len(nums)
}

// @lc code=end

