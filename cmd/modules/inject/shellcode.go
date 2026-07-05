package inject

import (
	"encoding/hex"
	"os"
)

func LoadShellcode(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	if len(data) > 2 && data[0] == '0' && data[1] == 'x' {
		return hex.DecodeString(string(data[2:]))
	}
	if len(data) > 0 && ((data[0] >= '0' && data[0] <= '9') || (data[0] >= 'a' && data[0] <= 'f') || (data[0] >= 'A' && data[0] <= 'F')) {
		decoded, err := hex.DecodeString(string(data))
		if err == nil {
			return decoded, nil
		}
	}
	return data, nil
}

func DecryptShellcode(encrypted []byte, key []byte) []byte {
	decrypted := make([]byte, len(encrypted))
	for i := 0; i < len(encrypted); i++ {
		decrypted[i] = encrypted[i] ^ key[i%len(key)]
	}
	return decrypted
}