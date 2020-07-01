package network

import (
	"net"
	"strings"
)

func GetAddr() (ip string) {
	addrs, _ := net.InterfaceAddrs()
	addr := addrs[0].String()
	result := strings.Index(addr,"/")
	s := string([]byte(addr)[:result])
	return s
}
