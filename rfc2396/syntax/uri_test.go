package syntax

import (
	"testing"
)

// https://tools.ietf.org/html/rfc2396#section-1.3
func TestExampleURI(t *testing.T) {
	// ftp scheme for File Transfer Protocol services
	t.Run("ftp", func(t *testing.T) {
		ftp := "ftp://ftp.is.co.za/rfc/rfc1808.txt"
		tree := absoluteURI([]rune(ftp)).Best()
		if tree.String() != ftp {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "ftp" {
			t.Errorf("did not get correct ftp scheme: %s", scheme)
		}

		if authority := tree.GetSubNode("authority").String(); authority != "ftp.is.co.za" {
			t.Errorf("did not get correct authority: %s", authority)
		}

		abs := tree.GetSubNode("abs_path")
		if abs.String() != "/rfc/rfc1808.txt" {
			t.Errorf("did not get correct absolute path: %s", abs)
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

	// gopher scheme for Gopher and Gopher+ Protocol services
	t.Run("gopher", func(t *testing.T) {
		gopher := "gopher://spinaltap.micro.umn.edu/00/Weather/California/Los%20Angeles"
		tree := absoluteURI([]rune(gopher)).Best()
		if tree.String() != gopher {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "gopher" {
			t.Errorf("did not get correct gopher scheme: %s", scheme)
		}

		if authority := tree.GetSubNode("authority").String(); authority != "spinaltap.micro.umn.edu" {
			t.Errorf("did not get correct authority: %s", authority)
		}

		abs := tree.GetSubNode("abs_path")
		if abs.String() != "/00/Weather/California/Los%20Angeles" {
			t.Errorf("did not get correct absolute path: %s", abs)
		}

		segments := abs.GetSubNodes("segment")
		if len(segments) != 4 {
			t.Error("expected four segments")
			return
		}

		if s := segments[0].String(); s != "00" {
			t.Errorf("invalid segment: %s", s)
		}

		if s := segments[1].String(); s != "Weather" {
			t.Errorf("invalid segment: %s", s)
		}

		if s := segments[2].String(); s != "California" {
			t.Errorf("invalid segment: %s", s)
		}

		if s := segments[3].String(); s != "Los%20Angeles" {
			t.Errorf("invalid segment: %s", s)
		}
	})

	// http scheme for Hypertext Transfer Protocol services
	t.Run("http", func(t *testing.T) {
		http := "http://www.math.uio.no/faq/compression-faq/part1.html"
		tree := absoluteURI([]rune(http)).Best()
		if tree.String() != http {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "http" {
			t.Errorf("did not get correct http scheme: %s", scheme)
		}

		if authority := tree.GetSubNode("authority").String(); authority != "www.math.uio.no" {
			t.Errorf("did not get correct authority: %s", authority)
		}

		abs := tree.GetSubNode("abs_path")
		if abs.String() != "/faq/compression-faq/part1.html" {
			t.Errorf("did not get correct absolute path: %s", abs)
		}

		segments := abs.GetSubNodes("segment")
		if len(segments) != 3 {
			t.Error("expected three segments")
			return
		}

		if s := segments[0].String(); s != "faq" {
			t.Errorf("invalid segment: %s", s)
		}

		if s := segments[1].String(); s != "compression-faq" {
			t.Errorf("invalid segment: %s", s)
		}

		if s := segments[2].String(); s != "part1.html" {
			t.Errorf("invalid segment: %s", s)
		}
	})

	// mailto scheme for electronic mail addresses
	t.Run("mailto", func(t *testing.T) {
		mailto := "mailto:mduerst@ifi.unizh.ch"
		tree := absoluteURI([]rune(mailto)).Best()
		if tree.String() != mailto {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "mailto" {
			t.Errorf("did not get correct mailto scheme: %s", scheme)
		}

		if s := tree.GetSubNode("opaque_part").String(); s != "mduerst@ifi.unizh.ch" {
			t.Errorf("invalid opaque part: %s", s)
		}
	})

	// news scheme for USENET news groups and articles
	t.Run("news", func(t *testing.T) {
		news := "news:comp.infosystems.www.servers.unix"
		tree := absoluteURI([]rune(news)).Best()
		if tree.String() != news {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "news" {
			t.Errorf("did not get correct news scheme: %s", scheme)
		}

		if s := tree.GetSubNode("opaque_part").String(); s != "comp.infosystems.www.servers.unix" {
			t.Errorf("invalid opaque part: %s", s)
		}
	})

	// telnet scheme for interactive services via the TELNET Protocol
	t.Run("telnet", func(t *testing.T) {
		telnet := "telnet://melvyl.ucop.edu/"
		tree := absoluteURI([]rune(telnet)).Best()
		if tree.String() != telnet {
			t.Error("could not parse string")
		}

		if scheme := tree.GetSubNode("scheme").String(); scheme != "telnet" {
			t.Errorf("did not get correct telnet scheme: %s", scheme)
		}

		if authority := tree.GetSubNode("authority").String(); authority != "melvyl.ucop.edu" {
			t.Errorf("did not get correct authority: %s", authority)
		}

		abs := tree.GetSubNode("abs_path")
		if abs.String() != "/" {
			t.Errorf("did not get correct absolute path: %s", abs)
		}

		segments := abs.GetSubNodes("segment")
		if len(segments) != 1 {
			t.Error("expected one segments")
		}

		if s := segments[0].String(); s != "" {
			t.Errorf("invalid segment: %s", s)
		}
	})
}
