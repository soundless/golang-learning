/*
 * [17] Letter Combinations of a Phone Number
 *
 * https://leetcode.com/problems/letter-combinations-of-a-phone-number
 *
 * Medium (33.90%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '""'
 *
 * Given a digit string, return all possible letter combinations that the
 * number could represent.
 *
 *
 *
 * A mapping of digit to letters (just like on the telephone buttons) is given
 * below.
 *
 *
 *
 * Input:Digit string "23"
 * Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
 *
 *
 *
 * Note:
 * Although the above answer is in lexicographical order, your answer could be
 * in any order you want.
 *
 */
package gogo

// import "log"

func letterCombinations(digits string) []string {
	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}

	hash := map[byte]([]rune){
		'2': {'a', 'b', 'c'},
		'3': {'d', 'e', 'f'},
		'4': {'g', 'h', 'i'},
		'5': {'j', 'k', 'l'},
		'6': {'m', 'n', 'o'},
		'7': {'p', 'q', 'r', 's'},
		'8': {'t', 'u', 'v'},
		'9': {'w', 'x', 'y', 'z'},
	}

	result = helper(result, digits, make([]rune, 0), 0, hash)
	// log.Println(result)
	return result
}

func helper(result []string, digits string, temp []rune, index int, hash map[byte]([]rune)) []string {
	if len(temp) == len(digits) {
		// log.Println(temp)
		return append(result, string(temp))
	}

	c := digits[index]
	chars := hash[c]

	for j := 0; j < len(chars); j++ {
		temp = append(temp, chars[j])
		result = helper(result, digits, temp, index+1, hash)
		temp = temp[:len(temp)-1]
	}

	return result
}
