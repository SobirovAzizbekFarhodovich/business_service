package main

import (
	"log"
	"net"

	pb "business/genprotos"
	"business/service"
	"business/storage/postgres"
	"google.golang.org/grpc"
)

func main(){
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal(err)
	}

	liss, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	pb.RegisterBookmarked_BusinessesServer(s, service.NewBookmarkedBusinessService(db))
	pb.RegisterBusinessServer(s, service.NewBusinessService(db))
	pb.RegisterBusiness_PhotosServer(s, service.NewBusinessPhotosService(db))
	pb.RegisterLocationServer(s, service.NewLocationService(db))
	pb.RegisterReviewsServer(s,service.NewReviewService(db))

	log.Printf("server listening at %v", liss.Addr())
	if err := s.Serve(liss); err != nil {
		log.Fatalf("failed to serve %v", err)
	}
}
