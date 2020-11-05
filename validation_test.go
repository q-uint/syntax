package syntax_test

import (
	"testing"

	"github.com/di-wu/syntax"
	"github.com/di-wu/syntax/rfc3339"
)

func TestIsValid(t *testing.T) {
	if !syntax.IsValid(rfc3339.DateTime, "1985-04-12T23:20:50.52Z") {
		t.Error()
	}
}
