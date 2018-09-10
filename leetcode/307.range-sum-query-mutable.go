/*
 * [307] Range Sum Query - Mutable
 *
 * https://leetcode.com/problems/range-sum-query-mutable
 *
 * Medium (19.91%)
 * Total Accepted:    
 * Total Submissions: 
 * Testcase Example:  '["NumArray","sumRange","update","sumRange"]\n[[[1,3,5]],[0,2],[1,2],[0,2]]'
 *
 * Given an integer array nums, find the sum of the elements between indices i
 * and j (i ≤ j), inclusive.
 * 

package gogo

 * The update(i, val) function modifies nums by updating the element at index i
 * to val.
 * 
 * Example:
 * 
 * Given nums = [1, 3, 5]
 * 
 * sumRange(0, 2) -> 9
 * update(1, 2)
 * sumRange(0, 2) -> 8
 * 
 * 
 * 
 * Note:
 * 

package gogo

 * The array is only modifiable by the update function.

package gogo

 * You may assume the number of calls to update and sumRange function is
 * distributed evenly.
 * 
 * 
 */
type NumArray struct {
    
}



package gogo

func Constructor(nums []int) NumArray {
    
}



package gogo

func (this *NumArray) Update(i int, val int)  {
    
}



package gogo

func (this *NumArray) SumRange(i int, j int) int {
    
}


/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * obj.Update(i,val);
 * param_2 := obj.SumRange(i,j);
 */
