package bank

import (
	"errors"
	"time"
)

const (
	TransactionTypeUnknown string = "UNKNOWN"
	TransactionTypeIn      string = "IN"
	TransactionTypeOut     string = "OUT"
)

type ExchangeRate struct {
	FromCurrency       string
	ToCurrency         string
	Rate               float64
	ValidFromTimestamp time.Time
	ValidToTimestamp   time.Time
}

type Transaction struct {
	Amount          float64
	Timestamp       time.Time
	TransactionType string
	Notes           string
}

type TransactionSummary struct {
	SummaryOnDate time.Time
	SumIn         float64
	SumOut        float64
	SumTotal      float64
}

type TransferTransaction struct {
	FromAccountNumber string
	ToAccountNumber   string
	Currency          string
	Amount            float64
}

var ErrTransferSourceAccountNotFound = errors.New("source account not found")
var ErrTransferDestinationAccountNotFound = errors.New("destination account not found")
var ErrTransferRecordFailed = errors.New("can't create transfer record")
var ErrTransferTransactionPair = errors.New("can't create transfer transaction pair, " +
	"possibly insufficient balance on source account")
