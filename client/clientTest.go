package main

import (
	pb "../proto"
	"flag"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"os"
)

func main() {
	flag.Parse()
	if flag.NArg() < 1 {
		log.Fatalln(os.Stderr, "no args")
		os.Exit(1)
	}
	conn, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("could not connect to backend: ", err)
	}
	client := pb.NewCountriesClient(conn)
	operationChoice(client)
	if err != nil {
		log.Fatalln(os.Stderr, err)
		os.Exit(1)
	}
}
func operationChoice(client pb.CountriesClient) {
	var err error
	switch cmd := flag.Arg(0); cmd {
	case "create":
		country := pb.Country{2, "Belarus", 95000000}
		createRequest := &pb.CreateRequest{"golang-grpc", &country}
		createResponse, err := client.Create(context.Background(), createRequest)
		if err != nil {
			log.Println("response has been failed")
		}
		log.Printf("Country id: %d", createResponse.Id)
	case "read":
		readRequest := &pb.ReadRequest{"golang-grpc", 4}
		readResponse, err := client.Read(context.Background(), readRequest)
		if err != nil {
			log.Println("response has been failed")
		}
		log.Println(readResponse.Country)
	case "update":
		country := pb.Country{4, "Belarus", 95000000}
		var id int64 = 1
		updateRequest := &pb.UpdateRequest{"golang-grpc", &country, id}
		updateResponse, err := client.Update(context.Background(), updateRequest)
		if err != nil {
			log.Println("response has been failed")
		}
		log.Println(updateResponse.Updated)
	case "readAll":
		readAllRequest := &pb.ReadAllRequest{"golang-grpc"}
		readAllResponse, err := client.ReadAll(context.Background(), readAllRequest)
		if err != nil {
			log.Println("response has been failed")
		}
		log.Println(readAllResponse.Country)
	default:
		if err != nil {
			log.Fatalln("unknown sub command:", cmd)
		}
	}
}
