/*
 * @lc app=leetcode.cn id=1 lang=golang
 *
 * [1] 两数之和
 */

// @lc code=start
func twoSum(nums []int, target int) []int {
	bucket := make(map[int]int)
	for first, value := range nums {
		if second, ok := bucket[target - value]; ok {
			return []int{first, second}
		} else {
			bucket[value] = first
		}
	}
	return nil
}
// @lc code=end

