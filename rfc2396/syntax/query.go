package syntax

import . "github.com/elimity-com/abnf/operators"

func query(s []rune) Alternatives {
	return Repeat0Inf(`query`,
		uric,
	)(s)
}
