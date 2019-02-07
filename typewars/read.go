package main

import (
	"encoding/json"
	"fmt"
	"os"
	"typewars/pb"

	"github.com/golang/protobuf/proto"
)

func main() {
	redWing := pb.Xwing{}
	bytes := make([]byte, 1024)
	os.Stdin.Read(bytes)
	proto.Unmarshal(bytes, &redWing)
	jbytes, _ := json.Marshal(&redWing)
	fmt.Printf("%s\n", string(jbytes))
}
