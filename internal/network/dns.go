package network

import "net"

func DNSQuery(domain string) ([]string, error) {
	return net.LookupTXT(domain)
}