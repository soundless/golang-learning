/*
 * [12] Integer to Roman
 *
 * https://leetcode.com/problems/integer-to-roman
 *
 * Medium (44.09%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '1'
 *
 * Given an integer, convert it to a roman numeral.
 *
 *
 * Input is guaranteed to be within the range from 1 to 3999.
 */

package gogo

import "sort"

func intToRoman(num int) string {
	hash := map[int]string{
		1000: "M",
		900:  "CM",
		500:  "D",
		400:  "CD",
		100:  "C",
		90:   "XC",
		50:   "L",
		40:   "XL",
		10:   "X",
		9:    "IX",
		5:    "V",
		4:    "IV",
		1:    "I",
	}
	keys := make([]int, 0, len(hash))
	for k := range hash {
		keys = append(keys, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(keys)))
	result := ""
	for num > 0 {
		for _, key := range keys {
			for num/key > 0 {
				num -= key
				result += hash[key]
			}
		}
	}
	return result
}
