package ocpi

// VersionsData defines model for versions_data.
type Version struct {
	URL     string        `json:"url"`
	Version VersionNumber `json:"version"`
}
