# ABNF Syntax Parsers
Contains auto-generated ABNF parsers w/ tests.
## Included
- [RFC2396](https://tools.ietf.org/html/rfc2396): Uniform Resource Identifiers (URI): Generic Syntax
- [RFC3339](https://tools.ietf.org/html/rfc3339): Date and Time on the Internet: Timestamps
- [RFC3986](https://tools.ietf.org/html/rfc3986): Uniform Resource Identifier (URI): Generic Syntax
- [RFC7159](https://tools.ietf.org/html/rfc7159) /
  [RFC8259](https://tools.ietf.org/html/rfc8259): The JavaScript Object Notation (JSON) Data Interchange Format

- [ISO8601](https://tools.ietf.org/html/rfc3339#appendix-A): ISO 8601 Collected ABNF (! based on the 1988 version)

## Validation
```go
import "github.com/di-wu/syntax/rfc3339"

if !IsValid(rfc3339.DateTime, "1985-04-12T23:20:50.52Z") {
    // date time is not valid
}
```

##### CLI
```shell script
go run cli/cli.go -rfc=3339
```

```shell script
go run cli/cli.go -in="./abnf/iso8601_date.abnf" -out="./iso8601/date.go" -pkg="iso8601" -core="DIGIT"
```
