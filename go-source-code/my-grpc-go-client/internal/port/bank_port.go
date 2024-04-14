package port

import (
	"context"

	"github.com/timpamungkas/my-grpc-proto/protogen/go/bank"
	"google.golang.org/grpc"
)

type BankClientPort interface {
	GetCurrentBalance(ctx context.Context, in *bank.CurrentBalanceRequest,
		opts ...grpc.CallOption) (*bank.CurrentBalanceResponse, error)
	FetchExchangeRates(ctx context.Context, in *bank.ExchangeRateRequest,
		opts ...grpc.CallOption) (bank.BankService_FetchExchangeRatesClient, error)
	SummarizeTransactions(ctx context.Context,
		opts ...grpc.CallOption) (bank.BankService_SummarizeTransactionsClient, error)
	TransferMultiple(ctx context.Context,
		opts ...grpc.CallOption) (bank.BankService_TransferMultipleClient, error)
	CreateAccount(ctx context.Context, in *bank.CreateAccountRequest,
		opts ...grpc.CallOption) (*bank.CreateAccountResponse, error)
}
