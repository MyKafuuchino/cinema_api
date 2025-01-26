package types

type QueryParamRequest struct {
	Status string `json:"status,omitempty"`
	Limit  int    `json:"limit"`
	Offset int    `json:"offset"`
}
