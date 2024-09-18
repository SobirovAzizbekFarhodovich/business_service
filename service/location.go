package service

import (
	"context"

	pb "business/genprotos"
	"business/storage"
)

type LocationService struct {
	storage storage.StorageI
	pb.UnimplementedLocationServer
}

func NewLocationService(storage storage.StorageI) *LocationService {
	return &LocationService{storage: storage}
}

func (l *LocationService) CreateLocation(ctx context.Context, req *pb.CreateLocationRequest) (*pb.CreateLocationResponse, error) {
	_, err := l.storage.Location().CreateLocation(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *LocationService) DeleteLocation(ctx context.Context, req *pb.DeleteLocationRequest) (*pb.DeleteLocationResponse, error) {
	_, err := l.storage.Location().DeleteLocation(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *LocationService) GetLocationById(ctx context.Context, req *pb.GetLocationByIdRequest) (*pb.GetLocationByIdResponse, error) {
	res, err := l.storage.Location().GetLocationById(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (l *LocationService) GetAllLocation(ctx context.Context, req *pb.GetAllLocationRequest) (*pb.GetAllLocationResponse, error) {
	res, err := l.storage.Location().GetAllLocations(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
