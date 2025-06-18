package ocpi221

import (
	"context"
	"errors"

	"github.com/si3nloong/ocpi-go/ocpi"
)

var ErrNotImplemented = errors.New("not implemented")

type UnimplementedServer struct{}

// GetCredential implements Credentials.
func (UnimplementedServer) GetCredential(ctx context.Context, tokenA string) (*Credential, error) {
	return nil, ErrNotImplemented
}

// PostCredential implements Credentials.
func (UnimplementedServer) PostCredential(ctx context.Context, tokenA string, body ocpi.RawMessage[Credential]) (*Credential, error) {
	return nil, ErrNotImplemented
}

// PutCredential implements Credentials.
func (UnimplementedServer) PutCredential(ctx context.Context, tokenA string, body ocpi.RawMessage[Credential]) (*Credential, error) {
	return nil, ErrNotImplemented
}

// DeleteCredential implements Credentials.
func (UnimplementedServer) DeleteCredential(ctx context.Context, tokenA string) (*Credential, error) {
	return nil, ErrNotImplemented
}

// VerifyToken implements Credentials.
func (UnimplementedServer) VerifyToken(token string) bool {
	return true
}

func (UnimplementedServer) GetLocation(ctx context.Context, countryCode string, partyID string, locationID string) (*Location, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) GetLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string) (*EVSE, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) GetLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string) (*Connector, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PutLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[Location]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PutLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[Location]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PutLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[Location]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PatchLocation(ctx context.Context, countryCode string, partyID string, locationID string, body ocpi.RawMessage[PatchedLocation]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PatchLocationEVSE(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, body ocpi.RawMessage[PatchedLocation]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PatchLocationConnector(ctx context.Context, countryCode string, partyID string, locationID string, evseUID string, connectorID string, body ocpi.RawMessage[PatchedLocation]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) GetTariff(ctx context.Context, countryCode string, partyID string, sessionID string) (*Tariff, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PutTariff(ctx context.Context, countryCode string, partyID string, tariffID string, body ocpi.RawMessage[Tariff]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PatchTariff(ctx context.Context, countryCode string, partyID string, tariffID string, body ocpi.RawMessage[Tariff]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) GetSession(ctx context.Context, countryCode string, partyID string, sessionID string) (*Session, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PutSession(ctx context.Context, countryCode string, partyID string, sessionID string, body ocpi.RawMessage[Session]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PatchSession(ctx context.Context, countryCode string, partyID string, sessionID string, body ocpi.RawMessage[PatchedSession]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) GetToken(ctx context.Context, countryCode string, partyID string, tokenUID string, tokenType ...TokenType) (*Token, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PutToken(ctx context.Context, countryCode string, partyID string, tokenUID string, body ocpi.RawMessage[Token]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PatchToken(ctx context.Context, countryCode string, partyID string, tokenUID string, body ocpi.RawMessage[PatchedToken]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) GetCDR(ctx context.Context, cdrID string) (*ChargeDetailRecord, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PostCDR(ctx context.Context, body ocpi.RawMessage[ChargeDetailRecord]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) GetTerminal(ctx context.Context, terminalID string) (*Terminal, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) GetFinancialAdviceConfirmation(ctx context.Context, financialAdviceConfirmationID string) (*FinancialAdviceConfirmation, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PostTerminal(ctx context.Context, body ocpi.RawMessage[Terminal]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PostFinancialAdviceConfirmation(ctx context.Context, body ocpi.RawMessage[FinancialAdviceConfirmation]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) GetLocations(ctx context.Context, params GetLocationsParams) ([]Location, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) GetTariffs(ctx context.Context, params GetTariffsParams) ([]Tariff, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) GetSessions(ctx context.Context, params GetOcpiSessionsParams) ([]Session, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) GetTokens(ctx context.Context, params GetTokensParams) ([]Token, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PostToken(ctx context.Context, tokenUID string, body ocpi.RawMessage[LocationReferences], tokenType ...TokenType) (*AuthorizationInfo, error) {
	return nil, ErrNotImplemented
}
