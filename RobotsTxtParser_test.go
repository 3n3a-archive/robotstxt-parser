package RobotsTxtParser

import (
	_ "regexp"
	"testing"

)

func TestParsing(t *testing.T) {
	txtFile := `
User-agent: Googlebot
Disallow: /nogooglebot/

User-agent: *
Allow: /

Sitemap: https://www.example.com/sitemap.xml
	`

	msg, err := ParseTxt(txtFile)
	if err != nil {
		t.Fatalf(`Msg: %v, Err: %v`, msg, err)
	}
}