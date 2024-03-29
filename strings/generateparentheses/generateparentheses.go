// generateparenthesesgo
// description: Generate Parentheses
// details:
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
// author(s) [red_byte](https://github.com/i-redbyte)
// see generateparentheses_test.go

package generateparentheses

import (
	"fmt"
	"strings"
)

func GenerateParenthesis(n int) []string {
	i := 0
	result := make([]string, 0)
	maxLen := 2 * n
	var recursiveComputation func(s []string, left int, right int)
	recursiveComputation = func(s []string, left int, right int) {
		i++
		fmt.Println(i, s, left, right)
		if len(s) == maxLen {
			result = append(result, strings.Join(s, ""))
			return
		}
		if left < n {
			s = append(s, "(")
			recursiveComputation(s, left+1, right)
			fmt.Println(i, "left before", s)
			s = s[:len(s)-1] // 去掉最后一个元素
			fmt.Println(i, "left after", s)
		}
		if right < left {
			s = append(s, ")")
			recursiveComputation(s, left, right+1)
			fmt.Println(i, "right before", s)
			_ = s[:len(s)-1]
			fmt.Println(i, "right after", s)
		}
	}
	recursiveComputation(make([]string, 0), 0, 0)
	return result
}
