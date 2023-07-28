package models

type DomainResponse struct {
	Name       string `json:"domainName"`
	IsCatchall string `json:"isCatchall"`
	Events     int64  `json:"events"`
	Bounced    bool   `json:"bounced"`
}
