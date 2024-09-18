package postgres

import (
	pb "business/genprotos"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type ReviewsStorage struct {
	db *sql.DB
}

func NewReviewsStorage(db *sql.DB) *ReviewsStorage {
	return &ReviewsStorage{db: db}
}

func (r *ReviewsStorage) CreateReview(req *pb.CreateReviewRequest) (*pb.CreateReviewResponse, error) {
	query := `
	INSERT INTO reviews(
		business_id,
		user_id,
		rating,
		text)
	VALUES($1, $2, $3, $4)`
	_, err := r.db.Exec(query, req.BusinessId, req.UserId, req.Rating, req.Text)
	if err != nil {
		return nil, err
	}
	return &pb.CreateReviewResponse{}, nil
}

func (r *ReviewsStorage) UpdateReview(req *pb.UpdateReviewRequest) (*pb.UpdateReviewResponse, error) {
	query := `UPDATE reviews SET `
	var conditions []string
	var args []interface{}

	if req.Rating != 0 {
		conditions = append(conditions, fmt.Sprintf("rating = $%d", len(args)+1))
		args = append(args, req.Rating)
	}
	if req.Text != "" && req.Text != "string" {
		conditions = append(conditions, fmt.Sprintf("text = $%d", len(args)+1))
		args = append(args, req.Text)
	}

	if len(conditions) == 0 {
		return nil, errors.New("nothing to update")
	}

	query += strings.Join(conditions, ", ")
	query += fmt.Sprintf(" WHERE id = $%d AND business_id = $%d AND user_id = $%d RETURNING id, business_id, user_id, rating, text", len(args)+1, len(args)+2, len(args)+3)

	args = append(args, req.Id, req.BusinessId, req.UserId)

	res := pb.UpdateReviewResponse{}

	err := r.db.QueryRow(query, args...).Scan(&res.Id, &res.BusinessId, &res.UserId, &res.Rating, &res.Text)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("review with id %s, business_id %s, and user_id %s not found", req.Id, req.BusinessId, req.UserId)
		}
		return nil, err
	}

	return &res, nil
}

func (r *ReviewsStorage) DeleteReview(req *pb.DeleteReviewRequest) (*pb.DeleteReviewResponse, error) {
	query := `DELETE FROM reviews WHERE id = $1 AND user_id = $2`
	_, err := r.db.Exec(query, req.Id, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteReviewResponse{}, nil
}

func (r *ReviewsStorage) GetOwnReviews(req *pb.GetOwnReviewsRequest) (*pb.GetOwnReviewsResponse, error) {
	query := `SELECT id, business_id, user_id, rating, text FROM reviews WHERE user_id = $1`
	rows, err := r.db.Query(query, req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*pb.UpdateReviewRequest
	for rows.Next() {
		var id, businessId, userId, text string
		var rating int32
		err := rows.Scan(&id, &businessId, &userId, &rating, &text)
		if err != nil {
			return nil, err
		}

		review := &pb.UpdateReviewRequest{
			Id:         id,
			BusinessId: businessId,
			UserId:     userId,
			Rating:     rating,
			Text:       text,
		}
		reviews = append(reviews, review)
	}

	response := &pb.GetOwnReviewsResponse{
		Reviews: reviews,
	}
	return response, nil
}

func (r *ReviewsStorage) GetReviewByBusinessId(req *pb.GetReviewByBusinessIdRequest) (*pb.GetReviewByBusinessIdResponse, error) {
	query := `SELECT id, business_id, user_id, rating, text FROM reviews WHERE business_id = $1`
	rows, err := r.db.Query(query, req.BusinessId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reviews []*pb.UpdateReviewRequest
	for rows.Next() {
		var id, businessId, userId, text string
		var rating int32
		err := rows.Scan(&id, &businessId, &userId, &rating, &text)
		if err != nil {
			return nil, err
		}

		review := &pb.UpdateReviewRequest{
			Id:         id,
			BusinessId: businessId,
			UserId:     userId,
			Rating:     rating,
			Text:       text,
		}
		reviews = append(reviews, review)
	}

	response := &pb.GetReviewByBusinessIdResponse{
		Reviews: reviews,
	}
	return response, nil
}
