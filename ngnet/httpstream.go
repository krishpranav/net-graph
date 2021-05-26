package ngnet

import (
	"fmt"
	"regexp"
)

var (
	httpRequestFirtLine  *regexp.Regexp
	httpResponseFirtLine *regexp.Regexp
)

func init() {
	httpRequestFirtLine = regexp.MustCompile(`([A-Z]+) (.+) (HTTP/.+)\r\n`)
	httpResponseFirtLine = regexp.MustCompile(`(HTTP/.+) (\d{3}) (.+)\r\n`)
}

type streamKey struct {
	net, tcp gopacket.Flow
}

func (k streamKey) String() string {
	return fmt.Sprintf("{%v:%v} -> {%v:%v}", k.net.Src(), k.tcp.Src(), k.net.Dst(), k.tcp.Dst())
}
