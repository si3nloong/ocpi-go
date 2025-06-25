package ocpi230

import (
	"net/http"
	"sync"
)

type Client interface {
}

type ClientConn struct {
	rw         sync.RWMutex
	tokenA     string
	tokenC     string
	versionUrl string
	httpClient *http.Client
}

var _ Client = (*ClientConn)(nil)
