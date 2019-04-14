package server

import (
	"../dao/entities"
	. "../dao/implementations"
	"../dao/interfaces"
	pb "../proto"
	"errors"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
)

const api = "v1"

type server struct {
}

func checkAPI(reqAPI string) error {
	if len(reqAPI) > 0 {
		if reqAPI != api {
			return status.Errorf(codes.Unimplemented,
				"unsupported API version: service implements"+
					" API version '%s', but asked for '%s'", api, reqAPI)
		}
	} else {
		return errors.New("API version is empty")
	}
	return nil
}
func (s *server) Read(ctx context.Context, req *pb.ReadRequest) (*pb.ReadResponse, error) {
	var country *entities.Country
	var dao interfaces.CountryDAO
	dao = CountryDAOImpl{}
	country, err := dao.Read(req.Id)
	if err != nil {
		log.Fatalln("country hasn't been read")
	}
	respCountry := pb.Country{
		Id:           country.Id,
		Name:         country.Name,
		PeopleNumber: country.PeopleNumber,
	}
	return &pb.ReadResponse{Api: api, Country: &respCountry}, nil
}
func (s *server) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	if err := checkAPI(req.GetApi()); err != nil {
		return nil, err
	}
	country := entities.Country{
		Name:         req.Country.Name,
		PeopleNumber: req.Country.PeopleNumber,
	}
	var dao interfaces.CountryDAO
	dao = CountryDAOImpl{}
	id, err := dao.Create(&country)
	if err != nil {
		log.Fatalln("country hasn't been added")
	}
	return &pb.CreateResponse{Api: api, Id: id}, nil
}

func (s *server) Update(ctx context.Context, req *pb.UpdateRequest) (*pb.UpdateResponse, error) {
	if err := checkAPI(req.Api); err != nil {
		return nil, err
	}
	country := entities.Country{
		Id:           req.Id,
		Name:         req.Country.Name,
		PeopleNumber: req.Country.PeopleNumber,
	}
	var dao interfaces.CountryDAO
	dao = CountryDAOImpl{}
	err := dao.Update(&country)
	if err != nil {
		log.Fatalln("country hasn't been updated")
	}
	return &pb.UpdateResponse{Api: api, Updated: req.Country.Id}, nil
}
func (s *server) Delete(ctx context.Context, req *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	var dao interfaces.CountryDAO
	dao = CountryDAOImpl{}
	err := dao.Delete(req.Id)
	if err != nil {
		log.Fatalln("country hasn't been read")
	}
	return &pb.DeleteResponse{Api: api, Deleted: req.Id}, nil
}
func (s *server) ReadAll(ctx context.Context, req *pb.ReadAllRequest) (*pb.ReadAllResponse, error) {
	var dao interfaces.CountryDAO
	dao = CountryDAOImpl{}
	var countries []*entities.Country
	countries, err := dao.ReadAll()
	if err != nil {
		log.Fatalln("countries haven't been read")
	}
	var respCountries []*pb.Country
	for i := 0; i < len(countries); i++ {
		country := &pb.Country{
			Name:         countries[i].Name,
			PeopleNumber: countries[i].PeopleNumber,
		}
		respCountries = append(respCountries, country)
	}

	return &pb.ReadAllResponse{Api: api, Country: respCountries}, nil
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
