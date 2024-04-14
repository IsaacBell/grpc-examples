package basic

import (
	"log"
	"my-protobuf/protogen/basic"
)

func BasicHello() {
	h := basic.Hello{
		Name: "Clark Kent",
	}

	log.Println(&h)
}
