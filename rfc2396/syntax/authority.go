package syntax

import . "github.com/di-wu/abnf"

func authority(s []rune) *AST {
	return Alts(`authority`,
		server,
		regName,
	)(s)
}

func regName(s []rune) *AST {
	return Repeat1Inf(`reg_name`,
		Alts(`unreserved | escaped | "$" | "," | ";" | ":" | "@" | "&" | "=" | "+"`,
			unreserved,
			escaped,
			Rune(`$`, '$'),
			Rune(`,`, ','),
			Rune(`;`, ';'),
			Rune(`:`, ':'),
			Rune(`@`, '@'),
			Rune(`&`, '&'),
			Rune(`=`, '='),
			Rune(`+`, '+'),
		),
	)(s)
}

func server(s []rune) *AST {
	return Optional(`server`,
		Concat(`[ userinfo "@" ] hostport`,
			Optional(`[ userinfo "@" ]`,
				Concat(`userinfo "@"`,
					userinfo,
					Rune(`@`, '@'),
				),
			),
			hostport,
		),
	)(s)
}

func userinfo(s []rune) *AST {
	return Repeat1Inf(`userinfo`,
		Alts(`unreserved | escaped | ";" | ":" | "&" | "=" | "+" | "$" | ","`,
			unreserved,
			escaped,
			Rune(`;`, ';'),
			Rune(`:`, ':'),
			Rune(`&`, '&'),
			Rune(`=`, '='),
			Rune(`+`, '+'),
			Rune(`$`, '$'),
			Rune(`,`, ','),
		),
	)(s)
}

func hostport(s []rune) *AST {
	return Concat(`hostport`,
		host,
		Optional(`[ ":" port ]`,
			Concat(`":" port`,
				Rune(`:`, ':'),
				port,
			),
		),
	)(s)
}

func host(s []rune) *AST {
	return Alts(`host`,
		hostname,
		iPv4address,
	)(s)
}

func hostname(s []rune) *AST {
	return Concat(`hostname`,
		Repeat0Inf(`*( domainlabel "." )`,
			Concat(`domainlabel "."`,
				domainlabel,
				Rune(`.`, '.'),
			),
		),
		toplabel,
		Optional(`[ "." ]`, Rune(`.`, '.')),
	)(s)
}

func domainlabel(s []rune) *AST {
	return Alts(`domainlabel`,
		alphanum,
		Concat(`alphanum *( alphanum | "-" ) alphanum`,
			alphanum,
			Repeat0Inf(`*( alphanum | "-" )`,
				Alts(`alphanum | "-"`,
					alphanum,
					Rune(`-`, '-'),
				),
			),
			alphanum,
		),
	)(s)
}

func toplabel(s []rune) *AST {
	return Alts(`toplabel`,
		alpha,
		Concat(`alpha *( alphanum | "-" ) alphanum`,
			alpha,
			Repeat0Inf(`*( alphanum | "-" )`,
				Alts(`alphanum | "-"`,
					alphanum,
					Rune(`-`, '-'),
				),
			),
			alphanum,
		),
	)(s)
}

func iPv4address(s []rune) *AST {
	return Concat(`IPv4address`,
		Repeat1Inf(`1*digit`, digit),
		Rune(`.`, '.'),
		Repeat1Inf(`1*digit`, digit),
		Rune(`.`, '.'),
		Repeat1Inf(`1*digit`, digit),
		Rune(`.`, '.'),
		Repeat1Inf(`1*digit`, digit),
	)(s)
}

func port(s []rune) *AST {
	return Repeat0Inf(`*digit`, digit)(s)
}
