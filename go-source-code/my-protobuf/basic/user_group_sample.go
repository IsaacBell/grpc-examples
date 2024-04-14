package basic

import (
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
)

func BasicUserGroup() {
	batman := basic.User{
		Id:       97,
		Username: "batman",
		IsActive: true,
		Password: []byte("batmanpassword"),
		Gender:   basic.Gender_GENDER_MALE,
	}

	nightwing := basic.User{
		Id:       96,
		Username: "nightwing",
		IsActive: true,
		Password: []byte("nightwingpassword"),
		Gender:   basic.Gender_GENDER_MALE,
	}

	robin := basic.User{
		Id:       95,
		Username: "robin",
		IsActive: true,
		Password: []byte("robinpassword"),
		Gender:   basic.Gender_GENDER_MALE,
	}

	batFamily := basic.UserGroup{
		GroupId:     999,
		GroupName:   "Bat Family",
		Users:       []*basic.User{&batman, &nightwing, &robin},
		Description: "The classic bat family",
	}

	jsonBytes, _ := protojson.Marshal(&batFamily)
	log.Println(string(jsonBytes))
}
