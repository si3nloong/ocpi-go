package ocpi211

import "context"

type CPO interface {
	GetLocations(ctx context.Context) ([]Location, error)
}
