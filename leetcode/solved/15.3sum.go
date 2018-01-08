/*
 * [15] 3Sum
 *
 * https://leetcode.com/problems/3sum
 *
 * Medium (21.52%)
 * Total Accepted:    
 * Total Submissions: 
 * Testcase Example:  '[-1,0,1,2,-1,-4]'
 *
 * Given an array S of n integers, are there elements a, b, c in S such that a
 * + b + c = 0? Find all unique triplets in the array which gives the sum of
 * zero.
 * 
 * Note: The solution set must not contain duplicate triplets.
 * 
 * 
 * For example, given array S = [-1, 0, 1, 2, -1, -4],
 * 
 * A solution set is:
 * [
 * ⁠ [-1, 0, 1],
 * ⁠ [-1, -1, 2]
 * ]
 * 
 */
package gogo

import (
    "reflect"
    "sort"
)

func threeSum(nums []int) [][]int {
    var result [][]int
    if len(nums) == 0 {
        return result
    }
    for i := 0; i < len(nums); i++ {
        target := 0 - nums[i]
        hash := make(map[int]int)
        for j := i+1; j < len(nums); j++ {
            temp := target - nums[j]
            if _, ok := hash[temp]; ok {
                list := make([]int, 0)
                list = append(list, nums[i], temp, nums[j])
                if containSet(result, list) != true {
                    result = append(result, list)
                }
            } else {
                hash[nums[j]] = j
            }
        }
    }
    return result
}

func containSet(slices [][]int, item []int) bool {
    for _, slice := range(slices) {
        sort.Ints(slice)
        sort.Ints(item)
        if reflect.DeepEqual(slice, item) == true {
            return true
        }
    }
    return false
}
