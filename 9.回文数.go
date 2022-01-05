/*
 * @lc app=leetcode.cn id=9 lang=golang
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	cur := 0
	num := x
	for num > 0 {
		cur = cur * 10 + num % 10
		num = num / 10
	}
	return cur == x
}
// @lcode=end

