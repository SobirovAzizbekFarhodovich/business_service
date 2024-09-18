package service

import (
	"context"

	pb "business/genprotos"
	"business/storage"
)

type BusinessPhotosService struct {
	storage storage.StorageI
	pb.UnimplementedBusiness_PhotosServer
}

func NewBusinessPhotosService(storage storage.StorageI) *BusinessPhotosService {
	return &BusinessPhotosService{storage: storage}
}

func (l *BusinessPhotosService) CreateBusinessPhotos(ctx context.Context, req *pb.CreateBusinessPhotosRequest) (*pb.CreateBusinessPhotosResponse, error) {
	_, err := l.storage.BusinessPhotos().CreateBusinessPhotos(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *BusinessPhotosService) UpdateBusinessPhotos(ctx context.Context, req *pb.UpdateBusinessPhotosRequest) (*pb.UpdateBusinessPhotosResponse, error) {
	res, err := l.storage.BusinessPhotos().UpdateBusinessPhotos(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (l *BusinessPhotosService) DeleteBusinessPhotos(ctx context.Context, req *pb.DeleteBusinessPhotosRequest) (*pb.DeleteBusinessPhotosResponse, error) {
	_, err := l.storage.BusinessPhotos().DeleteBusinessPhotos(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *BusinessPhotosService) GetByBusinessId(ctx context.Context, req *pb.GetBusinessIdRequest) (*pb.GetBusinessIdResponse, error) {
	res, err := l.storage.BusinessPhotos().GetByBusinessId(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (l *BusinessPhotosService) GetBusinessPhotosByOwner(ctx context.Context, req *pb.GetBusinessPhotosByOwnerRequest) (*pb.GetBusinessPhotosByOwnerResponse, error) {
	res, err := l.storage.BusinessPhotos().GetBusinessPhotosByOwner(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
