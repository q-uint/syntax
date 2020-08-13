package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/elimity-com/abnf"
)

func main() {
	rfcNumber := flag.Int("rfc", 0, "the number of the rfc spec. e.g. 2396")

	fileInput := flag.String("in", "", "the file path to the abnf spec.")
	fileOutput := flag.String("out", "", "the file path to the output file.")
	packageName := flag.String("pkg", "syntax", "the name of the go package.")
	operators := flag.String("core", "", "list of core operators seperated by a comma")
	flag.Parse()

	if n := *rfcNumber; n != 0 {
		rfc(n)
	} else {
		other(*fileInput, *fileOutput, *packageName, strings.Split(*operators, ",")...)
	}
}

func rfc(n int) {
	raw, err := ioutil.ReadFile(fmt.Sprintf("./abnf/rfc%d.abnf", n))
	if err != nil {
		log.Fatal(err)
	}

	g := abnf.CodeGenerator{
		PackageName: fmt.Sprintf("rfc%d", n),
		RawABNF:     raw,
	}

	externalABNF, ok := dependencies[n]
	if ok {
		g.ExternalABNF = externalABNF
	}

	f := g.GenerateABNFAsAlternatives()
	if err := ioutil.WriteFile(
		fmt.Sprintf("./rfc%d/syntax.go", n),
		[]byte(fmt.Sprintf("%#v", f)), 0644,
	); err != nil {
		log.Fatal(err)
	}
}

func other(in, out, pkg string, operators ...string) {
	raw, err := ioutil.ReadFile(in)
	if err != nil {
		log.Fatal(err)
	}

	g := abnf.CodeGenerator{
		PackageName: pkg,
		RawABNF:     raw,
	}

	externalABNF := make(map[string]abnf.ExternalABNF)
	for _, op := range operators {
		if e, ok := coreOperators[op]; ok {
			externalABNF[op] = e
		}
	}
	g.ExternalABNF = externalABNF

	f := g.GenerateABNFAsAlternatives()
	if err := ioutil.WriteFile(
		out,
		[]byte(fmt.Sprintf("%#v", f)), 0644,
	); err != nil {
		log.Fatal(err)
	}
}

var (
	corePkg = abnf.ExternalABNF{
		IsOperator:  true,
		PackageName: "github.com/elimity-com/abnf/core",
	}

	coreOperators = map[string]abnf.ExternalABNF{
		"DIGIT":  corePkg,
		"ALPHA":  corePkg,
		"HEXDIG": corePkg,
	}

	dependencies = map[int]map[string]abnf.ExternalABNF{
		3339: {
			"DIGIT": corePkg,
		},
		3986: {
			"ALPHA":  corePkg,
			"DIGIT":  corePkg,
			"HEXDIG": corePkg,
		},
		8259: {
			"DIGIT":  corePkg,
			"HEXDIG": corePkg,
		},
	}
)
