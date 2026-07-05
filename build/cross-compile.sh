#!/bin/bash

echo "[*] Building AIZEN..."

GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui" -o build/windows/x64/aizen.exe cmd/aizen/main.go
if [ $? -eq 0 ]; then
    echo "[+] Windows x64 built: build/windows/x64/aizen.exe"
    upx --brute build/windows/x64/aizen.exe 2>/dev/null || true
fi

GOOS=windows GOARCH=386 go build -ldflags="-s -w -H=windowsgui" -o build/windows/x86/aizen.exe cmd/aizen/main.go
if [ $? -eq 0 ]; then
    echo "[+] Windows x86 built: build/windows/x86/aizen.exe"
    upx --brute build/windows/x86/aizen.exe 2>/dev/null || true
fi

GOOS=windows GOARCH=amd64 go build -ldflags="-s -w -H=windowsgui" -o build/windows/x64/pdfgen.exe cmd/pdfgen/main.go
if [ $? -eq 0 ]; then
    echo "[+] PDFGen x64 built: build/windows/x64/pdfgen.exe"
fi

echo "[*] Done."