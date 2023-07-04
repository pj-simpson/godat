package models

type Company struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Platform        string           `json:"platform"`
	Redirect        string           `json:"redirect"`
	LastSync        string           `json:"lastSync"`
	DataConnections []DataConnection `json:"dataConnections"`
	Created         string           `json:"created,omitempty"`
}

type DataConnection struct {
	ID             string `json:"id"`
	IntegrationID  string `json:"integrationId"`
	IntegrationKey string `json:"integrationKey"`
	SourceID       string `json:"sourceId"`
	PlatformName   string `json:"platformName"`
	LinkURL        string `json:"linkUrl"`
	Status         string `json:"status"`
	LastSync       string `json:"lastSync"`
	Created        string `json:"created"`
	SourceType     string `json:"sourceType"`
}
