/*
 * [6] ZigZag Conversion
 *
 * https://leetcode.com/problems/zigzag-conversion
 *
 * Medium (26.57%)
 * Total Accepted:
 * Total Submissions:
 * Testcase Example:  '""\n1'
 *
 *
 * The string "PAYPALISHIRING" is written in a zigzag pattern on a given number
 * of rows like this: (you may want to display this pattern in a fixed font for
 * better legibility)
 *
 * P   A   H   N
 * A P L S I I G
 * Y   I   R
 *
 *
 * And then read line by line: "PAHNAPLSIIGYIR"
 *
 *
 * Write the code that will take a string and make this conversion given a
 * number of rows:
 *
 * string convert(string text, int nRows);
 *
 * convert("PAYPALISHIRING", 3) should return "PAHNAPLSIIGYIR".
 *
 */
package gogo

import "bytes"

func convert(s string, numRows int) string {
    if numRows == 1 || len(s) <= numRows {
        return s
    }

    result := bytes.Buffer{}
    // pace
    p := numRows * 2 - 2

    for i := 0; i < len(s); i += p {
        result.WriteByte(s[i])
    }

    for r := 1; r <= numRows - 2; r++ {
        result.WriteByte(s[r])

        for k := p; k - r < len(s); k += p {
            result.WriteByte(s[k-r])
            if k + r < len(s) {
                result.WriteByte(s[k + r])
            }
        }
    }

    for i := numRows - 1; i < len(s); i += p {
        result.WriteByte(s[i])
    }

    return result.String()
}
