package service

import (
	"context"

	pb "business/genprotos"
	"business/storage"
)

type BusinessService struct {
	storage storage.StorageI
	pb.UnimplementedBusinessServer
}

func NewBusinessService(storage storage.StorageI) *BusinessService {
	return &BusinessService{storage: storage}
}

func (l *BusinessService) CreateBusiness(ctx context.Context, req *pb.CreateBusinessRequest) (*pb.CreateBusinessResponse, error) {
	_, err := l.storage.Business().CreateBusiness(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *BusinessService) UpdateBusiness(ctx context.Context, req *pb.UpdateBusinessRequest) (*pb.UpdateBusinessResponse, error) {
	res, err := l.storage.Business().UpdateBusiness(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (l *BusinessService) DeleteBusiness(ctx context.Context, req *pb.DeleteBusinessRequest) (*pb.DeleteBusinessResponse, error) {
	_, err := l.storage.Business().DeleteBusiness(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *BusinessService) GetByIdBusiness(ctx context.Context, req *pb.GetByIdBusinessRequest) (*pb.GetByIdBusinessResponse, error) {
	res, err := l.storage.Business().GetByIdBusiness(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (l *BusinessService) GetAllBusinesses(ctx context.Context, req *pb.GetAllBusinessesRequest) (*pb.GetAllBusinessesResponse, error) {
	res, err := l.storage.Business().GetAllBusinesses(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
