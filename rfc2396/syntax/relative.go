package syntax

import . "github.com/elimity-com/abnf/operators"

func relativeURI(s []rune) Alternatives {
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

func relPath(s []rune) Alternatives {
	return Concat(`rel_path`,
		relSegment,
		Optional(`[ abs_path ]`,
			absPath,
		),
	)(s)
}

func relSegment(s []rune) Alternatives {
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
