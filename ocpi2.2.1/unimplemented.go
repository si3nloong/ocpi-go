package ocpi221

import (
	"context"
	"errors"

	"github.com/si3nloong/ocpi-go/ocpi"
)

var ErrNotImplemented = errors.New("not implemented")

type UnimplementedServer struct{}

func (UnimplementedServer) GetLocation(ctx context.Context, countryCode string, partyId string, locationId string) (*Location, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) GetLocationEVSE(ctx context.Context, countryCode string, partyId string, locationId string, evseUid string) (*EVSE, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) GetLocationConnector(ctx context.Context, countryCode string, partyId string, locationId string, evseUid string, connectorId string) (*Connector, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PutLocation(ctx context.Context, countryCode string, partyId string, locationId string, body ocpi.RawMessage[Location]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PutLocationEVSE(ctx context.Context, countryCode string, partyId string, locationId string, evseUid string, body ocpi.RawMessage[Location]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PutLocationConnector(ctx context.Context, countryCode string, partyId string, locationId string, evseUid string, connectorId string, body ocpi.RawMessage[Location]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PatchLocation(ctx context.Context, countryCode string, partyId string, locationId string, body ocpi.RawMessage[PatchedLocation]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PatchLocationEVSE(ctx context.Context, countryCode string, partyId string, locationId string, evseUid string, body ocpi.RawMessage[PatchedLocation]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PatchLocationConnector(ctx context.Context, countryCode string, partyId string, locationId string, evseUid string, connectorId string, body ocpi.RawMessage[PatchedLocation]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) GetTariff(ctx context.Context, countryCode string, partyId string, sessionId string) (*Tariff, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PutTariff(ctx context.Context, countryCode string, partyId string, tariffId string, body ocpi.RawMessage[Tariff]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PatchTariff(ctx context.Context, countryCode string, partyId string, tariffId string, body ocpi.RawMessage[Tariff]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) GetSession(ctx context.Context, countryCode string, partyId string, sessionId string) (*Session, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PutSession(ctx context.Context, countryCode string, partyId string, sessionId string, body ocpi.RawMessage[Session]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PatchSession(ctx context.Context, countryCode string, partyId string, sessionId string, body ocpi.RawMessage[PatchedSession]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) GetToken(ctx context.Context, countryCode string, partyId string, tokenUid string, tokenType ...TokenType) (*Token, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PutToken(ctx context.Context, countryCode string, partyId string, tokenUid string, body ocpi.RawMessage[Token]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PatchToken(ctx context.Context, countryCode string, partyId string, tokenUid string, body ocpi.RawMessage[PatchedToken]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) GetCDR(ctx context.Context, cdrId string) (*ChargeDetailRecord, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PostCDR(ctx context.Context, body ocpi.RawMessage[ChargeDetailRecord]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) GetTerminal(ctx context.Context, terminalId string) (*Terminal, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) GetFinancialAdviceConfirmation(ctx context.Context, financialAdviceConfirmationId string) (*FinancialAdviceConfirmation, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PostTerminal(ctx context.Context, body ocpi.RawMessage[Terminal]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) PostFinancialAdviceConfirmation(ctx context.Context, body ocpi.RawMessage[FinancialAdviceConfirmation]) error {
	return ErrNotImplemented
}

func (UnimplementedServer) GetLocations(ctx context.Context, params GetOcpiLocationsParams) ([]Location, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) GetTariffs(ctx context.Context, params GetOcpiTariffsParams) ([]Tariff, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) GetSessions(ctx context.Context, params GetOcpiSessionsParams) ([]Session, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) GetTokens(ctx context.Context, params GetOcpiTokensParams) ([]Token, error) {
	return nil, ErrNotImplemented
}

func (UnimplementedServer) PostToken(ctx context.Context, tokenUid string, tokenType ...TokenType) (*AuthorizationInfo, error) {
	return nil, ErrNotImplemented
}
