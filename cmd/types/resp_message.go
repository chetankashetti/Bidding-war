package types

type RespMessage struct {
	Message  string `json:"message"`
	GasSpent uint64 `json:"gasSpent"`
}
