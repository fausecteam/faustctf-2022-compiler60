package api_client

import (
	"encoding/json"
	"executor/pkg/compiler60_shared"
	"executor/pkg/signed_elf"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

// MaxStdoutSize in base64 plus some slack for the other fields
const maxResponseBody = 100 + compiler60_shared.MaxStdoutSize/3*4

func getExecuteURL(host string) string {
	return fmt.Sprintf("http://%s/execute", net.JoinHostPort(host, compiler60_shared.ExecutorPort))
}

func ExecuteCode(host string, signedElf *signed_elf.SignedElf) (*compiler60_shared.JsonRunnerResult, error) {
	jsonBytes, err := json.Marshal(signedElf)
	if err != nil {
		log.Printf("Unexpected err from json encoding for execute request: %s\n", err)
		return nil, err
	}

	// use json syntax with spaces like python requests
	jsonStr := string(jsonBytes)
	jsonStr = strings.ReplaceAll(jsonStr, ",", ", ")
	jsonStr = strings.ReplaceAll(jsonStr, ":", ": ")

	resp, err := http.Post(getExecuteURL(host), compiler60_shared.HttpJsonType, strings.NewReader(jsonStr))
	if err != nil {
		log.Printf("Err from execute request: %s\n", err)
		return nil, err
	}

	switch resp.StatusCode {
	default:
		log.Printf("Unknown response status in execute response: %d - %s\n", resp.StatusCode, resp.Status)
		return nil, nil // -> Faulty
	case http.StatusServiceUnavailable:
		// read a bit from body, so you can potentially differ HAProxy from other 503 responses
		var body [32]byte
		_, _ = resp.Body.Read(body[:])
		log.Printf("Got 503 Service Unavailable as compile response: %s - %s\n", resp.Status, string(body[:]))
		return nil, HTTP503Error // -> Down
	case http.StatusOK:
		reader := http.MaxBytesReader(nil, resp.Body, maxResponseBody)
		dec := json.NewDecoder(reader)

		var runnerResult compiler60_shared.JsonRunnerResult
		if err := dec.Decode(&runnerResult); err != nil {
			log.Printf("Json parse err from 200 execute response: %s\n", err)
			return nil, nil // -> Faulty
		}

		return &runnerResult, nil
	}
}
