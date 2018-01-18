/*
 * [18] 4Sum
 *
 * https://leetcode.com/problems/4sum
 *
 * Medium (26.40%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '[1,0,-1,0,-2,2]\n0'
 *
 * Given an array S of n integers, are there elements a, b, c, and d in S such
 * that a + b + c + d = target? Find all unique quadruplets in the array which
 * gives the sum of target.
 *
 * Note: The solution set must not contain duplicate quadruplets.
 *
 *
 *
 * For example, given array S = [1, 0, -1, 0, -2, 2], and target = 0.
 *
 * A solution set is:
 * [
 * ⁠ [-1,  0, 0, 1],
 * ⁠ [-2, -1, 1, 2],
 * ⁠ [-2,  0, 0, 2]
 * ]
 *
 */
package gogo

import "sort"

func fourSum(nums []int, target int) [][]int {
	result := make([][]int, 0)
	if len(nums) == 0 {
		return result
	}

	sort.Ints(nums)

	for a := 0; a < len(nums); a++ {
		for b := a + 1; b < len(nums); b++ {
			hash := make(map[int]int)
			for c := b + 1; c < len(nums); c++ {
				d_num := target - nums[a] - nums[b] - nums[c]
				if _, ok := hash[d_num]; ok {
					result = append(result, []int{nums[a], nums[b], d_num, nums[c]})
				} else {
					hash[nums[c]] = c
				}
			}
		}
	}
	return result
}
