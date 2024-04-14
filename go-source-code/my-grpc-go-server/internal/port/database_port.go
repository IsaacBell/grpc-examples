package port

import (
	"time"

	"github.com/google/uuid"
	db "github.com/timpamungkas/my-grpc-go-server/internal/adapter/database"
)

type DummyDatabasePort interface {
	Save(data *db.DummyOrm) (uuid.UUID, error)
	GetByUuid(uuid *uuid.UUID) (db.DummyOrm, error)
}

type BankDatabasePort interface {
	GetBankAccountByAccountNumber(acct string) (db.BankAccountOrm, error)
	CreateExchangeRate(r db.BankExchangeRateOrm) (uuid.UUID, error)
	GetExchangeRateAtTimestamp(fromCur string, toCur string, ts time.Time) (db.BankExchangeRateOrm, error)
	CreateTransaction(acct db.BankAccountOrm, t db.BankTransactionOrm) (uuid.UUID, error)
	CreateTransfer(transfer db.BankTransferOrm) (uuid.UUID, error)
	CreateTransferTransactionPair(fromAccountOrm db.BankAccountOrm, toAccountOrm db.BankAccountOrm,
		fromTransactionOrm db.BankTransactionOrm, toTransactionOrm db.BankTransactionOrm) (bool, error)
	UpdateTransferStatus(transfer db.BankTransferOrm, status bool) error
}
