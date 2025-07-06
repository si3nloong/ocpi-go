package ocpi221

import (
	"context"
	"errors"

	"github.com/si3nloong/ocpi-go/ocpi"
)

var ErrNotImplemented = errors.New("not implemented")

type UnimplementedServer struct{}

func (UnimplementedServer) IsClientRegistered(ctx context.Context, token string) bool {
	return true
}

// VerifyCredentialsToken implements Credentials.
func (UnimplementedServer) VerifyCredentialsToken(ctx context.Context, token string) bool {
	return true
}

// VerifyCredentialsToken implements Credentials.
func (UnimplementedServer) StoreVersionDetails(ctx context.Context, token VersionDetails) error {
	return nil
}

// GetCredential implements Credentials.
func (UnimplementedServer) OnGetCredential(ctx context.Context, token string) (*Credential, error) {
	return nil, ErrNotImplemented
}

// PostCredential implements Credentials.
func (UnimplementedServer) OnPostCredential(ctx context.Context, token string, body ocpi.RawMessage[Credential]) (*Credential, error) {
	return nil, ErrNotImplemented
}

// PutCredential implements Credentials.
func (UnimplementedServer) OnPutCredential(ctx context.Context, token string, body ocpi.RawMessage[Credential]) (*Credential, error) {
	return nil, ErrNotImplemented
}

// DeleteCredential implements Credentials.
func (UnimplementedServer) OnDeleteCredential(ctx context.Context, token string) error {
	return ErrNotImplemented
}

// OnGetClientOwnedLocation implements EMSP.
func (UnimplementedServer) OnGetClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string) (*Location, error) {
	return nil, ErrNotImplemented
}

// OnGetClientOwnedLocationConnector implements EMSP.
func (UnimplementedServer) OnGetClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string) (*Connector, error) {
	return nil, ErrNotImplemented
}

// OnGetClientOwnedLocationEVSE implements EMSP.
func (UnimplementedServer) OnGetClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string) (*EVSE, error) {
	return nil, ErrNotImplemented
}

// OnPutClientOwnedLocation implements EMSP.
func (UnimplementedServer) OnPutClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[Location]) error {
	return ErrNotImplemented
}

// OnPutClientOwnedLocationConnector implements EMSP.
func (UnimplementedServer) OnPutClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[Location]) error {
	return ErrNotImplemented
}

// OnPutClientOwnedLocationEVSE implements EMSP.
func (UnimplementedServer) OnPutClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[Location]) error {
	return ErrNotImplemented
}

// OnPatchClientOwnedLocation implements EMSP.
func (UnimplementedServer) OnPatchClientOwnedLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[PartialLocation]) error {
	return ErrNotImplemented
}

// OnPatchClientOwnedLocationConnector implements EMSP.
func (UnimplementedServer) OnPatchClientOwnedLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[PartialLocation]) error {
	return ErrNotImplemented
}

// OnPatchClientOwnedLocationEVSE implements EMSP.
func (UnimplementedServer) OnPatchClientOwnedLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[PartialLocation]) error {
	return ErrNotImplemented
}

// OnGetClientOwnedSession implements EMSP.
func (UnimplementedServer) OnGetClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*Session, error) {
	return nil, ErrNotImplemented
}

// OnPutClientOwnedSession implements EMSP.
func (UnimplementedServer) OnPutClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, body ocpi.RawMessage[Session]) error {
	return ErrNotImplemented
}

// OnPatchClientOwnedSession implements EMSP.
func (UnimplementedServer) OnPatchClientOwnedSession(ctx context.Context, countryCode string, partyID string, sessionID string, body ocpi.RawMessage[PartialSession]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) OnGetCDR(ctx context.Context, cdrID string) (*ChargeDetailRecord, error) {
	return nil, ErrNotImplemented
}

// OnPostCDR implements EMSP.
func (UnimplementedServer) OnPostCDR(ctx context.Context, body ocpi.RawMessage[ChargeDetailRecord]) (*ChargeDetailRecordResponse, error) {
	return nil, ErrNotImplemented
}

// OnPostAsyncCommand implements EMSP.
func (UnimplementedServer) OnPostAsyncCommand(ctx context.Context, commandType CommandType, uid string, body ocpi.RawMessage[CommandResult]) error {
	return ErrNotImplemented
}

// OnGetClientOwnedTariff implements EMSP.
func (UnimplementedServer) OnGetClientOwnedTariff(ctx context.Context, countryCode string, partyID string, sessionID string) (*Tariff, error) {
	return nil, ErrNotImplemented
}

// OnPutClientOwnedTariff implements EMSP.
func (UnimplementedServer) OnPutClientOwnedTariff(ctx context.Context, countryCode string, partyID string, tariffID string, body ocpi.RawMessage[Tariff]) error {
	return ErrNotImplemented
}

// OnDeleteClientOwnedTariff implements EMSP.
func (UnimplementedServer) OnDeleteClientOwnedTariff(ctx context.Context, countryCode string, partyID string, tariffID string) error {
	return ErrNotImplemented
}

// OnGetTokens implements EMSP.
func (UnimplementedServer) OnGetTokens(ctx context.Context, params GetTokensParams) (*ocpi.PaginationResponse[Token], error) {
	return nil, ErrNotImplemented
}

// OnPostToken implements EMSP.
func (UnimplementedServer) OnPostToken(ctx context.Context, tokenUID string, body ocpi.RawMessage[LocationReferences], tokenType ...TokenType) (*AuthorizationInfo, error) {
	return nil, ErrNotImplemented
}

// OnGetHubClientInfo implements EMSP.
func (UnimplementedServer) OnGetHubClientInfo(ctx context.Context, countryCode string, partyID string) (*ClientInfo, error) {
	return nil, ErrNotImplemented
}

// OnPutHubClientInfo implements EMSP.
func (UnimplementedServer) OnPutHubClientInfo(ctx context.Context, countryCode string, partyID string, body ocpi.RawMessage[ClientInfo]) error {
	return ErrNotImplemented
}
