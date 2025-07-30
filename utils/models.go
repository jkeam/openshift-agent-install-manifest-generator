package utils

type ChannelResponse struct {
	CsvName          string   `json:"csvName"`
	DisplayName      string   `json:"displayName"`
	AdditionalImages []string `json:"additionalImages"`
}

type GetPackageResponse struct {
	PackageName        string                      `json:"packageName"`
	DefaultChannel     string                      `json:"defaultChannel"`
	DefaultDisplayName string                      `json:"defaultDisplayName"`
	Channels           map[string]*ChannelResponse `json:"channels"`
}
