# ABNF Syntax Parsers
Contains auto-generated ABNF parsers w/ tests.
## Included
- [RFC2396](https://tools.ietf.org/html/rfc2396): Uniform Resource Identifiers (URI): Generic Syntax
- [RFC3986](https://tools.ietf.org/html/rfc3986):  Uniform Resource Identifier (URI): Generic Syntax

### CLI
##### Flags
`i`: source file of the abnf syntax. \
`o`: destination for the generated abnf parser.
```shell
go build cli.go
./cli -i="abnf/rfc3986.abnf" -o="rfc3986/syntax.go"
```
