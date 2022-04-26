package balancedparentheses

import "fmt"

// IsBalanced returns true if provided input string is properly nested.
//
// Input is a sequence of brackets: '(', ')', '[', ']', '{', '}'.
//
// A sequence of brackets `s` is considered properly nested
// if any of the following conditions are true:
// 	- `s` is empty;
// 	- `s` has the form (U) or [U] or {U} where U is a properly nested string;
// 	- `s` has the form VW where V and W are properly nested strings.
//
// For example, the string "()()[()]" is properly nested but "[(()]" is not.
//
// **Note** Providing characters other then brackets would return false,
// despite brackets sequence in the string. Make sure to filter
// input before usage.

func BalanceChecker(input string) bool {
	fmt.Println(input)
	if len(input) == 0 {
		return true
	}
	if len(input)%2 != 0 {
		return false
	}
	var stack []byte
	for i := 0; i < len(input); i++ {
		fmt.Println(input[i])
		if input[i] == '(' || input[i] == '{' || input[i] == '[' {
			stack = append(stack, input[i])
		} else {
			fmt.Println(stack)
			if len(stack) > 0 {
				pair := string(stack[len(stack)-1]) + string(input[i])
				stack = stack[:len(stack)-1]
				if pair != "[]" && pair != "()" && pair != "{}" {
					return false
				}

			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
