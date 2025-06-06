package ocpi

//go:generate stringer -type StatusCode
type StatusCode uint16

const (
	GenericSuccessCode StatusCode = 1000 // Success
	GenericClientError StatusCode = 2000 // Client Error
	GenericServerError StatusCode = 3000 // Server Error
)
