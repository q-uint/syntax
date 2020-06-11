// This file is generated - do not edit.

package syntax

import operators "github.com/elimity-com/abnf/operators"

// URI = scheme ":" hier-part [ "?" query ] [ "#" fragment ]
func URI(s []rune) operators.Alternatives {
	return operators.Concat(
		"URI",
		Scheme,
		operators.Rune("\":\"", 58),
		HierPart,
		operators.Optional("[ \"?\" query ]", operators.Concat(
			"\"?\" query",
			operators.Rune("\"?\"", 63),
			Query,
		)),
		operators.Optional("[ \"#\" fragment ]", operators.Concat(
			"\"#\" fragment",
			operators.Rune("\"#\"", 35),
			Fragment,
		)),
	)(s)
}

// HierPart = "//" authority path-abempty / path-absolute / path-rootless / path-empty
func HierPart(s []rune) operators.Alternatives {
	return operators.Alts(
		"hier-part",
		operators.Concat(
			"\"//\" authority path-abempty",
			operators.String("\"//\"", "//"),
			Authority,
			PathAbempty,
		),
		PathAbsolute,
		PathRootless,
		PathEmpty,
	)(s)
}

// URIReference = URI / relative-ref
func URIReference(s []rune) operators.Alternatives {
	return operators.Alts(
		"URI-reference",
		URI,
		RelativeRef,
	)(s)
}

// AbsoluteURI = scheme ":" hier-part [ "?" query ]
func AbsoluteURI(s []rune) operators.Alternatives {
	return operators.Concat(
		"absolute-URI",
		Scheme,
		operators.Rune("\":\"", 58),
		HierPart,
		operators.Optional("[ \"?\" query ]", operators.Concat(
			"\"?\" query",
			operators.Rune("\"?\"", 63),
			Query,
		)),
	)(s)
}

// RelativeRef = relative-part [ "?" query ] [ "#" fragment ]
func RelativeRef(s []rune) operators.Alternatives {
	return operators.Concat(
		"relative-ref",
		RelativePart,
		operators.Optional("[ \"?\" query ]", operators.Concat(
			"\"?\" query",
			operators.Rune("\"?\"", 63),
			Query,
		)),
		operators.Optional("[ \"#\" fragment ]", operators.Concat(
			"\"#\" fragment",
			operators.Rune("\"#\"", 35),
			Fragment,
		)),
	)(s)
}

// RelativePart = "//" authority path-abempty / path-absolute / path-noscheme / path-empty
func RelativePart(s []rune) operators.Alternatives {
	return operators.Alts(
		"relative-part",
		operators.Concat(
			"\"//\" authority path-abempty",
			operators.String("\"//\"", "//"),
			Authority,
			PathAbempty,
		),
		PathAbsolute,
		PathNoscheme,
		PathEmpty,
	)(s)
}

// Scheme = ALPHA *( ALPHA / DIGIT / "+" / "-" / "." )
func Scheme(s []rune) operators.Alternatives {
	return operators.Concat(
		"scheme",
		ALPHA,
		operators.Repeat0Inf("*( ALPHA / DIGIT / \"+\" / \"-\" / \".\" )", operators.Alts(
			"ALPHA / DIGIT / \"+\" / \"-\" / \".\"",
			ALPHA,
			DIGIT,
			operators.Rune("\"+\"", 43),
			operators.Rune("\"-\"", 45),
			operators.Rune("\".\"", 46),
		)),
	)(s)
}

// Authority = [ userinfo "@" ] host [ ":" port ]
func Authority(s []rune) operators.Alternatives {
	return operators.Concat(
		"authority",
		operators.Optional("[ userinfo \"@\" ]", operators.Concat(
			"userinfo \"@\"",
			Userinfo,
			operators.Rune("\"@\"", 64),
		)),
		Host,
		operators.Optional("[ \":\" port ]", operators.Concat(
			"\":\" port",
			operators.Rune("\":\"", 58),
			Port,
		)),
	)(s)
}

// Userinfo = *( unreserved / pct-encoded / sub-delims / ":" )
func Userinfo(s []rune) operators.Alternatives {
	return operators.Repeat0Inf("userinfo", operators.Alts(
		"unreserved / pct-encoded / sub-delims / \":\"",
		Unreserved,
		PctEncoded,
		SubDelims,
		operators.Rune("\":\"", 58),
	))(s)
}

// Host = IP-literal / IPv4address / reg-name
func Host(s []rune) operators.Alternatives {
	return operators.Alts(
		"host",
		IPLiteral,
		IPv4address,
		RegName,
	)(s)
}

// Port = *DIGIT
func Port(s []rune) operators.Alternatives {
	return operators.Repeat0Inf("port", DIGIT)(s)
}

// IPLiteral = "[" ( IPv6address / IPvFuture ) "]"
func IPLiteral(s []rune) operators.Alternatives {
	return operators.Concat(
		"IP-literal",
		operators.Rune("\"[\"", 91),
		operators.Alts(
			"( IPv6address / IPvFuture )",
			IPv6address,
			IPvFuture,
		),
		operators.Rune("\"]\"", 93),
	)(s)
}

// IPvFuture = "v" 1*HEXDIG "." 1*( unreserved / sub-delims / ":" )
func IPvFuture(s []rune) operators.Alternatives {
	return operators.Concat(
		"IPvFuture",
		operators.Rune("\"v\"", 118),
		operators.Repeat1Inf("1*HEXDIG", HEXDIG),
		operators.Rune("\".\"", 46),
		operators.Repeat1Inf("1*( unreserved / sub-delims / \":\" )", operators.Alts(
			"unreserved / sub-delims / \":\"",
			Unreserved,
			SubDelims,
			operators.Rune("\":\"", 58),
		)),
	)(s)
}

// IPv6address = 6( h16 ":" ) ls32 / "::" 5( h16 ":" ) ls32 / [ h16 ] "::" 4( h16 ":" ) ls32 / [ *1( h16 ":" ) h16 ] "::" 3( h16 ":" ) ls32 / [ *2( h16 ":" ) h16 ] "::" 2( h16 ":" ) ls32 / [ *3( h16 ":" ) h16 ] "::" h16 ":" ls32 / [ *4( h16 ":" ) h16 ] "::" ls32 / [ *5( h16 ":" ) h16 ] "::" h16 / [ *6( h16 ":" ) h16 ] "::"
func IPv6address(s []rune) operators.Alternatives {
	return operators.Alts(
		"IPv6address",
		operators.Concat(
			"6( h16 \":\" ) ls32",
			operators.Repeat0Inf("6( h16 \":\" )", operators.Concat(
				"h16 \":\"",
				H16,
				operators.Rune("\":\"", 58),
			)),
			Ls32,
		),
		operators.Concat(
			"\"::\" 5( h16 \":\" ) ls32",
			operators.String("\"::\"", "::"),
			operators.Repeat0Inf("5( h16 \":\" )", operators.Concat(
				"h16 \":\"",
				H16,
				operators.Rune("\":\"", 58),
			)),
			Ls32,
		),
		operators.Concat(
			"[ h16 ] \"::\" 4( h16 \":\" ) ls32",
			operators.Optional("[ h16 ]", H16),
			operators.String("\"::\"", "::"),
			operators.Repeat0Inf("4( h16 \":\" )", operators.Concat(
				"h16 \":\"",
				H16,
				operators.Rune("\":\"", 58),
			)),
			Ls32,
		),
		operators.Concat(
			"[ *1( h16 \":\" ) h16 ] \"::\" 3( h16 \":\" ) ls32",
			operators.Optional("[ *1( h16 \":\" ) h16 ]", operators.Concat(
				"*1( h16 \":\" ) h16",
				operators.Repeat("*1( h16 \":\" )", 0, 1, operators.Concat(
					"*1( h16 \":\" )",
					H16,
					operators.Rune("\":\"", 58),
				)),
				H16,
			)),
			operators.String("\"::\"", "::"),
			operators.Repeat0Inf("3( h16 \":\" )", operators.Concat(
				"h16 \":\"",
				H16,
				operators.Rune("\":\"", 58),
			)),
			Ls32,
		),
		operators.Concat(
			"[ *2( h16 \":\" ) h16 ] \"::\" 2( h16 \":\" ) ls32",
			operators.Optional("[ *2( h16 \":\" ) h16 ]", operators.Concat(
				"*2( h16 \":\" ) h16",
				operators.Repeat("*2( h16 \":\" )", 0, 2, operators.Concat(
					"*2( h16 \":\" )",
					H16,
					operators.Rune("\":\"", 58),
				)),
				H16,
			)),
			operators.String("\"::\"", "::"),
			operators.Repeat0Inf("2( h16 \":\" )", operators.Concat(
				"h16 \":\"",
				H16,
				operators.Rune("\":\"", 58),
			)),
			Ls32,
		),
		operators.Concat(
			"[ *3( h16 \":\" ) h16 ] \"::\" h16 \":\" ls32",
			operators.Optional("[ *3( h16 \":\" ) h16 ]", operators.Concat(
				"*3( h16 \":\" ) h16",
				operators.Repeat("*3( h16 \":\" )", 0, 3, operators.Concat(
					"*3( h16 \":\" )",
					H16,
					operators.Rune("\":\"", 58),
				)),
				H16,
			)),
			operators.String("\"::\"", "::"),
			H16,
			operators.Rune("\":\"", 58),
			Ls32,
		),
		operators.Concat(
			"[ *4( h16 \":\" ) h16 ] \"::\" ls32",
			operators.Optional("[ *4( h16 \":\" ) h16 ]", operators.Concat(
				"*4( h16 \":\" ) h16",
				operators.Repeat("*4( h16 \":\" )", 0, 4, operators.Concat(
					"*4( h16 \":\" )",
					H16,
					operators.Rune("\":\"", 58),
				)),
				H16,
			)),
			operators.String("\"::\"", "::"),
			Ls32,
		),
		operators.Concat(
			"[ *5( h16 \":\" ) h16 ] \"::\" h16",
			operators.Optional("[ *5( h16 \":\" ) h16 ]", operators.Concat(
				"*5( h16 \":\" ) h16",
				operators.Repeat("*5( h16 \":\" )", 0, 5, operators.Concat(
					"*5( h16 \":\" )",
					H16,
					operators.Rune("\":\"", 58),
				)),
				H16,
			)),
			operators.String("\"::\"", "::"),
			H16,
		),
		operators.Concat(
			"[ *6( h16 \":\" ) h16 ] \"::\"",
			operators.Optional("[ *6( h16 \":\" ) h16 ]", operators.Concat(
				"*6( h16 \":\" ) h16",
				operators.Repeat("*6( h16 \":\" )", 0, 6, operators.Concat(
					"*6( h16 \":\" )",
					H16,
					operators.Rune("\":\"", 58),
				)),
				H16,
			)),
			operators.String("\"::\"", "::"),
		),
	)(s)
}

// H16 = 1*4HEXDIG
func H16(s []rune) operators.Alternatives {
	return operators.Repeat("h16", 1, 4, HEXDIG)(s)
}

// Ls32 = ( h16 ":" h16 ) / IPv4address
func Ls32(s []rune) operators.Alternatives {
	return operators.Alts(
		"ls32",
		operators.Concat(
			"( h16 \":\" h16 )",
			H16,
			operators.Rune("\":\"", 58),
			H16,
		),
		IPv4address,
	)(s)
}

// IPv4address = dec-octet "." dec-octet "." dec-octet "." dec-octet
func IPv4address(s []rune) operators.Alternatives {
	return operators.Concat(
		"IPv4address",
		DecOctet,
		operators.Rune("\".\"", 46),
		DecOctet,
		operators.Rune("\".\"", 46),
		DecOctet,
		operators.Rune("\".\"", 46),
		DecOctet,
	)(s)
}

// DecOctet = DIGIT
func DecOctet(s []rune) operators.Alternatives {
	return operators.Alts(
		"dec-octet",
		DIGIT,
		operators.Concat(
			"%x31-39 DIGIT",
			operators.Range("%x31-39", 49, 57),
			DIGIT,
		),
		operators.Concat(
			"\"1\" 2DIGIT",
			operators.Rune("\"1\"", 49),
			operators.Repeat0Inf("2DIGIT", DIGIT),
		),
		operators.Concat(
			"\"2\" %x30-34 DIGIT",
			operators.Rune("\"2\"", 50),
			operators.Range("%x30-34", 48, 52),
			DIGIT,
		),
		operators.Concat(
			"\"25\" %x30-35",
			operators.String("\"25\"", "25"),
			operators.Range("%x30-35", 48, 53),
		),
	)(s)
}

// RegName = *( unreserved / pct-encoded / sub-delims )
func RegName(s []rune) operators.Alternatives {
	return operators.Repeat0Inf("reg-name", operators.Alts(
		"unreserved / pct-encoded / sub-delims",
		Unreserved,
		PctEncoded,
		SubDelims,
	))(s)
}

// Path = path-abempty
func Path(s []rune) operators.Alternatives {
	return operators.Alts(
		"path",
		PathAbempty,
		PathAbsolute,
		PathNoscheme,
		PathRootless,
		PathEmpty,
	)(s)
}

// PathAbempty = *( "/" segment )
func PathAbempty(s []rune) operators.Alternatives {
	return operators.Repeat0Inf("path-abempty", operators.Concat(
		"\"/\" segment",
		operators.Rune("\"/\"", 47),
		Segment,
	))(s)
}

// PathAbsolute = "/" [ segment-nz *( "/" segment ) ]
func PathAbsolute(s []rune) operators.Alternatives {
	return operators.Concat(
		"path-absolute",
		operators.Rune("\"/\"", 47),
		operators.Optional("[ segment-nz *( \"/\" segment ) ]", operators.Concat(
			"segment-nz *( \"/\" segment )",
			SegmentNz,
			operators.Repeat0Inf("*( \"/\" segment )", operators.Concat(
				"\"/\" segment",
				operators.Rune("\"/\"", 47),
				Segment,
			)),
		)),
	)(s)
}

// PathNoscheme = segment-nz-nc *( "/" segment )
func PathNoscheme(s []rune) operators.Alternatives {
	return operators.Concat(
		"path-noscheme",
		SegmentNzNc,
		operators.Repeat0Inf("*( \"/\" segment )", operators.Concat(
			"\"/\" segment",
			operators.Rune("\"/\"", 47),
			Segment,
		)),
	)(s)
}

// PathRootless = segment-nz *( "/" segment )
func PathRootless(s []rune) operators.Alternatives {
	return operators.Concat(
		"path-rootless",
		SegmentNz,
		operators.Repeat0Inf("*( \"/\" segment )", operators.Concat(
			"\"/\" segment",
			operators.Rune("\"/\"", 47),
			Segment,
		)),
	)(s)
}

// PathEmpty = 0pchar
func PathEmpty(s []rune) operators.Alternatives {
	return operators.Repeat0Inf("path-empty", Pchar)(s)
}

// Segment = *pchar
func Segment(s []rune) operators.Alternatives {
	return operators.Repeat0Inf("segment", Pchar)(s)
}

// SegmentNz = 1*pchar
func SegmentNz(s []rune) operators.Alternatives {
	return operators.Repeat1Inf("segment-nz", Pchar)(s)
}

// SegmentNzNc = 1*( unreserved / pct-encoded / sub-delims / "@" )
func SegmentNzNc(s []rune) operators.Alternatives {
	return operators.Repeat1Inf("segment-nz-nc", operators.Alts(
		"unreserved / pct-encoded / sub-delims / \"@\"",
		Unreserved,
		PctEncoded,
		SubDelims,
		operators.Rune("\"@\"", 64),
	))(s)
}

// Pchar = unreserved / pct-encoded / sub-delims / ":" / "@"
func Pchar(s []rune) operators.Alternatives {
	return operators.Alts(
		"pchar",
		Unreserved,
		PctEncoded,
		SubDelims,
		operators.Rune("\":\"", 58),
		operators.Rune("\"@\"", 64),
	)(s)
}

// Query = *( pchar / "/" / "?" )
func Query(s []rune) operators.Alternatives {
	return operators.Repeat0Inf("query", operators.Alts(
		"pchar / \"/\" / \"?\"",
		Pchar,
		operators.Rune("\"/\"", 47),
		operators.Rune("\"?\"", 63),
	))(s)
}

// Fragment = *( pchar / "/" / "?" )
func Fragment(s []rune) operators.Alternatives {
	return operators.Repeat0Inf("fragment", operators.Alts(
		"pchar / \"/\" / \"?\"",
		Pchar,
		operators.Rune("\"/\"", 47),
		operators.Rune("\"?\"", 63),
	))(s)
}

// PctEncoded = "%" HEXDIG HEXDIG
func PctEncoded(s []rune) operators.Alternatives {
	return operators.Concat(
		"pct-encoded",
		operators.Rune("\"%\"", 37),
		HEXDIG,
		HEXDIG,
	)(s)
}

// Unreserved = ALPHA / DIGIT / "-" / "." / "_" / "~"
func Unreserved(s []rune) operators.Alternatives {
	return operators.Alts(
		"unreserved",
		ALPHA,
		DIGIT,
		operators.Rune("\"-\"", 45),
		operators.Rune("\".\"", 46),
		operators.Rune("\"_\"", 95),
		operators.Rune("\"~\"", 126),
	)(s)
}

// Reserved = gen-delims / sub-delims
func Reserved(s []rune) operators.Alternatives {
	return operators.Alts(
		"reserved",
		GenDelims,
		SubDelims,
	)(s)
}

// GenDelims = ":" / "/" / "?" / "#" / "[" / "]" / "@"
func GenDelims(s []rune) operators.Alternatives {
	return operators.Alts(
		"gen-delims",
		operators.Rune("\":\"", 58),
		operators.Rune("\"/\"", 47),
		operators.Rune("\"?\"", 63),
		operators.Rune("\"#\"", 35),
		operators.Rune("\"[\"", 91),
		operators.Rune("\"]\"", 93),
		operators.Rune("\"@\"", 64),
	)(s)
}

// SubDelims = "!" / "$" / "&" / "'" / "(" / ")" / "*" / "+" / "," / "
func SubDelims(s []rune) operators.Alternatives {
	return operators.Alts(
		"sub-delims",
		operators.Rune("\"!\"", 33),
		operators.Rune("\"$\"", 36),
		operators.Rune("\"&\"", 38),
		operators.Rune("\"'\"", 39),
		operators.Rune("\"(\"", 40),
		operators.Rune("\")\"", 41),
		operators.Rune("\"*\"", 42),
		operators.Rune("\"+\"", 43),
		operators.Rune("\",\"", 44),
		operators.Rune("\";\"", 59),
		operators.Rune("\"=\"", 61),
	)(s)
}

// ALPHA = %x41-5A / %x61-7A
func ALPHA(s []rune) operators.Alternatives {
	return operators.Alts(
		"ALPHA",
		operators.Range("%x41-5A", 65, 90),
		operators.Range("%x61-7A", 97, 122),
	)(s)
}

// DIGIT = %x30-39
func DIGIT(s []rune) operators.Alternatives {
	return operators.Range("DIGIT", 48, 57)(s)
}

// HEXDIG = DIGIT / "A" / "B" / "C" / "D" / "E" / "F" / "a" / "b" / "c" / "d" / "e" / "f"
func HEXDIG(s []rune) operators.Alternatives {
	return operators.Alts(
		"HEXDIG",
		DIGIT,
		operators.Rune("\"A\"", 65),
		operators.Rune("\"B\"", 66),
		operators.Rune("\"C\"", 67),
		operators.Rune("\"D\"", 68),
		operators.Rune("\"E\"", 69),
		operators.Rune("\"F\"", 70),
		operators.Rune("\"a\"", 97),
		operators.Rune("\"b\"", 98),
		operators.Rune("\"c\"", 99),
		operators.Rune("\"d\"", 100),
		operators.Rune("\"e\"", 101),
		operators.Rune("\"f\"", 102),
	)(s)
}
