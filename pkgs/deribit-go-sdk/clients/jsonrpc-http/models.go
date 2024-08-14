package jsonrpchttp

type MessageRequest[Params any] struct {
	JsonRpc string `json:"jsonrpc"`
	Id      int    `json:"id"`
	Method  string `json:"method"`
	Params  Params `json:"params"`
}

type MessageResponse[Result any] struct {
	Id      int          `json:"id"`
	Result  Result       `json:"result"`
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
