/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	var c int
	for i := 1; i < len(nums); i++ {
		if nums[c] != nums[i] {
			c++
			nums[c] = nums[i]
		}
	}
	return c + 1
}

// @lc code=end

