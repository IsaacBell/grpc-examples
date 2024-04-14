package main

import (
	"database/sql"
	"log"
	"math/rand"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	dbmigration "github.com/timpamungkas/my-grpc-go-server/db"

	mydb "github.com/timpamungkas/my-grpc-go-server/internal/adapter/database"
	mygrpc "github.com/timpamungkas/my-grpc-go-server/internal/adapter/grpc"
	app "github.com/timpamungkas/my-grpc-go-server/internal/application"
	"github.com/timpamungkas/my-grpc-go-server/internal/application/domain/bank"
)

func main() {
	log.SetFlags(0)
	log.SetOutput(logWriter{})

	sqlDB, err := sql.Open("pgx", "postgres://postgres:postgres@localhost:5432/grpc?sslmode=disable")

	if err != nil {
		log.Fatalln("Can't connect database :", err)
	}

	dbmigration.Migrate(sqlDB)

	databaseAdapter, err := mydb.NewDatabaseAdapter(sqlDB)

	if err != nil {
		log.Fatalln("Can't create database adapter :", err)
	}

	hs := &app.HelloService{}
	bs := app.NewBankService(databaseAdapter)
	rs := &app.ResiliencyService{}

	go generateExchangeRates(bs, "USD", "IDR", 5*time.Second)

	grpcAdapter := mygrpc.NewGrpcAdapter(hs, bs, rs, 9090)

	grpcAdapter.Run()
}

// func runDummyOrm(da *mydb.DatabaseAdapter) {
// 	now := time.Now()

// 	uuid, _ := da.Save(
// 		&mydb.DummyOrm{
// 			UserId:    uuid.New(),
// 			UserName:  "Tim " + time.Now().Format("15:04:05"),
// 			CreatedAt: now,
// 			UpdatedAt: now,
// 		},
// 	)

// 	res, _ := da.GetByUuid(&uuid)

// 	log.Println("res :", res)
// }

func generateExchangeRates(bs *app.BankService, fromCurrency, toCurrency string, duration time.Duration) {
	ticker := time.NewTicker(duration)

	for range ticker.C {
		now := time.Now()
		validFrom := now.Truncate(time.Second).Add(3 * time.Second)
		validTo := validFrom.Add(duration).Add(-1 * time.Millisecond)

		dummyRate := bank.ExchangeRate{
			FromCurrency:       fromCurrency,
			ToCurrency:         toCurrency,
			ValidFromTimestamp: validFrom,
			ValidToTimestamp:   validTo,
			Rate:               2000 + float64(rand.Intn(300)),
		}

		bs.CreateExchangeRate(dummyRate)
	}
}
