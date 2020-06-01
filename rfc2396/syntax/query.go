package syntax

import . "github.com/di-wu/abnf"

func query(s []rune) *AST {
	return Repeat0Inf(`query`,
		uric,
	)(s)
}
