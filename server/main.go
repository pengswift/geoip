package main

import (
	"errors"
	"log"
	"net"
	"os"
	"strings"

	"github.com/oschwald/geoip2-golang"
	"github.com/oschwald/maxminddb-golang"
	pb "github.com/pengswift/geoip/geoip"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	port = ":60052"
)

var (
	ERROR_CANNOT_QUERY_IP = errors.New("cannot query ip")
)

type Reader struct {
	mmdbReader *maxminddb.Reader
}

type server struct {
	db *geoip2.Reader
}

func (s *server) init() {
	dbpath := s.dataPath()
	if dbpath == "" {
		log.Println("geoip dbpath does not exist")
		return
	}

	log.Println("Loading GEOIP City...")
	db, err := geoip2.Open(dbpath)
	if err != nil {
		log.Panic(err)
	}

	s.db = db
	log.Println("GEOIP City Load Complete")
}

func (s *server) dataPath() (path string) {
	paths := strings.Split(os.Getenv("GOPATH"), ":")
	for k := range paths {
		path = paths[k] + "/data/GeoLite2-City.mmdb"
		//path = paths[k] + "/data/GeoIP2-City.mmdb"
		_, err := os.Lstat(path)
		if err == nil {
			return
		}
	}
	return
}

func (s *server) QueryCountry(ctx context.Context, in *pb.GeoIPRequest) (*pb.GeoIPResponse, error) {
	ip := net.ParseIP(in.Ip)
	if record, err := s.db.City(ip); err == nil {
		return &pb.GeoIPResponse{Name: record.Country.IsoCode}, nil
	}
	return nil, ERROR_CANNOT_QUERY_IP
}

func (s *server) QuerySubdivision(ctx context.Context, in *pb.GeoIPRequest) (*pb.GeoIPResponse, error) {
	ip := net.ParseIP(in.Ip)
	if record, err := s.db.City(ip); err == nil {
		return &pb.GeoIPResponse{Name: record.Subdivisions[0].Names["en"]}, nil
	}
	return nil, ERROR_CANNOT_QUERY_IP
}

func (s *server) QueryCity(ctx context.Context, in *pb.GeoIPRequest) (*pb.GeoIPResponse, error) {
	ip := net.ParseIP(in.Ip)
	if record, err := s.db.City(ip); err == nil {
		return &pb.GeoIPResponse{Name: record.Country.Names["en"]}, nil
	}
	return nil, ERROR_CANNOT_QUERY_IP
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	ins := &server{}
	ins.init()
	pb.RegisterGeoIPServiceServer(s, ins)
	s.Serve(lis)
}
