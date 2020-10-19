package main

import (
	"fmt"
	"github.com/bigwhite/protobufbench_goproto/submit"
	"github.com/golang/protobuf/proto"
)

func main() {
	fmt.Println(111)

	var request = submit.Request{
		Recvtime: 170123456,
		Uniqueid: "a1b2c3d4e5f6g7h8i9",
		Token:    "xxxx-1111-yyyy-2222-zzzz-3333",
		Phone:    "13900010002",
		Content:  "Customizing the fields of the messages to be the fields that you actually want to use removes the need to copy between the structs you use and structs you use to serialize. gogoprotobuf also offers more serialization formats and generation of tests and even more methods.",
		Sign:     "tonybaiXZYDFDS",
		Type:     "submit",
		Extend:   "",
		Version:  "v1.0.0",
	}
	_, _ = proto.Marshal(&request)

}
