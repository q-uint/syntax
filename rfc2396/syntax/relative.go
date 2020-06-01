package syntax

import . "github.com/di-wu/abnf"

func relativeURI(s []rune) *AST {
	return Concat(`relativeURI`,
		Alts(`net_path | abs_path | rel_path`,
			netPath,
			absPath,
			relPath,
		),
		Optional(`[ "?" query ]`,
			Concat(`"?" query`,
				Rune(`?`, '?'),
				query,
			),
		),
	)(s)
}

func relPath(s []rune) *AST {
	return Concat(`rel_path`,
		relSegment,
		Optional(`[ abs_path ]`,
			absPath,
		),
	)(s)
}

func relSegment(s []rune) *AST {
	return Repeat1Inf(`rel_segment`,
		Alts(`unreserved | escaped | ";" | "@" | "&" | "=" | "+" | "$" | ","`,
			unreserved,
			escaped,
			Rune(`;`, ';'),
			Rune(`@`, '@'),
			Rune(`&`, '&'),
			Rune(`=`, '='),
			Rune(`+`, '+'),
			Rune(`$`, '$'),
			Rune(`,`, ','),
		),
	)(s)
}
