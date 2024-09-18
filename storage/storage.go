package storage

import pb "business/genprotos"

type StorageI interface {
	Business() BusinessI
	Location() LocationI
	Review() ReviewI
	BusinessPhotos() BusinessPhotosI
	BookmarkedBusiness() BookmarkedBusinessI
}

type BusinessI interface {
	CreateBusiness(req *pb.CreateBusinessRequest) (*pb.CreateBusinessResponse, error)
	UpdateBusiness(req *pb.UpdateBusinessRequest) (*pb.UpdateBusinessResponse, error)
	DeleteBusiness(req *pb.DeleteBusinessRequest) (*pb.DeleteBusinessResponse, error)
	GetByIdBusiness(req *pb.GetByIdBusinessRequest) (*pb.GetByIdBusinessResponse, error)
	GetAllBusinesses(req *pb.GetAllBusinessesRequest) (*pb.GetAllBusinessesResponse, error)
	GetAverageRatingByBusinessId(businessId string) (float32, error)
	
}

type LocationI interface {
	CreateLocation(req *pb.CreateLocationRequest) (*pb.CreateLocationResponse, error)
	DeleteLocation(req *pb.DeleteLocationRequest) (*pb.DeleteLocationResponse, error)
	GetLocationById(req *pb.GetLocationByIdRequest) (*pb.GetLocationByIdResponse, error)
	GetAllLocations(req *pb.GetAllLocationRequest) (*pb.GetAllLocationResponse, error)
}

type ReviewI interface {
	CreateReview(req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error)
	UpdateReview(req *pb.UpdateReviewRequest) (*pb.UpdateReviewResponse, error)
	DeleteReview(req *pb.DeleteReviewRequest) (*pb.DeleteReviewResponse, error)
	GetOwnReviews(req *pb.GetOwnReviewsRequest) (*pb.GetOwnReviewsResponse, error)
	GetReviewByBusinessId(req *pb.GetReviewByBusinessIdRequest) (*pb.GetReviewByBusinessIdResponse, error)
}

type BusinessPhotosI interface{
	CreateBusinessPhotos(req *pb.CreateBusinessPhotosRequest) (*pb.CreateBusinessPhotosResponse, error)
	UpdateBusinessPhotos(req *pb.UpdateBusinessPhotosRequest) (*pb.UpdateBusinessPhotosResponse, error)
	DeleteBusinessPhotos(req *pb.DeleteBusinessPhotosRequest) (*pb.DeleteBusinessPhotosResponse, error)
	GetByBusinessId(req *pb.GetBusinessIdRequest) (*pb.GetBusinessIdResponse, error)
	GetBusinessPhotosByOwner(req *pb.GetBusinessPhotosByOwnerRequest) (*pb.GetBusinessPhotosByOwnerResponse, error)
}

type BookmarkedBusinessI interface{
	CreateBookmarkedBus(req *pb.CreateBookmarkedBusinessRequest) (*pb.CreateBookmarkedBusinessResponse, error)
	DeleteBookmarkedBus(req *pb.DeleteBookmarkedBusinessRequest) (*pb.DeleteBookmarkedBusinessResponse, error)
	GetBookmarkedBusinessById(req *pb.GetBookmarkedBusinessByIdRequest) (*pb.GetBookmarkedBusinessByIdResponse, error)
	GetAllBookmarkedBus(req *pb.GetAllBookmarkedBusinessRequest) (*pb.GetAllBookmarkedBusinessResponse, error)

}