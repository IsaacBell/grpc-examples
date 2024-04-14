package basic

import (
	"io/ioutil"
	"log"
	"my-protobuf/protogen/basic"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func WriteProtoToJSON(msg proto.Message, fname string) {
	b, err := protojson.Marshal(msg)

	if err != nil {
		log.Fatalln("Can not marshal message", err)
	}

	if err := ioutil.WriteFile(fname, b, 0644); err != nil {
		log.Fatalln("Can not write to file", err)
	}
}

func ReadProtoFromJSON(fname string, dest proto.Message) {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Can not read file", err)
	}

	if err := protojson.Unmarshal(in, dest); err != nil {
		log.Fatalln("Can not unmarshal", err)
	}
}

func WriteToJSONSample() {
	u := dummyUser()
	WriteProtoToJSON(&u, "superman_file.json")
}

func ReadFromJSONSample() {
	dest := basic.User{}

	ReadProtoFromJSON("superman_file.json", &dest)

	log.Println(&dest)
}
