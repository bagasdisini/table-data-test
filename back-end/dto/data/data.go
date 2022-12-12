package datadto

type DataRequest struct {
	ProductName  string `json:"productName"`
	Amount       int    `json:"amount"`
	CustomerName string `json:"customerName"`
	StatusID     int    `json:"statusID"`
	CreateBy     string `json:"createBy"`
}

type DataResponse struct {
	ID           int    `json:"id"`
	ProductID    int64  `json:"productID"`
	ProductName  string `json:"productName"`
	Amount       int    `json:"amount"`
	CustomerName string `json:"customerName"`
	StatusID     int    `json:"statusID"`
	CreateBy     string `json:"createBy"`
}
