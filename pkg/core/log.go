package core

type Log struct {
	BlockNumber int64
	TxHash      string
	Address     string
	Topics      []string
	Data        string
}
