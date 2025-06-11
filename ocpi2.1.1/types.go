package ocpi211

// DisplayText defines model for cdrBody_tariffs_tariff_alt_text.
type DisplayText struct {
	Language string `json:"language" validate:"required,len=2"`
	Text     string `json:"text" validate:"required,max=512"`
}
