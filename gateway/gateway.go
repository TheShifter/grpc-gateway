package main

import (
	gw "../proto"
	"../server"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net/http"
)

func main() {
	proxyAddr := ":8080"
	serviceAddr := "127.0.0.1:8081"
	go server.GRPCService(serviceAddr)
	HTTPProxy(proxyAddr, serviceAddr)
}

func HTTPProxy(proxyAddr string, serviceAddr string) {
	grpcConn, err := grpc.Dial(serviceAddr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to connect to grpc ", err)
	}
	defer grpcConn.Close()
	grpcGWMux := runtime.NewServeMux()
	err = gw.RegisterCountriesHandler(context.Background(), grpcGWMux, grpcConn)
	if err != nil {
		log.Fatalln("failed to start HTTP server", err)
	}
	mux := http.NewServeMux()
	mux.Handle("/api/v1/countries", grpcGWMux)
	log.Fatalln(http.ListenAndServe(proxyAddr, mux))
}
