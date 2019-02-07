package main

import (
	"os"
	"typewars/pb"

	"github.com/golang/protobuf/proto"
)

func main() {
	redWing := pb.Xwing{
		Pilot:      "Luke",
		AstroDroid: "R2D2",
	}
	bytes, _ := proto.Marshal(&redWing)
	os.Stdout.Write(bytes)
}
