package httputil

import (
	"net/http"
	"strings"
)

func GetHostname(r *http.Request) string {
	hostname := "http://" + r.Host
	if r.TLS != nil {
		hostname = "https://" + r.Host
		return hostname
	}
	scheme := r.Header.Get("X-Forwarded-Proto")
	if scheme == "https" {
		hostname = "https://" + r.Host
		return hostname
	}
	fwd := r.Header.Get("Forwarded")
	if strings.Contains(fwd, "proto=https") {
		hostname = "https://" + r.Host
		return hostname
	}
	return hostname
}
