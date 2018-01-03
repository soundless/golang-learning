/*
 * [13] Roman to Integer
 *
 * https://leetcode.com/problems/roman-to-integer
 *
 * Easy (45.07%)
 * Total Accepted:    
 * Total Submissions: 
 * Testcase Example:  '"DCXXI"'
 *
 * Given a roman numeral, convert it to an integer.
 * 
 * Input is guaranteed to be within the range from 1 to 3999.
 */
package gogo

func romanToInt(s string) int {
    result := 0
    if len(s) == 0 {
        return result
    }

    hash := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    last := 0;
    for i := 0; i < len(s); i++ {
        curr := hash[s[i]]
        if last < curr {
            result += curr - 2 * last
        } else {
            result += curr
        }
        last = curr
    }

    return result
}
