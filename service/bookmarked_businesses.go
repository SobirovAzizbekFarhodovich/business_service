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

func (l *BookmarkedBusinessService) CreateBookmarkedBusiness(ctx context.Context, req *pb.CreateBookmarkedBusinessRequest) (*pb.CreateBookmarkedBusinessResponse, error) {
	_, err := l.storage.BookmarkedBusiness().CreateBookmarkedBus(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *BookmarkedBusinessService) DeleteBookmarkedBusiness(ctx context.Context, req *pb.DeleteBookmarkedBusinessRequest) (*pb.DeleteBookmarkedBusinessResponse, error) {
	_, err := l.storage.BookmarkedBusiness().DeleteBookmarkedBus(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *BookmarkedBusinessService) GetBookmarkedBusinessById(ctx context.Context, req *pb.GetBookmarkedBusinessByIdRequest) (*pb.GetBookmarkedBusinessByIdResponse, error) {
	res, err := l.storage.BookmarkedBusiness().GetBookmarkedBusinessById(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (l *BookmarkedBusinessService) GetAllBookmarkedBusiness(ctx context.Context, req *pb.GetAllBookmarkedBusinessRequest) (*pb.GetAllBookmarkedBusinessResponse, error) {
	res, err := l.storage.BookmarkedBusiness().GetAllBookmarkedBus(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
