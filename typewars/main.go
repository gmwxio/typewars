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
	var script *pb.Script
	for idx, fi := range fds.File {
		fmt.Printf("%d proto '%s'\n", idx, fi.GetName())
		if fi.GetName() == "anewhope.proto" {
			ext, err := proto.GetExtension(fds.File[idx].Options, pb.E_File)
			if err != nil {
				glog.Fatalf("%v\n", err)
			}
			script = ext.(*pb.Script)
		}
	}
	fmt.Printf("Props\n")
	for i, prop := range script.Props {
		fmt.Printf("Prop %d\n", i)
		xwings := prop.GetXwings()
		fmt.Printf("  XWings\n")
		for i, xwing := range xwings.Xwingies {
			fmt.Printf("  %d\t%+v\n", i, xwing)
		}
	}
}
