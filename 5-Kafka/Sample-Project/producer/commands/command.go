package commands

type OpenAccountCommand struct {
	AccountHolder string
	AccountType   int
	Balance       float64
}

type DepositFundCommand struct {
	ID     string
	Amount float64
}

type WithdrawFundCommand struct {
	ID     string
	Amount float64
}

type CloseAccountCommand struct {
	ID string
}

type ShowBalanceCommand struct {
	ID string
}

type ShowTransactionsCommand struct {
	ID string
}
