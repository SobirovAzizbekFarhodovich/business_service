package service

import (
	"context"

	pb "business/genprotos"
	"business/storage"
)

type BookmarkedBusinessService struct {
	storage storage.StorageI
	pb.UnimplementedBookmarked_BusinessesServer
}

func NewBookmarkedBusinessService(storage storage.StorageI) *BookmarkedBusinessService {
	return &BookmarkedBusinessService{storage: storage}
}

func (l *BookmarkedBusinessService) CreateBookmarkedBusiness(ctx context.Context, req *pb.CreateBookmarkedBusRequest) (*pb.CreateBookmarkedBusResponse, error) {
	_, err := l.storage.BookmarkedBusiness().CreateBookmarkedBus(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *BookmarkedBusinessService) DeleteBookmarkedBusiness(ctx context.Context, req *pb.DeleteBookmarkedBusRequest) (*pb.DeleteBookmarkedBusResponse, error) {
	_, err := l.storage.BookmarkedBusiness().DeleteBookmarkedBus(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *BookmarkedBusinessService) GetBookmarkedBusById(req *pb.GetBookmarkedBusByIdRequest) (*pb.GetBookmarkedBusByIdResponse, error) {
	res, err := l.storage.BookmarkedBusiness().GetBookmarkedBusById(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (l *BookmarkedBusinessService) GetAllBookmarkedBus(req *pb.GetAllBookmarkedBusRequest) (*pb.GetAllBookmarkedBusResponse, error) {
	res, err := l.storage.BookmarkedBusiness().GetAllBookmarkedBus(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
