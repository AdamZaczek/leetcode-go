import "strings"

/*
 * @lc app=leetcode id=14 lang=golang
 *
 * [14] Longest Common Prefix
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	switch len(strs) {
	case 0:
		return ""
	case 1:
		return strs[0]
	default:
		var prefix string
		var n int
		first := strs[0]
		remaining := strs[1:]
		for true {
			if n > len(first) {
				return prefix
			}
			prefix = first[:n]
			for _, s := range remaining {
				if !strings.HasPrefix(s, prefix) {
					return prefix[:len(prefix)-1]
				}
			}
			n++
		}
		return prefix
	}

}

// @lc code=end

