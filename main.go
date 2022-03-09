package main

import (
	"fmt"
	"github.com/tikv/client-go/v2/tikv"
	"google.golang.org/grpc"
)

func main() {
	//Suppose we have to use tikv v2.0.0-rc version
	fmt.Println(tikv.DCLabelKey)

	//Suppose we have to use grpc v1.44.0 version
	fmt.Println(grpc.Version)

	// exec `go mod tidy`  will occur error
	//	go: finding module for package google.golang.org/grpc/naming
	//  grpc-test imports
	//        github.com/tikv/client-go/v2/tikv imports
	//        go.etcd.io/etcd/clientv3 tested by
	//        go.etcd.io/etcd/clientv3.test imports
	//        go.etcd.io/etcd/integration imports
	//        go.etcd.io/etcd/proxy/grpcproxy imports
	//        google.golang.org/grpc/naming: module google.golang.org/grpc@latest found (v1.44.0), but does not contain package google.golang.org/grpc/naming

}
