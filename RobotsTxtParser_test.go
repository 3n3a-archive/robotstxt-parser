package RobotsTxtParser

import (
	"fmt"
	_ "regexp"
	"testing"

)

func TestParsing(t *testing.T) {
	txtFile := `# ===================================
# Webseite: http://example.com/
# ===================================

Sitemap: http://www.example.com/sitemap.xml

User-agent: ia_archiver
Disallow: /

User-Agent: *
Allow: /

# ===================================
# SchlieÃŸe folgende Spider komplett aus:
# ===================================

User-agent: WebReaper
User-agent: WebCopier
User-agent: Offline Explorer
User-agent: HTTrack
User-agent: Microsoft.URL.Control
User-agent: EmailCollector
User-agent: penthesilea`

	msg, err := ParseTxt(txtFile)
	fmt.Printf("%+v\n", msg)
	if err != nil {
		t.Fatalf(`Msg: %v, Err: %v`, msg, err)
	}
}