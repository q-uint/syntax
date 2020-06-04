package syntax

import . "github.com/elimity-com/abnf/operators"

func fragment(s []rune) Alternatives {
	return Repeat0Inf(`fragment`,
		Alts(`pchar / "/" / "?"`,
			pchar,
			Rune(`/`, '/'),
			Rune(`?`, '?'),
		),
	)(s)
}
