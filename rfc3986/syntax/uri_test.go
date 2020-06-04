package syntax

import (
	"testing"
)

// https://tools.ietf.org/html/rfc3986#section-1.1.2
func TestExampleURI(t *testing.T) {
	t.Run("ftp", func(t *testing.T) {
		ftp := "ftp://ftp.is.co.za/rfc/rfc1808.txt"
		tree := absoluteURI([]rune(ftp)).Best()
		if tree.String() != ftp {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "ftp" {
			t.Errorf("did not get the correct ftp scheme: %s", scheme)
		}

		if authority := tree.GetSubNode("authority").String(); authority != "ftp.is.co.za" {
			t.Errorf("did not get the correct authority: %s", authority)
		}

		abs := tree.GetSubNode("path-abempty")
		if abs.String() != "/rfc/rfc1808.txt" {
			t.Errorf("did not get the correct absolute path: %s", abs)
		}

		segments := abs.GetSubNodes("segment")
		if len(segments) != 2 {
			t.Error("expected two segments")
			return
		}

		if s := segments[0].String(); s != "rfc" {
			t.Errorf("invalid segment: %s", s)
		}

		if s := segments[1].String(); s != "rfc1808.txt" {
			t.Errorf("invalid segment: %s", s)
		}
	})

	t.Run("http", func(t *testing.T) {
		http := "http://www.ietf.org/rfc/rfc2396.txt"
		tree := absoluteURI([]rune(http)).Best()
		if tree.String() != http {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "http" {
			t.Errorf("did not get the correct http scheme: %s", scheme)
		}

		if authority := tree.GetSubNode("authority").String(); authority != "www.ietf.org" {
			t.Errorf("did not get the correct authority: %s", authority)
		}

		abs := tree.GetSubNode("path-abempty")
		if abs.String() != "/rfc/rfc2396.txt" {
			t.Errorf("did not get the correct absolute path: %s", abs)
		}

		segments := abs.GetSubNodes("segment")
		if len(segments) != 2 {
			t.Error("expected two segments")
			return
		}

		if s := segments[0].String(); s != "rfc" {
			t.Errorf("invalid segment: %s", s)
		}

		if s := segments[1].String(); s != "rfc2396.txt" {
			t.Errorf("invalid segment: %s", s)
		}
	})

	t.Run("ldap", func(t *testing.T) {
		ldap := "ldap://[2001:db8::7]/c=GB?objectClass?one"
		tree := absoluteURI([]rune(ldap)).Best()
		if tree.String() != ldap {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "ldap" {
			t.Errorf("did not get the correct ldap scheme: %s", scheme)
		}

		if ip := tree.GetSubNode("IP-literal").String(); ip != "[2001:db8::7]" {
			t.Errorf("did not get the correct ip literal: %s", ip)
		}

		abs := tree.GetSubNode("path-abempty")
		if abs.String() != "/c=GB" {
			t.Errorf("did not get the correct absolute path: %s", abs)
		}

		segments := abs.GetSubNodes("segment")
		if len(segments) != 1 {
			t.Error("expected one segment")
			return
		}

		if s := segments[0].String(); s != "c=GB" {
			t.Errorf("invalid segment: %s", s)
		}

		query := tree.GetSubNode("query")
		if query.String() != "objectClass?one" {
			t.Errorf("did not get the corrent query: %s", query)
		}
	})

	t.Run("mailto", func(t *testing.T) {
		mailto := "mailto:John.Doe@example.com"
		tree := absoluteURI([]rune(mailto)).Best()
		if tree.String() != mailto {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "mailto" {
			t.Errorf("did not get the correct mailto scheme: %s", scheme)
		}

		if rootless := tree.GetSubNode("path-rootless").String(); rootless != "John.Doe@example.com" {
			t.Errorf("did not get the correct rootless path: %s", rootless)
		}
	})
	
	t.Run("news", func(t *testing.T) {
		news := "news:comp.infosystems.www.servers.unix"
		tree := absoluteURI([]rune(news)).Best()
		if tree.String() != news {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "news" {
			t.Errorf("did not get the correct news scheme: %s", scheme)
		}

		if rootless := tree.GetSubNode("path-rootless").String(); rootless != "comp.infosystems.www.servers.unix" {
			t.Errorf("did not get the correct rootless path: %s", rootless)
		}
	})


	t.Run("tel", func(t *testing.T) {
		tel := "tel:+1-816-555-1212"
		tree := absoluteURI([]rune(tel)).Best()
		if tree.String() != tel {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "tel" {
			t.Errorf("did not get the correct tel scheme: %s", scheme)
		}

		if rootless := tree.GetSubNode("path-rootless").String(); rootless != "+1-816-555-1212" {
			t.Errorf("did not get the correct rootless path: %s", rootless)
		}
	})

	t.Run("telnet", func(t *testing.T) {
		telnet := "telnet://192.0.2.16:80/"
		tree := absoluteURI([]rune(telnet)).Best()
		if tree.String() != telnet {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "telnet" {
			t.Errorf("did not get the correct telnet scheme: %s", scheme)
		}

		authority := tree.GetSubNode("authority")
		if authority.String() != "192.0.2.16:80" {
			t.Errorf("did not get the correct authority: %s", authority)
		}

		if ipv4 := authority.GetSubNode("IPv4address"); ipv4.String() != "192.0.2.16" {
			t.Errorf("invalid IPv4address: %s", ipv4)
		}

		if port := authority.GetSubNode("port"); port.String() != "80" {
			t.Errorf("invalid port: %s", port)
		}
	})

	t.Run("urn", func(t *testing.T) {
		urn := "urn:oasis:names:specification:docbook:dtd:xml:4.1.2"
		tree := absoluteURI([]rune(urn)).Best()
		if tree.String() != urn {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "urn" {
			t.Errorf("did not get the correct urn scheme: %s", scheme)
		}

		if rootless := tree.GetSubNode("path-rootless").String(); rootless != "oasis:names:specification:docbook:dtd:xml:4.1.2" {
			t.Errorf("did not get the correct rootless path: %s", rootless)
		}
	})
}
