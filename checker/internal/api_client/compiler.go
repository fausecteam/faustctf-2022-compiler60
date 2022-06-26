package api_client

import (
	"encoding/json"
	"executor/pkg/signed_elf"
	"fmt"
	"log"
	"net"
	"net/http"
	"strings"
)

const CompilerPort = "6061"

const TextPlain = "text/plain"

func getCompileURL(host string) string {
	return fmt.Sprintf("http://%s/compile", net.JoinHostPort(host, CompilerPort))
}

func CompileCode(host string, code string) (*signed_elf.SignedElf, *CompilerError, error) {
	resp, err := http.Post(getCompileURL(host), TextPlain, strings.NewReader(code))
	if err != nil {
		log.Printf("Err from compile request: %s\n", err)
		return nil, nil, err
	}

	reader := http.MaxBytesReader(nil, resp.Body, signed_elf.MaxEncodedSize)
	dec := json.NewDecoder(reader)
	switch resp.StatusCode {
	default:
		log.Printf("Unknown response status in compile response: %d - %s\n", resp.StatusCode, resp.Status)
		return nil, nil, nil // -> Faulty
	case http.StatusServiceUnavailable:
		// read a bit from body, so you can potentially differ HAProxy from other 503 responses
		var body [32]byte
		_, _ = reader.Read(body[:])
		log.Printf("Got 503 Service Unavailable as compile response: %s - %s\n", resp.Status, string(body[:]))
		return nil, nil, HTTP503Error // -> Down
	case http.StatusOK:
		// successful compiler
		var signedElf signed_elf.SignedElf
		if err := dec.Decode(&signedElf); err != nil {
			log.Printf("Json parse err from 200 compile response: %s\n", err)
			return nil, nil, nil // -> Faulty
		}

		if err := signedElf.VerifyStructure(); err != nil {
			log.Printf("Elf structure wrong from 200 compile response: %s\n", err)
			return nil, nil, nil // -> Faulty
		}
		return &signedElf, nil, nil
	case http.StatusBadRequest:
		// expecting compile error
		compilerErr := CompilerError{Line: -1}
		if err := dec.Decode(&compilerErr); err != nil {
			log.Printf("Json parse err from 400 compile response: %s\n", err)
			return nil, nil, nil // -> Faulty
		}

		if len(compilerErr.Error) == 0 {
			log.Printf("No 'err' key for 400 compile response\n")
			return nil, nil, nil // -> Faulty
		}
		return nil, &compilerErr, nil
	}
}
