package request

type ChatRequest struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
}

type OneHistoryRequest struct {
	Id int `json:"id"`
}

type ChangeRequest struct {
	Message string `json:"message"`
}
