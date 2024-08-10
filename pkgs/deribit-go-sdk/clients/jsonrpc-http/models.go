package jsonrpchttp

type MessageRequest struct {
	JsonRpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Method  string `json:"method"`
	Params  any    `json:"params"`
}

type MessageResponse struct {
	Id      int          `json:"id"`
	Result  any          `json:"result"`
	Error   ErrorMessage `json:"error"`
	TestNet bool         `json:"testnet"`
	UsIn    int          `json:"usIn"`
	UsOut   int          `json:"usOut"`
	UsDiff  int          `json:"usDiff"`
}

type ErrorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
