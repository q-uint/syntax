#!/bin/bash

go run cli/cli.go -rfc=2396
go run cli/cli.go -rfc=3339
go run cli/cli.go -rfc=3986
go run cli/cli.go -rfc=8259

go run cli/cli.go -in="./abnf/iso8601_date.abnf" -out="./iso8601/date.go" -pkg="iso8601" -core="DIGIT"
go run cli/cli.go -in="./abnf/iso8601_duration.abnf" -out="./iso8601/duration.go" -pkg="iso8601" -core="DIGIT"
go run cli/cli.go -in="./abnf/iso8601_period.abnf" -out="./iso8601/period.go" -pkg="iso8601"
go run cli/cli.go -in="./abnf/iso8601_time.abnf" -out="./iso8601/time.go" -pkg="iso8601" -core="DIGIT"
