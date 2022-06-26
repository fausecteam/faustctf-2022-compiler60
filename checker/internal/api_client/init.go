package api_client

import (
	"github.com/fausecteam/ctf-gameserver/go/checkerlib"
	"net/http"
)

// setup http.DefaultClient additional to what checkerlib does
func init() {
	// ignore redirects
	http.DefaultClient.CheckRedirect = ignoreRedirects

	transport := http.DefaultTransport.(*http.Transport)
	// the servers do not support gzip
	transport.DisableCompression = true
	// more timeouts and limits
	transport.ResponseHeaderTimeout = checkerlib.Timeout
	transport.MaxResponseHeaderBytes = 1 << 20 // 1 MB like in go server
}

func ignoreRedirects(_ *http.Request, _ []*http.Request) error {
	return http.ErrUseLastResponse
}
