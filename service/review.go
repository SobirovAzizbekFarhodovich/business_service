package service

import (
	"context"

	pb "business/genprotos"
	"business/storage"
)

type ReviewService struct {
	storage storage.StorageI
	pb.UnimplementedReviewsServer
}

func NewReviewService(storage storage.StorageI) *ReviewService {
	return &ReviewService{storage: storage}
}

func (l *ReviewService) CreateReview(ctx context.Context, req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	_, err := l.storage.Review().CreateReview(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *ReviewService) UpdateReview(ctx context.Context, req *pb.UpdateReviewRequest) (*pb.UpdateReviewResponse, error) {
	res, err := l.storage.Review().UpdateReview(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (l *ReviewService) DeleteReview(ctx context.Context, req *pb.DeleteReviewRequest) (*pb.DeleteReviewResponse, error) {
	_, err := l.storage.Review().DeleteReview(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *ReviewService) GetOwnReviews(ctx context.Context, req *pb.GetOwnReviewsRequest) (*pb.GetOwnReviewsResponse, error) {
	res, err := l.storage.Review().GetOwnReviews(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (l *ReviewService) GetReviewByBusinessId(ctx context.Context, req *pb.GetReviewByBusinessIdRequest) (*pb.GetReviewByBusinessIdResponse, error) {
	res, err := l.storage.Review().GetReviewByBusinessId(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
