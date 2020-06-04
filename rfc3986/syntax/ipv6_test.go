package syntax

import (
	"testing"
)

func TestIPv6Examples(t *testing.T) {
	// https://tools.ietf.org/html/rfc5952#section-1
	for _, address := range []string{
		"2001:db8:0:0:1:0:0:1",
		"2001:0db8:0:0:1:0:0:1",
		"2001:db8::1:0:0:1",
		"2001:db8::0:1:0:0:1",
		"2001:0db8::1:0:0:1",
		"2001:db8:0:0:1::1",
		"2001:db8:0000:0:1::1",
		"2001:DB8:0:0:1::1",
	} {
		t.Run(address, func(t *testing.T) {
			a := iPv6address([]rune(address)).Best()
			if a == nil {
				t.Error("could not parse IPv6 address")
				return
			}
			if a.String() != address {
				t.Errorf("invalid IPv6 address: %s", a)
			}
		})
	}
}
