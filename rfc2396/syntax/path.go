package syntax

import . "github.com/elimity-com/abnf/operators"

func path(s []rune) Alternatives {
	return Optional(`path`,
		Alts(`abs_path | opaque_part`,
			absPath,
			opaquePart,
		),
	)(s)
}

func pathSegments(s []rune) Alternatives {
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

func segment(s []rune) Alternatives {
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

func param(s []rune) Alternatives {
	return Repeat0Inf(`param`,
		pchar,
	)(s)
}

func pchar(s []rune) Alternatives {
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
