package jsonhelper

import (
	"github.com/savsgio/gotils/strconv"
	"regexp"
)

func FixJsonNewLines(data []byte) []byte {
	str := strconv.B2S(data)
	matchNewlines := regexp.MustCompile(`[\r\n]`)
	escapeNewlines := func(s string) string {
		return matchNewlines.ReplaceAllString(s, "\\n")
	}
	re := regexp.MustCompile(`"[^"\\]*(?:\\[\s\S][^"\\]*)*"`)
	fixedJson := re.ReplaceAllStringFunc(str, escapeNewlines)
	return strconv.S2B(fixedJson)
}
