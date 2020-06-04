package syntax

import . "github.com/elimity-com/abnf/operators"

func path(s []rune) Alternatives {
	return Alts(`path`,
		pathAbempty,
		pathAbsolute,
		pathNoscheme,
		pathRootless,
		pathEmpty,
	)(s)
}

func pathAbempty(s []rune) Alternatives {
	return Repeat0Inf(`path-abempty`,
		Concat(`"/" segment`,
			Rune(`/`, '/'),
			segment,
		),
	)(s)
}

func pathAbsolute(s []rune) Alternatives {
	return Concat(`path-absolute`,
		Rune(`/`, '/'),
		Optional(`[ segment-nz *( "/" segment ) ]`,
			Concat(`segment-nz *( "/" segment )`,
				segmentNz,
				Repeat0Inf(`*( "/" segment )`,
					Concat(`"/" segment`,
						Rune(`/`, '/'),
						segment,
					),
				),
			),
		),
	)(s)
}

func pathNoscheme(s []rune) Alternatives {
	return Concat(`path-noscheme`,
		segmentNzNc,
		Repeat0Inf(`*( "/" segment )`,
			Concat(`"/" segment`,
				Rune(`/`, '/'),
				segment,
			),
		),
	)(s)
}

func pathRootless(s []rune) Alternatives {
	return Concat(`path-rootless`,
		segmentNz,
		Repeat0Inf(`*( "/" segment )`,
			Concat(`"/" segment`,
				Rune(`/`, '/'),
				segment,
			),
		),
	)(s)
}

func pathEmpty(s []rune) Alternatives {
	return RepeatN(`path-empty`, 0, pchar)(s)
}

func segment(s []rune) Alternatives {
	return Repeat0Inf(`segment`, pchar)(s)
}

func segmentNz(s []rune) Alternatives {
	return Repeat1Inf(`segment-nz`, pchar)(s)
}

func segmentNzNc(s []rune) Alternatives {
	return Repeat1Inf(`segment-nz-nc`,
		Alts(`unreserved / pct-encoded / sub-delims / "@"`,
			unreserved,
			pctEncoded,
			subDelims,
			Rune(`@`, '@'),
		),
	)(s)
}

func pchar(s []rune) Alternatives {
	return Alts(`pchar`,
		unreserved,
		pctEncoded,
		subDelims,
		Rune(`:`, ':'),
		Rune(`@`, '@'),
	)(s)
}
