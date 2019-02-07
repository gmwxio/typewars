package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"typewars/pb"

	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func main() {
	flag.Parse() // needed to make glog happy
	fds := &descriptor.FileDescriptorSet{}
	bytes, _ := ioutil.ReadAll(os.Stdin)
	err := proto.Unmarshal(bytes, fds)
	if err != nil {
		glog.Fatalf("%v\n", err)
	}
	ext, err := proto.GetExtension(fds.File[1].Options, pb.E_File)
	if err != nil {
		glog.Fatalf("%v\n", err)
	}
	script := ext.(*pb.Script)
	fmt.Printf("XWings\n")
	for i, xwing := range script.GetXwing() {
		fmt.Printf("%d\t%+v\n", i, xwing)
	}
}
