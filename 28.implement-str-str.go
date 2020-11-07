/*
 * @lc app=leetcode id=28 lang=golang
 *
 * [28] Implement strStr()
 */
import "strings"

// @lc code=start
func strStr(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}

// @lc code=end

