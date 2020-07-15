package main

// Generate protobuf files:
//go:generate protoc -I. -I$GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis --go_out=plugins=grpc:. --grpc-gateway_out=logtostderr=true,paths=source_relative:. --swagger_out=logtostderr=true:. protos/texture_studio.proto

import (
	"flag"
	"github.com/Mrucznik/texture-studio-api/protos"
	"net/http"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "localhost:9090", "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := protos.RegisterTextureStudioServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(":8080", mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}