/*
 * [8] String to Integer (atoi)
 *
 * https://leetcode.com/problems/string-to-integer-atoi
 *
 * Medium (13.94%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '""'
 *
 * Implement atoi to convert a string to an integer.
 *
 * Hint: Carefully consider all possible input cases. If you want a challenge,
 * please do not see below and ask yourself what are the possible input
 * cases.
 *
 *
 * Notes:
 * It is intended for this problem to be specified vaguely (ie, no given input
 * specs). You are responsible to gather all the input requirements up
 * front.
 *
 *
 * Update (2015-02-10):
 * The signature of the C++ function had been updated. If you still see your
 * function signature accepts a const char * argument, please click the reload
 * button  to reset your code definition.
 *
 *
 * spoilers alert... click to show requirements for atoi.
 *
 * Requirements for atoi:
 *
 * The function first discards as many whitespace characters as necessary until
 * the first non-whitespace character is found. Then, starting from this
 * character, takes an optional initial plus or minus sign followed by as many
 * numerical digits as possible, and interprets them as a numerical value.
 *
 * The string can contain additional characters after those that form the
 * integral number, which are ignored and have no effect on the behavior of
 * this function.
 *
 * If the first sequence of non-whitespace characters in str is not a valid
 * integral number, or if no such sequence exists because either str is empty
 * or it contains only whitespace characters, no conversion is performed.
 *
 * If no valid conversion could be performed, a zero value is returned. If the
 * correct value is out of the range of representable values, INT_MAX
 * (2147483647) or INT_MIN (-2147483648) is returned.
 *
 *
 */

package gogo

import (
	"math"
	"strings"
)

func myAtoi(str string) int {
	s := strings.TrimSpace(str)
	if len(s) == 0 {
		return 0
	}

	sign, x := getSign(s)
	x = trimInvalid(x)

	return convertToInt(sign, x)
}

func getSign(s string) (int, string) {
	sign := 1
	switch s[0] {
	case '-':
		{
			sign = -1
			s = s[1:]
		}
	case '+':
		{
			s = s[1:]
		}
	default:
		{
		}
	}
	return sign, s
}

func trimInvalid(s string) string {
	for i := range s {
		if s[i] < '0' || s[i] > '9' {
			return s[:i]
		}
	}
	return s
}

func convertToInt(sign int, s string) int {
	base := 1 * sign
	result := 0
	overflow := false

	for i := len(s) - 1; i >= 0; i-- {
		result, overflow = isOverflow(result + int(s[i]-48)*base)
		if overflow {
			return result
		}
		base *= 10
	}

	return result
}

func isOverflow(i int) (int, bool) {
	if i > math.MaxInt32 {
		return math.MaxInt32, true
	}
	if i < math.MinInt32 {
		return math.MinInt32, true
	}
	return i, false
}
