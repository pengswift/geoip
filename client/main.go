package main

import (
	"log"

	pb "github.com/pengswift/geoip/geoip"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:60052"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGeoIPServiceClient(conn)

	r, err := c.QueryCity(context.Background(), &pb.GeoIPRequest{Ip: "36.63.199.104"})
	if err != nil {
		log.Fatalf("could not QueryCity: %v", err)
	}
	log.Printf("City: %s\n", r.Name)

	r, err = c.QueryCountry(context.Background(), &pb.GeoIPRequest{Ip: "36.63.199.104"})
	if err != nil {
		log.Fatalf("could not QueryCountry: %v", err)
	}
	log.Printf("Country: %s\n", r.Name)

	r, err = c.QuerySubdivision(context.Background(), &pb.GeoIPRequest{Ip: "36.63.199.104"})
	if err != nil {
		log.Fatalf("could not QuerySubdivision: %v", err)
	}
	log.Printf("Subdivision: %s\n", r.Name)
}
