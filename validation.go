package main

import "github.com/elimity-com/abnf/operators"

func IsValid(operator operators.Operator, value string) bool {
	return operator([]byte(value)).Best() != nil
}
