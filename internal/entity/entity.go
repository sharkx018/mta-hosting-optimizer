package entity

type InefficentMTAResponseParams struct {
	Data  []string             `json:"data,omitempty"`
	Error *CommonErrorResponse `json:"errors,omitempty"`
}

type CommonErrorResponse struct {
	Message string `json:"message"`
}

type MockServiceResponse struct {
	Data []MtaData
}

type MtaData struct {
	IP       string
	Hostname string
	Active   bool
}
