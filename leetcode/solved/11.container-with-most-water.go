/*
 * [11] Container With Most Water
 *
 * https://leetcode.com/problems/container-with-most-water
 *
 * Medium (36.36%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '[1,1]'
 *
 * Given n non-negative integers a1, a2, ..., an, where each represents a point
 * at coordinate (i, ai). n vertical lines are drawn such that the two
 * endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together
 * with x-axis forms a container, such that the container contains the most
 * water.
 *
 * Note: You may not slant the container and n is at least 2.
 *
 */

package gogo

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	volume := 0
	left := 0
	right := len(height) - 1
	for left < right {
		temp := MinInt(height[left], height[right]) * (right - left)
		if temp > volume {
			volume = temp
		}
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return volume
}
