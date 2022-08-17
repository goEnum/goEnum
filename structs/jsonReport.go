package structs

type JSONReport struct {
	Vulnerability string   `json:"vulnerbility"`
	Locations     []string `json:"locations"`
	Description   string   `json:"description"`
}

func NewJSONReport(vulnerability string, locations []string, description string) *JSONReport {
	return &JSONReport{
		Vulnerability: vulnerability,
		Locations:     locations,
		Description:   description,
	}
}
