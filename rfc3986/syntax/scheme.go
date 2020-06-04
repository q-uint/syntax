package syntax

import . "github.com/elimity-com/abnf/operators"

func scheme(s []rune) Alternatives {
	return Concat(`scheme`,
		alpha,
		Repeat0Inf(`*( ALPHA / DIGIT / "+" / "-" / "." )`,
			Alts(`ALPHA / DIGIT / "+" / "-" / "."`,
				alpha,
				digit,
				Rune(`+`, '+'),
				Rune(`-`, '-'),
				Rune(`.`, '.'),
			),
		),
	)(s)
}
