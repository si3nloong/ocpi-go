package ocpi

//go:generate stringer -type=StatusCode -linecomment
type StatusCode uint16

const (
	// StatusCode is a numeric code that indicates the status of the request.
	StatusCodeSuccess StatusCode = 1000 // Success

	// Client errors: Errors detected by a server in the message sent by a client: The client did something wrong
	StatusCodeClientError                           StatusCode = 2000 // Client Error
	StatusCodeClientErrorInvalidOrMissingParameters StatusCode = 2001 // Invalid or missing parameters
	StatusCodeClientErrorNotEnoughInfo              StatusCode = 2002 // Not enough information, for example: Authorization request with too little information.
	StatusCodeClientErrorUnknownLocation            StatusCode = 2003 // Unknown Location, for example: Command: START_SESSION with unknown location.

	// Server errors: Error during processing of the OCPI payload in the server. The message was syntactically correct but could not be processed by the server.
	StatusCodeServerError                     StatusCode = 3000 // Generic Server Error
	StatusCodeServerErrorUnableToUseClientAPI StatusCode = 3001 // Unable to use the client's API. For example during the credentials registration: When the initializing party requests data from the other party during the open POST call to its credentials endpoint. If one of the GETs can not be processed, the party should return this error in the POST response.
	StatusCodeServerErrorUnsupportedVersion   StatusCode = 3002 // Unsupported version
	StatusCodeServerErrorNoMatchingEndpoints  StatusCode = 3003 // No matching endpoints or expected endpoints missing between parties. Used during the registration process if the two parties do not have any mutual modules or endpoints available, or the minimum expected by the other party implementation.
)
