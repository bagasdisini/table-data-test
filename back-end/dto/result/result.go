package dto

type SuccessResult struct {
	Data   interface{} `json:"data"`
	Status interface{} `json:"status"`
}

type ErrorResult struct {
	Message    string `json:"message"`
	StatusCode int    `json:"statusCode"`
}
