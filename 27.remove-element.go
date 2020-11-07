/*
 * @lc app=leetcode id=27 lang=golang
 *
 * [27] Remove Element
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	n := 0
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		if num != val {
			nums[n] = num
			n++
		}
	}
	return n
}

// @lc code=end

