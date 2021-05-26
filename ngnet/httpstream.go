package ngnet

import "regexp"

var (
	httpRequestFirtLine  *regexp.Regexp
	httpResponseFirtLine *regexp.Regexp
)

func init() {
	httpRequestFirtLine = regexp.MustCompile(`([A-Z]+) (.+) (HTTP/.+)\r\n`)
	httpResponseFirtLine = regexp.MustCompile(`(HTTP/.+) (\d{3}) (.+)\r\n`)
}
