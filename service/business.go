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
	businessRes, err := l.storage.Business().GetByIdBusiness(req)
	if err != nil {
		return nil, err
	}

	averageRating, err := l.storage.Business().GetAverageRatingByBusinessId(req.Id)
	if err != nil {
		return nil, err
	}
	businessRes.AverageRatings = float32(averageRating)

	if businessRes.LocationId != "" {
		locationRes, err := l.storage.Location().GetLocationById(&pb.GetLocationByIdRequest{Id: businessRes.LocationId})
		if err != nil {
			return nil, err
		}
		businessRes.Location = &pb.Location{
			Id:        locationRes.Id,
			Latitude:  locationRes.Latitude,
			Longitude: locationRes.Longitude,
			Address:   locationRes.Address,
		}
	}

	return businessRes, nil
}

func (l *BusinessService) GetAllBusinesses(ctx context.Context, req *pb.GetAllBusinessesRequest) (*pb.GetAllBusinessesResponse, error) {
	businessesRes, err := l.storage.Business().GetAllBusinesses(req)
	if err != nil {
		return nil, err
	}

	for _, business := range businessesRes.Businesses {
		averageRating, err := l.storage.Business().GetAverageRatingByBusinessId(business.Id)
		if err != nil {
			return nil, err
		}
		business.AverageRatings = float32(averageRating)

		if business.LocationId != "" {
			locationRes, err := l.storage.Location().GetLocationById(&pb.GetLocationByIdRequest{Id: business.LocationId})
			if err != nil {
				return nil, err
			}
			business.Location = &pb.Location{
				Id:        locationRes.Id,
				Latitude:  locationRes.Latitude,
				Longitude: locationRes.Longitude,
				Address:   locationRes.Address,
			}
		} else {
			business.Location = nil
		}
	}

	return businessesRes, nil
}
