package basic

import (
	"log"
	"math/rand"
	"my-protobuf/protogen/basic"

	"google.golang.org/genproto/googleapis/type/date"
	"google.golang.org/genproto/googleapis/type/latlng"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func BasicUser() {
	addr := basic.Address{
		Street:  "Daily Planet",
		City:    "Metropolis",
		Country: "US",
		Coordinate: &basic.Address_Coordinate{
			Latitude:  40.70797893425118,
			Longitude: -74.01163838107261,
		},
	}

	comm := randomCommunicationChannel()
	sr := map[string]uint32{"fly": 5, "speed": 5, "durability": 4}

	u := basic.User{
		Id:                   99,
		Username:             "superman",
		IsActive:             true,
		Password:             []byte("supermanpassword"),
		Address:              &addr,
		CommunicationChannel: &comm,
		SkillRating:          sr,
		LastLoginTimestamp:   timestamppb.Now(),
		BirthDate:            &date.Date{Year: 2000, Month: 5, Day: 27},
		LastKnownLocation: &latlng.LatLng{
			Latitude:  -6.29847717,
			Longitude: 106.8290577,
		},
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}

func ProtoToJsonUser() {
	p := basic.User{
		Id:       98,
		Username: "wonderwoman",
		IsActive: true,
		Password: []byte("wonderwomanpassword"),
		Emails:   []string{"wonderwoman@movie.com", "wonderwoman@dc.com"},
		Gender:   basic.Gender_GENDER_FEMALE,
	}

	jsonBytes, _ := protojson.Marshal(&p)
	log.Println(string(jsonBytes))
}

func JsonToProtoUser() {
	json := `
	{
			"id":97,
			"username":"batman",
			"is_active":true,
			"password":"YmF0bWFucGFzc3dvcmQ=",
			"emails":[
				 "batman@movie.com",
				 "batman@dc.com"
			],
			"gender":"GENDER_MALE"
	 }
	`

	var p basic.User

	err := protojson.Unmarshal([]byte(json), &p)

	if err != nil {
		log.Fatalln("Got error : ", err)
	}

	log.Println(&p)
}

func randomCommunicationChannel() anypb.Any {
	paperMail := basic.PaperMail{
		PaperMailAddress: "Some paper mail address",
	}

	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "byteMe",
		SocialMediaUsername: "krypton.man",
	}

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct:  "whatsup",
		InstantMessagingUsername: "krypton.last",
	}

	var a anypb.Any

	switch r := rand.Intn(10) % 3; r {
	case 0:
		anypb.MarshalFrom(&a, &paperMail, proto.MarshalOptions{})
	case 1:
		anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})
	default:
		anypb.MarshalFrom(&a, &instantMessaging, proto.MarshalOptions{})
	}

	return a
}

func BasicUnmarshalAnyKnown() {
	socialMedia := basic.SocialMedia{
		SocialMediaPlatform: "my-social-media-platform",
		SocialMediaUsername: "my-social-media-username",
	}

	var a anypb.Any
	anypb.MarshalFrom(&a, &socialMedia, proto.MarshalOptions{})

	// known type (Social Media)
	sm := basic.SocialMedia{}

	if err := a.UnmarshalTo(&sm); err != nil {
		return
	}

	jsonBytes, _ := protojson.Marshal(&sm)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyNotKnown() {
	a := randomCommunicationChannel()

	var unmarshaled protoreflect.ProtoMessage

	unmarshaled, err := a.UnmarshalNew()

	if err != nil {
		return
	}

	log.Println("Unmarshal as", unmarshaled.ProtoReflect().Descriptor().Name())

	jsonBytes, _ := protojson.Marshal(unmarshaled)
	log.Println(string(jsonBytes))
}

func BasicUnmarshalAnyIs() {
	a := randomCommunicationChannel()
	pm := basic.PaperMail{}

	if a.MessageIs(&pm) {
		if err := a.UnmarshalTo(&pm); err != nil {
			log.Fatalln(err)
		}

		jsonBytes, _ := protojson.Marshal(&pm)
		log.Println(string(jsonBytes))
	} else {
		log.Println("Not PaperMail, but :", a.TypeUrl)
	}
}

func BasicOneof() {
	// socialMedia := basic.SocialMedia{
	// 	SocialMediaPlatform: "byteMe",
	// 	SocialMediaUsername: "aquaman",
	// }

	// ecomm := basic.User_SocialMedia{
	// 	SocialMedia: &socialMedia,
	// }

	instantMessaging := basic.InstantMessaging{
		InstantMessagingProduct:  "whatsup",
		InstantMessagingUsername: "atlantis.king",
	}

	ecomm := basic.User_InstantMessaging{
		InstantMessaging: &instantMessaging,
	}

	u := basic.User{
		Id:                    96,
		Username:              "aquaman",
		IsActive:              true,
		ElectronicCommChannel: &ecomm,
	}

	jsonBytes, _ := protojson.Marshal(&u)
	log.Println(string(jsonBytes))
}
