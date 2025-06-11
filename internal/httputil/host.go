package httputil

import "net/http"

func GetHostname(r *http.Request) string {
	hostname := "http://" + r.Host
	if r.TLS != nil {
		hostname = "https://" + r.Host
	}
	return hostname
}
