/*
 * [16] 3Sum Closest
 *
 * https://leetcode.com/problems/3sum-closest
 *
 * Medium (30.88%)
 * Total Accepted:    
 * Total Submissions: 
 * Testcase Example:  '[0,0,0]\n1'
 *
 * Given an array S of n integers, find three integers in S such that the sum
 * is closest to a given number, target. Return the sum of the three integers.
 * You may assume that each input would have exactly one solution.
 * 
 * 
 * ⁠   For example, given array S = {-1 2 1 -4}, and target = 1.
 * 
 * ⁠   The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
 * 
 */
package gogo

import "math"

func threeSumClosest(nums []int, target int) int {
    result := 0
    min := math.MaxInt32
    for i := 0; i < len(nums) - 2; i++ {
        j := i + 1
        k := len(nums) - 1
        for j < k {
            sum := nums[i] + nums[j] + nums[k]
            diff := target - sum
            // math.abs() is not avilable for integer
            if diff < 0 {
                diff = -diff
            }
            if diff == 0 {
                return sum
            }
            if diff < min {
                min = diff
                result = sum
            }
            if diff < target {
                j++
            } else {
                k--
            }
        }
    }
    return result
}
