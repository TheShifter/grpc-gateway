package server

import (
	"../database/config"
	"../database/model"
	pb "../proto"
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type server struct {
}

func (s *server) Create(ctx context.Context, in *pb.CreateRequest) (*pb.CreateResponse, error) {
	db, err := config.GetMySQLDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	countryModel := model.Database{Db: db}
	err = countryModel.Create(in.Country)
	if err != nil {
		log.Fatalln("country hasn't been added")
	}
	return &pb.CreateResponse{"golang-grpc", in.Country.Id}, nil
}

func (s *server) Read(ctx context.Context, in *pb.ReadRequest) (*pb.ReadResponse, error) {
	db, err := config.GetMySQLDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	countryModel := model.Database{Db: db}
	country, err := countryModel.Read(in.Id)
	if err != nil {
		log.Fatalln("country hasn't been read")
	}
	return &pb.ReadResponse{"golang-grpc", &country}, nil
}
func (s *server) Update(ctx context.Context, in *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	db, err := config.GetMySQLDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	countryModel := model.Database{Db: db}
	err = countryModel.Update(in.Country)
	if err != nil {
		log.Fatalln("country hasn't been updated")
	}
	return &pb.UpdateResponse{"golang-grpc", in.Country.Id}, nil
}
func (s *server) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	db, err := config.GetMySQLDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	countryModel := model.Database{Db: db}
	err = countryModel.Delete(in.Id)
	if err != nil {
		log.Fatalln("country hasn't been deleted")
	}
	return &pb.DeleteResponse{"golang-grpc", in.Id}, nil
}
func (s *server) ReadAll(ctx context.Context, in *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	db, err := config.GetMySQLDB()
	defer db.Close()
	if err != nil {
		fmt.Println(err)
	}
	countryModel := model.Database{Db: db}
	countries, err := countryModel.ReadAll()
	if err != nil {
		log.Fatalln("countries haven't been received")
	}
	return &pb.ReadAllResponse{"golang-grpc", countries}, nil
}

func GRPCService(serviceAddr string) {
	listener, err := net.Listen("tcp", serviceAddr)
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	pb.RegisterCountriesServer(srv, &server{})
	reflection.Register(srv)
	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
