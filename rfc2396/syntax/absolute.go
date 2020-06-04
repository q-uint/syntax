package syntax

import . "github.com/elimity-com/abnf/operators"

func absoluteURI(s []rune) Alternatives {
	return Concat(`absoluteURI`,
		scheme,
		Rune(`:`, ':'),
		Alts(`hier_part | opaque_part`,
			hierPart,
			opaquePart,
		),
	)(s)
}

func hierPart(s []rune) Alternatives {
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

func netPath(s []rune) Alternatives {
	return Concat(`net_path`,
		String(`//`, "//"),
		authority,
		Optional(`[ abs_path ]`, absPath),
	)(s)
}

func absPath(s []rune) Alternatives {
	return Concat(`abs_path`,
		Rune(`/`, '/'),
		pathSegments,
	)(s)
}

func opaquePart(s []rune) Alternatives {
	return Concat(`opaque_part`,
		uricNoSlash,
		Repeat0Inf(`*uric`, uric),
	)(s)
}

func uricNoSlash(s []rune) Alternatives {
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
