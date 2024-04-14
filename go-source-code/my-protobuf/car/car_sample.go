package car

import (
	"log"
	"my-protobuf/protogen/car"

	"github.com/google/uuid"
)

func ValidateCar() {
	car := car.Car{
		CarId: uuid.New().String(),
		Brand: "Bmw",
		Model: "745i Sport",
		Price: 204000,
	}

	err := car.ValidateAll()

	if err != nil {
		log.Fatalln("Validation failed", err)
	}

	log.Println(&car)
}
