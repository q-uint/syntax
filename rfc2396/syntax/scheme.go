package syntax

import . "github.com/elimity-com/abnf/operators"

func scheme(s []rune) Alternatives {
	return Concat(`scheme`,
		alpha,
		Repeat0Inf(`*( alpha | digit | "+" | "-" | "." )`,
			Alts(`alpha | digit | "+" | "-" | "."`,
				alpha,
				digit,
				Rune(`+`, '+'),
				Rune(`-`, '-'),
				Rune(`.`, '.'),
			),
		),
	)(s)
}
