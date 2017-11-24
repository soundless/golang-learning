/*
 * [4] Median of Two Sorted Arrays
 *
 * https://leetcode.com/problems/median-of-two-sorted-arrays
 *
 * Hard (21.40%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '[1,3]\n[2]'
 *
 * There are two sorted arrays nums1 and nums2 of size m and n respectively.
 *
 * Find the median of the two sorted arrays. The overall run time complexity
 * should be O(log (m+n)).
 *
 * Example 1:
 *
 * nums1 = [1, 3]
 * nums2 = [2]
 *
 * The median is 2.0
 *
 *
 *
 * Example 2:
 *
 * nums1 = [1, 2]
 * nums2 = [3, 4]
 *
 * The median is (2 + 3)/2 = 2.5
 *
 *
 */

package gogo

// O(n)
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//	i, j := 0, 0
//	nums := []int{}
//	for i < len(nums1) && j < len(nums2) {
//		if nums1[i] < nums2[j] {
//			nums = append(nums, nums1[i])
//			i += 1
//		} else {
//			nums = append(nums, nums2[j])
//			j += 1
//		}
//	}
//
//	for ; i < len(nums1); i++ {
//		nums = append(nums, nums1[i])
//	}
//
//	for ; j < len(nums2); j++ {
//		nums = append(nums, nums2[j])
//	}
//
//	length := len(nums)
//	if length%2 == 0 {
//		return (float64(nums[length/2-1]) + float64(nums[length/2])) / 2.0
//	} else {
//		return float64(nums[(length+1)/2-1])
//	}
//}

// O(log(m+n))
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	total := l1 + l2
	if total%2 == 1 {
		return float64(findKthNum(nums1, nums2, total/2+1))
	} else {
		return 0.5 * float64(findKthNum(nums1, nums2, total/2)+findKthNum(nums1, nums2, total/2+1))
	}
}

func findKthNum(nums1 []int, nums2 []int, k int) int {
	l1, l2 := len(nums1), len(nums2)
	if l1 > l2 {
		return findKthNum(nums2, nums1, k)
	}

	if l1 == 0 {
		return nums2[k-1]
	}

	if k == 1 {
		if nums1[0] < nums2[0] {
			return nums1[0]
		} else {
			return nums2[0]
		}
	}

	var i, j int
	if l1 < k/2 {
		i = l1
	} else {
		i = k / 2
	}
	j = k - i

	if nums1[i-1] < nums2[j-1] {
		return findKthNum(nums1[i:], nums2, k-i)
	} else if nums1[i-1] > nums2[j-1] {
		return findKthNum(nums1, nums2[j:], k-j)
	} else {
		return nums1[i-1]
	}
}
