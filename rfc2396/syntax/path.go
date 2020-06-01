package syntax

import . "github.com/di-wu/abnf"

func path(s []rune) *AST {
	return Optional(`path`,
		Alts(`abs_path | opaque_part`,
			absPath,
			opaquePart,
		),
	)(s)
}

func pathSegments(s []rune) *AST {
	return Concat(`path_segments`,
		segment,
		Repeat0Inf(`*( "/" segment )`,
			Concat(`"/" segment`,
				Rune(`/`, '/'),
				segment,
			),
		),
	)(s)
}

func segment(s []rune) *AST {
	return Concat(`segment`,
		Repeat0Inf(`*pchar`,
			pchar,
		),
		Repeat0Inf(`*( ";" param )`,
			Concat(`";" param`,
				Rune(`;`, ';'),
				param,
			),
		),
	)(s)
}

func param(s []rune) *AST {
	return Repeat0Inf(`param`,
		pchar,
	)(s)
}

func pchar(s []rune) *AST {
	return Alts(`pchar`,
		unreserved,
		escaped,
		Rune(`:`, ':'),
		Rune(`@`, '@'),
		Rune(`&`, '&'),
		Rune(`=`, '='),
		Rune(`+`, '+'),
		Rune(`$`, '$'),
		Rune(`,`, ','),
	)(s)
}
