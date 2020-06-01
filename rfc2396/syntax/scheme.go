package syntax

import . "github.com/di-wu/abnf"

func scheme(s []rune) *AST {
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
