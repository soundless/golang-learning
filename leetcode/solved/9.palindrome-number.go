/*
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number
 *
 * Easy (34.87%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '-2147483648'
 *
 * Determine whether an integer is a palindrome. Do this without extra space.
 *
 * click to show spoilers.
 *
 * Some hints:
 *
 * Could negative integers be palindromes? (ie, -1)
 *
 * If you are thinking of converting the integer to string, note the
 * restriction of using extra space.
 *
 * You could also try reversing an integer. However, if you have solved the
 * problem "Reverse Integer", you know that the reversed integer might
 * overflow. How would you handle such case?
 *
 * There is a more generic way of solving this problem.
 *
 *
 */

package gogo

//import "strconv"

//func isPalindrome(x int) bool {
//    if x < 0 {
//        return false
//    }
//
//    s := strconv.Itoa(x)
//    for i,j := 0, len(s) - 1; i < len(s); i, j = i+1, j-1 {
//        if s[i] != s[j] {
//            return false
//        }
//    }
//    return true
//}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	copied, reversed := x, 0
	for copied > 0 {
		reversed *= 10
		reversed += copied % 10
		copied /= 10
	}
	if x == reversed {
		return true
	}
	return false
}
