package core

// BalanceChecker interface for checking wallet balances
type BalanceChecker interface {
	GetBalance(address string) (float64, error)
}
