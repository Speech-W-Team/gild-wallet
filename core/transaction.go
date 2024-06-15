package core

// Transaction represents a blockchain transaction
type Transaction struct {
	From   string
	To     string
	Amount float64
	Fee    float64
	Nonce  uint64
	Data   []byte
}

// TransactionManager interface for transaction management
type TransactionManager interface {
	CreateTransaction(from string, to string, amount float64, fee float64, nonce uint64, data []byte) (Transaction, error)
	SignTransaction(tx Transaction, privateKey string) (string, error)
	BroadcastTransaction(signedTx string) (string, error)
}
