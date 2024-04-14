package port

import (
	"time"

	"github.com/google/uuid"
	dbank "github.com/timpamungkas/my-grpc-go-server/internal/application/domain/bank"
)

type HelloServicePort interface {
	GenerateHello(name string) string
}

type BankServicePort interface {
	FindCurrentBalance(acct string) (float64, error)
	CreateExchangeRate(r dbank.ExchangeRate) (uuid.UUID, error)
	FindExchangeRate(fromCur string, toCur string, ts time.Time) (float64, error)
	CreateTransaction(acct string, t dbank.Transaction) (uuid.UUID, error)
	CalculateTransactionSummary(tcur *dbank.TransactionSummary, trans dbank.Transaction) error
	Transfer(tt dbank.TransferTransaction) (uuid.UUID, bool, error)
}

type ResiliencyServicePort interface {
	GenerateResiliency(minDelaySecond int32, maxDelaySecond int32, statusCodes []uint32) (string, uint32)
}
