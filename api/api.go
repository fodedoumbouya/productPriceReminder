package api

/// we put all the api response and request here [struct]

type APIResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
