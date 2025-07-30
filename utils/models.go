package utils

// Representation of an operator channel
type OperatorChannel struct {
	CsvName          string   `json:"csvName"`
	DisplayName      string   `json:"displayName"`
	AdditionalImages []string `json:"additionalImages"`
}

// Representation of an operator package
type OperatorPackage struct {
	PackageName        string                      `json:"packageName"`
	DefaultChannel     string                      `json:"defaultChannel"`
	DefaultDisplayName string                      `json:"defaultDisplayName"`
	Channels           map[string]*OperatorChannel `json:"channels"`
}
