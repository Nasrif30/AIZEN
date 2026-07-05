package pdf

import (
	"crypto/rand"
	"encoding/hex"
	"strings"
)

func ObfuscateJavaScript(js string, level int) string {
	if level <= 0 {
		return js
	}

	lines := strings.Split(js, "\n")
	var clean []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "//") || strings.HasPrefix(trimmed, "/*") {
			continue
		}
		clean = append(clean, line)
	}
	js = strings.Join(clean, "\n")

	vars := extractVars(js)
	for _, v := range vars {
		newName := "_0x" + randomString(4)
		js = strings.ReplaceAll(js, v, newName)
	}

	for i := 0; i < level; i++ {
		junk := generateJunkFunction()
		js = junk + "\n" + js
	}

	js = strings.ReplaceAll(js, "\n", "")
	js = strings.ReplaceAll(js, "\t", "")
	js = strings.ReplaceAll(js, "  ", " ")
	return js
}

func extractVars(js string) []string {
	var vars []string
	seen := make(map[string]bool)

	parts := strings.Split(js, "var ")
	for i := 1; i < len(parts); i++ {
		end := strings.IndexAny(parts[i], " =;\n")
		if end > 0 {
			name := strings.TrimSpace(parts[i][:end])
			if len(name) > 0 && !seen[name] {
				seen[name] = true
				vars = append(vars, name)
			}
		}
	}

	parts = strings.Split(js, "function ")
	for i := 1; i < len(parts); i++ {
		end := strings.IndexAny(parts[i], " (")
		if end > 0 {
			name := strings.TrimSpace(parts[i][:end])
			if len(name) > 0 && !seen[name] {
				seen[name] = true
				vars = append(vars, name)
			}
		}
	}
	return vars
}

func generateJunkFunction() string {
	name := "_0x" + randomString(6)
	return `function ` + name + `(){var a=0;for(var i=0;i<` + randomString(2) + `;i++){a+=i;}return a;}`
}

func randomString(n int) string {
	b := make([]byte, n)
	rand.Read(b)
	return hex.EncodeToString(b)[:n]
}