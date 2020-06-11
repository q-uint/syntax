// This file is generated - do not edit.

package syntax

import operators "github.com/elimity-com/abnf/operators"

// URIReference = [ absoluteURI / relativeURI ] [ "#" fragment ]
func URIReference(s []rune) operators.Alternatives {
	return operators.Concat(
		"URI-reference",
		operators.Optional("[ absoluteURI / relativeURI ]", operators.Alts(
			"absoluteURI / relativeURI",
			AbsoluteURI,
			RelativeURI,
		)),
		operators.Optional("[ \"#\" fragment ]", operators.Concat(
			"\"#\" fragment",
			operators.Rune("\"#\"", 35),
			Fragment,
		)),
	)(s)
}

// AbsoluteURI = scheme ":" ( hier-part / opaque-part )
func AbsoluteURI(s []rune) operators.Alternatives {
	return operators.Concat(
		"absoluteURI",
		Scheme,
		operators.Rune("\":\"", 58),
		operators.Alts(
			"( hier-part / opaque-part )",
			HierPart,
			OpaquePart,
		),
	)(s)
}

// RelativeURI = ( net-path / abs-path / rel-path ) [ "?" query ]
func RelativeURI(s []rune) operators.Alternatives {
	return operators.Concat(
		"relativeURI",
		operators.Alts(
			"( net-path / abs-path / rel-path )",
			NetPath,
			AbsPath,
			RelPath,
		),
		operators.Optional("[ \"?\" query ]", operators.Concat(
			"\"?\" query",
			operators.Rune("\"?\"", 63),
			Query,
		)),
	)(s)
}

// HierPart = ( net-path / abs-path ) [ "?" query ]
func HierPart(s []rune) operators.Alternatives {
	return operators.Concat(
		"hier-part",
		operators.Alts(
			"( net-path / abs-path )",
			NetPath,
			AbsPath,
		),
		operators.Optional("[ \"?\" query ]", operators.Concat(
			"\"?\" query",
			operators.Rune("\"?\"", 63),
			Query,
		)),
	)(s)
}

// OpaquePart = uric-no-slash *uric
func OpaquePart(s []rune) operators.Alternatives {
	return operators.Concat(
		"opaque-part",
		UricNoSlash,
		operators.Repeat0Inf("*uric", Uric),
	)(s)
}

// UricNoSlash = unreserved / escaped / "
func UricNoSlash(s []rune) operators.Alternatives {
	return operators.Alts(
		"uric-no-slash",
		Unreserved,
		Escaped,
		operators.Rune("\";\"", 59),
		operators.Rune("\"?\"", 63),
		operators.Rune("\":\"", 58),
		operators.Rune("\"@\"", 64),
		operators.Rune("\"&\"", 38),
		operators.Rune("\"=\"", 61),
		operators.Rune("\"+\"", 43),
		operators.Rune("\"$\"", 36),
		operators.Rune("\",\"", 44),
	)(s)
}

// NetPath = "//" authority [ abs-path ]
func NetPath(s []rune) operators.Alternatives {
	return operators.Concat(
		"net-path",
		operators.String("\"//\"", "//"),
		Authority,
		operators.Optional("[ abs-path ]", AbsPath),
	)(s)
}

// AbsPath = "/" path-segments
func AbsPath(s []rune) operators.Alternatives {
	return operators.Concat(
		"abs-path",
		operators.Rune("\"/\"", 47),
		PathSegments,
	)(s)
}

// RelPath = rel-segment [ abs-path ]
func RelPath(s []rune) operators.Alternatives {
	return operators.Concat(
		"rel-path",
		RelSegment,
		operators.Optional("[ abs-path ]", AbsPath),
	)(s)
}

// RelSegment = 1*( unreserved / escaped / "
func RelSegment(s []rune) operators.Alternatives {
	return operators.Repeat1Inf("rel-segment", operators.Alts(
		"unreserved / escaped / \";\" / \"@\" / \"&\" / \"=\" / \"+\" / \"$\" / \",\"",
		Unreserved,
		Escaped,
		operators.Rune("\";\"", 59),
		operators.Rune("\"@\"", 64),
		operators.Rune("\"&\"", 38),
		operators.Rune("\"=\"", 61),
		operators.Rune("\"+\"", 43),
		operators.Rune("\"$\"", 36),
		operators.Rune("\",\"", 44),
	))(s)
}

// Scheme = alpha *( alpha / digit / "+" / "-" / "." )
func Scheme(s []rune) operators.Alternatives {
	return operators.Concat(
		"scheme",
		Alpha,
		operators.Repeat0Inf("*( alpha / digit / \"+\" / \"-\" / \".\" )", operators.Alts(
			"alpha / digit / \"+\" / \"-\" / \".\"",
			Alpha,
			Digit,
			operators.Rune("\"+\"", 43),
			operators.Rune("\"-\"", 45),
			operators.Rune("\".\"", 46),
		)),
	)(s)
}

// Authority = server / reg-name
func Authority(s []rune) operators.Alternatives {
	return operators.Alts(
		"authority",
		Server,
		RegName,
	)(s)
}

// RegName = 1*( unreserved / escaped / "$" / "," / "
func RegName(s []rune) operators.Alternatives {
	return operators.Repeat1Inf("reg-name", operators.Alts(
		"unreserved / escaped / \"$\" / \",\" / \";\" / \":\" / \"@\" / \"&\" / \"=\" / \"+\"",
		Unreserved,
		Escaped,
		operators.Rune("\"$\"", 36),
		operators.Rune("\",\"", 44),
		operators.Rune("\";\"", 59),
		operators.Rune("\":\"", 58),
		operators.Rune("\"@\"", 64),
		operators.Rune("\"&\"", 38),
		operators.Rune("\"=\"", 61),
		operators.Rune("\"+\"", 43),
	))(s)
}

// Server = [ [ userinfo "@" ] hostport ]
func Server(s []rune) operators.Alternatives {
	return operators.Optional("server", operators.Concat(
		"[ userinfo \"@\" ] hostport",
		operators.Optional("[ userinfo \"@\" ]", operators.Concat(
			"userinfo \"@\"",
			Userinfo,
			operators.Rune("\"@\"", 64),
		)),
		Hostport,
	))(s)
}

// Userinfo = *( unreserved / escaped / "
func Userinfo(s []rune) operators.Alternatives {
	return operators.Repeat0Inf("userinfo", operators.Alts(
		"unreserved / escaped / \";\" / \":\" / \"&\" / \"=\" / \"+\" / \"$\" / \",\"",
		Unreserved,
		Escaped,
		operators.Rune("\";\"", 59),
		operators.Rune("\":\"", 58),
		operators.Rune("\"&\"", 38),
		operators.Rune("\"=\"", 61),
		operators.Rune("\"+\"", 43),
		operators.Rune("\"$\"", 36),
		operators.Rune("\",\"", 44),
	))(s)
}

// Hostport = host [ ":" port ]
func Hostport(s []rune) operators.Alternatives {
	return operators.Concat(
		"hostport",
		Host,
		operators.Optional("[ \":\" port ]", operators.Concat(
			"\":\" port",
			operators.Rune("\":\"", 58),
			Port,
		)),
	)(s)
}

// Host = hostname / IPv4address
func Host(s []rune) operators.Alternatives {
	return operators.Alts(
		"host",
		Hostname,
		IPv4address,
	)(s)
}

// Hostname = *( domainlabel "." ) toplabel [ "." ]
func Hostname(s []rune) operators.Alternatives {
	return operators.Concat(
		"hostname",
		operators.Repeat0Inf("*( domainlabel \".\" )", operators.Concat(
			"domainlabel \".\"",
			Domainlabel,
			operators.Rune("\".\"", 46),
		)),
		Toplabel,
		operators.Optional("[ \".\" ]", operators.Rune("\".\"", 46)),
	)(s)
}

// Domainlabel = alphanum / alphanum *( alphanum / "-" ) alphanum
func Domainlabel(s []rune) operators.Alternatives {
	return operators.Alts(
		"domainlabel",
		Alphanum,
		operators.Concat(
			"alphanum *( alphanum / \"-\" ) alphanum",
			Alphanum,
			operators.Repeat0Inf("*( alphanum / \"-\" )", operators.Alts(
				"alphanum / \"-\"",
				Alphanum,
				operators.Rune("\"-\"", 45),
			)),
			Alphanum,
		),
	)(s)
}

// Toplabel = alpha / alpha *( alphanum / "-" ) alphanum
func Toplabel(s []rune) operators.Alternatives {
	return operators.Alts(
		"toplabel",
		Alpha,
		operators.Concat(
			"alpha *( alphanum / \"-\" ) alphanum",
			Alpha,
			operators.Repeat0Inf("*( alphanum / \"-\" )", operators.Alts(
				"alphanum / \"-\"",
				Alphanum,
				operators.Rune("\"-\"", 45),
			)),
			Alphanum,
		),
	)(s)
}

// IPv4address = 1*digit "." 1*digit "." 1*digit "." 1*digit
func IPv4address(s []rune) operators.Alternatives {
	return operators.Concat(
		"IPv4address",
		operators.Repeat1Inf("1*digit", Digit),
		operators.Rune("\".\"", 46),
		operators.Repeat1Inf("1*digit", Digit),
		operators.Rune("\".\"", 46),
		operators.Repeat1Inf("1*digit", Digit),
		operators.Rune("\".\"", 46),
		operators.Repeat1Inf("1*digit", Digit),
	)(s)
}

// Port = *digit
func Port(s []rune) operators.Alternatives {
	return operators.Repeat0Inf("port", Digit)(s)
}

// Path = [ abs-path / opaque-part ]
func Path(s []rune) operators.Alternatives {
	return operators.Optional("path", operators.Alts(
		"abs-path / opaque-part",
		AbsPath,
		OpaquePart,
	))(s)
}

// PathSegments = segment *( "/" segment )
func PathSegments(s []rune) operators.Alternatives {
	return operators.Concat(
		"path-segments",
		Segment,
		operators.Repeat0Inf("*( \"/\" segment )", operators.Concat(
			"\"/\" segment",
			operators.Rune("\"/\"", 47),
			Segment,
		)),
	)(s)
}

// Segment = *pchar *( "
func Segment(s []rune) operators.Alternatives {
	return operators.Concat(
		"segment",
		operators.Repeat0Inf("*pchar", Pchar),
		operators.Repeat0Inf("*( \";\" param )", operators.Concat(
			"\";\" param",
			operators.Rune("\";\"", 59),
			Param,
		)),
	)(s)
}

// Param = *pchar
func Param(s []rune) operators.Alternatives {
	return operators.Repeat0Inf("param", Pchar)(s)
}

// Pchar = unreserved / escaped / ":" / "@" / "&" / "=" / "+" / "$" / ","
func Pchar(s []rune) operators.Alternatives {
	return operators.Alts(
		"pchar",
		Unreserved,
		Escaped,
		operators.Rune("\":\"", 58),
		operators.Rune("\"@\"", 64),
		operators.Rune("\"&\"", 38),
		operators.Rune("\"=\"", 61),
		operators.Rune("\"+\"", 43),
		operators.Rune("\"$\"", 36),
		operators.Rune("\",\"", 44),
	)(s)
}

// Query = *uric
func Query(s []rune) operators.Alternatives {
	return operators.Repeat0Inf("query", Uric)(s)
}

// Fragment = *uric
func Fragment(s []rune) operators.Alternatives {
	return operators.Repeat0Inf("fragment", Uric)(s)
}

// Uric = reserved / unreserved / escaped
func Uric(s []rune) operators.Alternatives {
	return operators.Alts(
		"uric",
		Reserved,
		Unreserved,
		Escaped,
	)(s)
}

// Reserved = "
func Reserved(s []rune) operators.Alternatives {
	return operators.Alts(
		"reserved",
		operators.Rune("\";\"", 59),
		operators.Rune("\"/\"", 47),
		operators.Rune("\"?\"", 63),
		operators.Rune("\":\"", 58),
		operators.Rune("\"@\"", 64),
		operators.Rune("\"&\"", 38),
		operators.Rune("\"=\"", 61),
		operators.Rune("\"+\"", 43),
		operators.Rune("\"$\"", 36),
		operators.Rune("\",\"", 44),
	)(s)
}

// Unreserved = alphanum / mark
func Unreserved(s []rune) operators.Alternatives {
	return operators.Alts(
		"unreserved",
		Alphanum,
		Mark,
	)(s)
}

// Mark = "-" / "-" / "." / "!" / "~" / "*" / "'" / "(" / ")"
func Mark(s []rune) operators.Alternatives {
	return operators.Alts(
		"mark",
		operators.Rune("\"-\"", 45),
		operators.Rune("\"-\"", 45),
		operators.Rune("\".\"", 46),
		operators.Rune("\"!\"", 33),
		operators.Rune("\"~\"", 126),
		operators.Rune("\"*\"", 42),
		operators.Rune("\"'\"", 39),
		operators.Rune("\"(\"", 40),
		operators.Rune("\")\"", 41),
	)(s)
}

// Escaped = "%" hex hex
func Escaped(s []rune) operators.Alternatives {
	return operators.Concat(
		"escaped",
		operators.Rune("\"%\"", 37),
		Hex,
		Hex,
	)(s)
}

// Hex = digit / "A" / "B" / "C" / "D" / "E" / "F" / "a" / "b" / "c" / "d" / "e" / "f"
func Hex(s []rune) operators.Alternatives {
	return operators.Alts(
		"hex",
		Digit,
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

// Alphanum = alpha / digit
func Alphanum(s []rune) operators.Alternatives {
	return operators.Alts(
		"alphanum",
		Alpha,
		Digit,
	)(s)
}

// Alpha = lowalpha / upalpha
func Alpha(s []rune) operators.Alternatives {
	return operators.Alts(
		"alpha",
		Lowalpha,
		Upalpha,
	)(s)
}

// Lowalpha = "a" / "b" / "c" / "d" / "e" / "f" / "g" / "h" / "i" / "j" / "k" / "l" / "m" / "n" / "o" / "p" / "q" / "r" / "s" / "t" / "u" / "v" / "w" / "x" / "y" / "z"
func Lowalpha(s []rune) operators.Alternatives {
	return operators.Alts(
		"lowalpha",
		operators.Rune("\"a\"", 97),
		operators.Rune("\"b\"", 98),
		operators.Rune("\"c\"", 99),
		operators.Rune("\"d\"", 100),
		operators.Rune("\"e\"", 101),
		operators.Rune("\"f\"", 102),
		operators.Rune("\"g\"", 103),
		operators.Rune("\"h\"", 104),
		operators.Rune("\"i\"", 105),
		operators.Rune("\"j\"", 106),
		operators.Rune("\"k\"", 107),
		operators.Rune("\"l\"", 108),
		operators.Rune("\"m\"", 109),
		operators.Rune("\"n\"", 110),
		operators.Rune("\"o\"", 111),
		operators.Rune("\"p\"", 112),
		operators.Rune("\"q\"", 113),
		operators.Rune("\"r\"", 114),
		operators.Rune("\"s\"", 115),
		operators.Rune("\"t\"", 116),
		operators.Rune("\"u\"", 117),
		operators.Rune("\"v\"", 118),
		operators.Rune("\"w\"", 119),
		operators.Rune("\"x\"", 120),
		operators.Rune("\"y\"", 121),
		operators.Rune("\"z\"", 122),
	)(s)
}

// Upalpha = "A" / "B" / "C" / "D" / "E" / "F" / "G" / "H" / "I" / "J" / "K" / "L" / "M" / "N" / "O" / "P" / "Q" / "R" / "S" / "T" / "U" / "V" / "W" / "X" / "Y" / "Z"
func Upalpha(s []rune) operators.Alternatives {
	return operators.Alts(
		"upalpha",
		operators.Rune("\"A\"", 65),
		operators.Rune("\"B\"", 66),
		operators.Rune("\"C\"", 67),
		operators.Rune("\"D\"", 68),
		operators.Rune("\"E\"", 69),
		operators.Rune("\"F\"", 70),
		operators.Rune("\"G\"", 71),
		operators.Rune("\"H\"", 72),
		operators.Rune("\"I\"", 73),
		operators.Rune("\"J\"", 74),
		operators.Rune("\"K\"", 75),
		operators.Rune("\"L\"", 76),
		operators.Rune("\"M\"", 77),
		operators.Rune("\"N\"", 78),
		operators.Rune("\"O\"", 79),
		operators.Rune("\"P\"", 80),
		operators.Rune("\"Q\"", 81),
		operators.Rune("\"R\"", 82),
		operators.Rune("\"S\"", 83),
		operators.Rune("\"T\"", 84),
		operators.Rune("\"U\"", 85),
		operators.Rune("\"V\"", 86),
		operators.Rune("\"W\"", 87),
		operators.Rune("\"X\"", 88),
		operators.Rune("\"Y\"", 89),
		operators.Rune("\"Z\"", 90),
	)(s)
}

// Digit = "0" / "1" / "2" / "3" / "4" / "5" / "6" / "7" / "8" / "9"
func Digit(s []rune) operators.Alternatives {
	return operators.Alts(
		"digit",
		operators.Rune("\"0\"", 48),
		operators.Rune("\"1\"", 49),
		operators.Rune("\"2\"", 50),
		operators.Rune("\"3\"", 51),
		operators.Rune("\"4\"", 52),
		operators.Rune("\"5\"", 53),
		operators.Rune("\"6\"", 54),
		operators.Rune("\"7\"", 55),
		operators.Rune("\"8\"", 56),
		operators.Rune("\"9\"", 57),
	)(s)
}
