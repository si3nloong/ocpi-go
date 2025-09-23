# ocpi-go

[![Go Report](https://goreportcard.com/badge/github.com/si3nloong/ocpi-go)](https://goreportcard.com/report/github.com/si3nloong/ocpi-go)
[![GoDoc](https://img.shields.io/badge/godoc-reference-5272B4)](https://godoc.org/github.com/si3nloong/ocpi-go)

> Open Charge Point Interface implementation in Go.

## Installation

Go version 1.24+ is required.

```sh
go get github.com/si3nloong/ocpi-go
```

## Development

```go
package main

import (
	"net/http"

	ocpi221 "github.com/si3nloong/ocpi-go/ocpi2.2.1"
)

type server struct {
	ocpi221.UnimplementedServer
}

// GenerateCredentialsTokenC implements ocpi221.OCPIServer.
func (s *server) GenerateCredentialsTokenC(ctx context.Context, tokenA string) (*ocpi221.Credentials, error) {
	panic("unimplemented")
}

// IsClientRegistered implements ocpi221.OCPIServer.
// Subtle: this method shadows the method (UnimplementedServer).IsClientRegistered of server.UnimplementedServer.
func (s *server) IsClientRegistered(ctx context.Context, tokenA string) bool {
	panic("unimplemented")
}

// OnDeleteCredential implements ocpi221.OCPIServer.
// Subtle: this method shadows the method (UnimplementedServer).OnDeleteCredential of server.UnimplementedServer.
func (s *server) OnDeleteCredential(ctx context.Context, tokenC string) error {
	panic("unimplemented")
}

// OnGetCredential implements ocpi221.OCPIServer.
// Subtle: this method shadows the method (UnimplementedServer).OnGetCredential of server.UnimplementedServer.
func (s *server) OnGetCredential(ctx context.Context, tokenC string) (*ocpi221.Credentials, error) {
	panic("unimplemented")
}

// OnPostCredential implements ocpi221.OCPIServer.
// Subtle: this method shadows the method (UnimplementedServer).OnPostCredential of server.UnimplementedServer.
func (s *server) OnPostCredential(ctx context.Context, tokenA string, body ocpi.RawMessage[ocpi221.Credentials]) (*ocpi221.Credentials, error) {
	panic("unimplemented")
}

// OnPutCredential implements ocpi221.OCPIServer.
// Subtle: this method shadows the method (UnimplementedServer).OnPutCredential of server.UnimplementedServer.
func (s *server) OnPutCredential(ctx context.Context, tokenC string, body ocpi.RawMessage[ocpi221.Credentials]) (*ocpi221.Credentials, error) {
	panic("unimplemented")
}

// StoreCredentialsTokenB implements ocpi221.OCPIServer.
func (s *server) StoreCredentialsTokenB(ctx context.Context, credentialsTokenB ocpi221.Credentials) error {
	panic("unimplemented")
}

// StoreVersionDetails implements ocpi221.OCPIServer.
// Subtle: this method shadows the method (UnimplementedServer).StoreVersionDetails of server.UnimplementedServer.
func (s *server) StoreVersionDetails(ctx context.Context, versionDetails ocpi221.VersionDetails) error {
	panic("unimplemented")
}

// VerifyCredentialsToken implements ocpi221.OCPIServer.
// Subtle: this method shadows the method (UnimplementedServer).VerifyCredentialsToken of server.UnimplementedServer.
func (s *server) VerifyCredentialsToken(ctx context.Context, token string) error {
	panic("unimplemented")
}

func main() {
    // Running a OCPI 2.2.1 server
    mux := http.NewServeMux()
	srv := ocpi230.NewServer(&server{}, nil)
	mux.Handle("/ocpi", http.StripPrefix("/ocpi", srv.Handler()))
	http.ListenAndServe(":8080", mux)
}
```

## Features and supported versions
### Supported versions

-   [x] OCPP 2.1.1
-   [x] OCPP 2.2.1
-   [x] OCPI 2.3.0
-   [ ] OCPI 3.0.0

## Tested with Local CPO

| CPO             | Version | Verify and Tested |
| --------------- | ------- | ----------------- |
| Charge Sini     | 2.2.1   | ✅                 |
| Gentari         | 2.2.1   | ✅                 |
| Zura Charge     | 2.1.1   | ✅                 |
| Kineta          | 2.2.1   | ✅                 |
| Electron (TNBX) | 2.3.0   | ✅                 |
| EVlution        | 2.1     | ❌                 |
| chargEV         | 2.1     | ❌                 |
| JomCharge       | 2.1     | ❌                 |

## Big Thanks To

Thanks to these awesome companies for their support of Open Source developers ❤

[![GitHub](https://jstools.dev/img/badges/github.svg)](https://github.com/open-source)

## License

[MIT](https://github.com/si3nloong/ocpi-go/blob/main/LICENSE)

Copyright (c) 2025-present, SianLoong.