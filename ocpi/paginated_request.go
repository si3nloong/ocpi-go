package ocpi

type PaginatedRequest[DT any] struct {
	// DateFrom Return tokens that have last_updated after or equal to this Date/Time (inclusive).
	DateFrom *DT `form:"date_from,omitempty" json:"date_from,omitempty"`

	// DateTo Return tokens that have last_updated up to Date/Time, but not including (exclusive).
	DateTo *DT `form:"date_to,omitempty" json:"date_to,omitempty"`

	// Offset The offset of the first object returned. Default is 0.
	Offset *int `form:"offset,omitempty" json:"offset,omitempty"`

	// Limit Maximum number of objects to GET.
	Limit *int `form:"limit,omitempty" json:"limit,omitempty"`
}
