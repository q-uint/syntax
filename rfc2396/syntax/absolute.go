package syntax

import . "github.com/di-wu/abnf"

func absoluteURI(s []rune) *AST {
	return Concat(`absoluteURI`,
		scheme,
		Rune(`:`, ':'),
		Alts(`hier_part | opaque_part`,
			hierPart,
			opaquePart,
		),
	)(s)
}

func hierPart(s []rune) *AST {
	return Concat(`hier_part`,
		Alts(`net_path | abs_path`,
			netPath,
			absPath,
		),
		Optional(`[ "?" query ]`,
			Concat(`"?" query`,
				Rune(`?`, '?'),
				query,
			),
		),
	)(s)
}

func netPath(s []rune) *AST {
	return Concat(`net_path`,
		String(`//`, "//", false),
		authority,
		Optional(`[ abs_path ]`, absPath),
	)(s)
}

func absPath(s []rune) *AST {
	return Concat(`abs_path`,
		Rune(`/`, '/'),
		pathSegments,
	)(s)
}

func opaquePart(s []rune) *AST {
	return Concat(`opaque_part`,
		uricNoSlash,
		Repeat0Inf(`*uric`, uric),
	)(s)
}

func uricNoSlash(s []rune) *AST {
	return Alts(`uric_no_slash`,
		unreserved,
		escaped,
		Rune(`;`, ';'),
		Rune(`?`, '?'),
		Rune(`:`, ':'),
		Rune(`@`, '@'),
		Rune(`&`, '&'),
		Rune(`=`, '='),
		Rune(`+`, '+'),
		Rune(`$`, '$'),
		Rune(`,`, ','),
	)(s)
}
