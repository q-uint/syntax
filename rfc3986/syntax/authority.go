package syntax

import . "github.com/elimity-com/abnf/operators"

func authority(s []rune) Alternatives {
	return Concat(`authority`,
		Optional(`[ userinfo "@" ]`,
			Concat(`userinfo "@"`,
				userinfo,
				Rune(`@`, '@'),
			),
		),
		host,
		Optional(`[ ":" port ]`,
			Concat(`":" port`,
				Rune(`:`, ':'),
				port,
			),
		),
	)(s)
}

func userinfo(s []rune) Alternatives {
	return Repeat1Inf(`userinfo`,
		Alts(`unreserved / pct-encoded / sub-delims / ":"`,
			unreserved,
			pctEncoded,
			subDelims,
			Rune(`:`, ':'),
		),
	)(s)
}

func host(s []rune) Alternatives {
	return Alts(`host`,
		iPLiteral,
		iPv4address,
		regName,
	)(s)
}

func iPLiteral(s []rune) Alternatives {
	return Concat(`IP-literal`,
		Rune(`[`, '['),
		Alts(`IPv6address / IPvFuture `,
			iPv6address,
			iPvFuture,
		),
		Rune(`]`, ']'),
	)(s)
}

func iPvFuture(s []rune) Alternatives {
	return Concat(`IPvFuture`,
		Rune(`v`, 'v'),
		Repeat1Inf(`1*HEXDIG`, hexdig),
		Rune(`.`, '.'),
		Repeat1Inf(`1*( unreserved / sub-delims / ":" )`,
			Alts(`unreserved / sub-delims / ":"`,
				unreserved,
				subDelims,
				Rune(`:`, ':'),
			),
		),
	)(s)
}

func iPv6address(s []rune) Alternatives {
	h16colon := func() Operator {
		return Concat(`h16 ":"`,
			h16,
			Rune(`:`, ':'),
		)
	}

	return Alts(`IPv6address`,
		Concat(`6( h16 ":" ) ls32`,
			RepeatN(`6( h16 ":" )`, 6, h16colon()),
			ls32,
		),
		Concat(`"::" 5( h16 ":" ) ls32`,
			String(`::`, "::"),
			RepeatN(`5( h16 ":" )`, 5, h16colon()),
			ls32,
		),
		Concat(`[ h16 ] "::" 4( h16 ":" ) ls32`,
			Optional(`[ h16 ]`, h16),
			String(`::`, "::"),
			RepeatN(`4( h16 ":" )`, 4, h16colon()),
			ls32,
		),
		Concat(`[ *1( h16 ":" ) h16 ] "::" 3( h16 ":" ) ls32`,
			Optional(`[ *1( h16 ":" ) h16 ]`,
				Concat(`*1( h16 ":" ) h16`,
					Repeat(`*1( h16 ":" )`, 0, 1, h16colon()),
					h16,
				),
			),
			String(`::`, "::"),
			RepeatN(`3( h16 ":" )`, 3, h16colon()),
			ls32,
		),
		Concat(`[ *2( h16 ":" ) h16 ] "::" 2( h16 ":" ) ls32`,
			Optional(`[ *2( h16 ":" ) h16 ]`,
				Concat(`*2( h16 ":" ) h16`,
					Repeat(`*2( h16 ":" )`, 0, 2, h16colon()),
					h16,
				),
			),
			String(`::`, "::"),
			RepeatN(`2( h16 ":" )`, 2, h16colon()),
			ls32,
		),
		Concat(`[ *3( h16 ":" ) h16 ] "::" h16 ":" ls32`,
			Optional(`[ *3( h16 ":" ) h16 ]`,
				Concat(`*3( h16 ":" ) h16`,
					Repeat(`*3( h16 ":" )`, 0, 3, h16colon()),
					h16,
				),
			),
			String(`::`, "::"),
			h16colon(),
			ls32,
		),
		Concat(`[ *4( h16 ":" ) h16 ] "::" ls32`,
			Optional(`[ *4( h16 ":" ) h16 ]`,
				Concat(`*4( h16 ":" ) h16`,
					Repeat(`*4( h16 ":" )`, 0, 4, h16colon()),
					h16,
				),
			),
			String(`::`, "::"),
			ls32,
		),
		Concat(`[ *5( h16 ":" ) h16 ] "::" h16`,
			Optional(`[ *5( h16 ":" ) h16 ]`,
				Concat(`*5( h16 ":" ) h16`,
					Repeat(`*5( h16 ":" )`, 0, 5, h16colon()),
					h16,
				),
			),
			String(`::`, "::"),
			h16,
		),
		Concat(`[ *6( h16 ":" ) h16 ] "::"`,
			Optional(`[ *6( h16 ":" ) h16 ]`,
				Concat(`*6( h16 ":" ) h16`,
					Repeat(`*6( h16 ":" )`, 0, 6, h16colon()),
					h16,
				),
			),
			String(`::`, "::"),
		),
	)(s)
}

func ls32(s []rune) Alternatives {
	return Alts(`ls32`,
		Concat(`h16 ":" h16`,
			h16,
			Rune(`:`, ':'),
			h16,
		),
		iPv4address,
	)(s)
}

func h16(s []rune) Alternatives {
	return Repeat(`h16`, 1, 4, hexdig)(s)
}

func iPv4address(s []rune) Alternatives {
	return Concat(`IPv4address`,
		decOcted,
		Rune(`.`, '.'),
		decOcted,
		Rune(`.`, '.'),
		decOcted,
		Rune(`.`, '.'),
		decOcted,
	)(s)
}

func decOcted(s []rune) Alternatives {
	return Alts(`dec-octet`,
		digit,
		Concat(`10-99`,
			Range(`1-9`, '1', '9'),
			digit,
		),
		Concat(`100-199`,
			Rune(`1`, '1'),
			RepeatN(`2DIGIT`, 2, digit),
		),
		Concat(`200-249`,
			Rune(`2`, '2'),
			Range(`0-4`, '0', '4'),
			digit,
		),
		Concat(`250-255`,
			String(`25`, "25"),
			Range(`0-5`, '0', '5'),
		),
	)(s)
}

func regName(s []rune) Alternatives {
	return Repeat0Inf(`reg-name`,
		Alts(`unreserved / pct-encoded / sub-delims`,
			unreserved,
			pctEncoded,
			subDelims,
		),
	)(s)
}

func port(s []rune) Alternatives {
	return Repeat0Inf(`port`, digit)(s)
}

