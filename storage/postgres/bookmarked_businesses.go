package postgres

import (
	pb "business/genprotos"
	"database/sql"
	"fmt"
)

type BookmarkedBusinessesStorage struct {
	db *sql.DB
}

func NewBookmarkedBusinessesStorage(db *sql.DB) *BookmarkedBusinessesStorage {
	return &BookmarkedBusinessesStorage{db: db}
}

func (b *BookmarkedBusinessesStorage) CreateBookmarkedBus(req *pb.CreateBookmarkedBusinessRequest) (*pb.CreateBookmarkedBusinessResponse, error) {
	query := `
	INSERT INTO bookmarked_businesses(
		user_id,
		business_id)
	VALUES($1, $2)`
	_, err := b.db.Exec(query, req.UserId, req.BusinessId)
	if err != nil {
		return nil, err
	}
	return &pb.CreateBookmarkedBusinessResponse{}, nil
}

func (b *BookmarkedBusinessesStorage) DeleteBookmarkedBus(req *pb.DeleteBookmarkedBusinessRequest) (*pb.DeleteBookmarkedBusinessResponse, error) {
	query := `DELETE FROM bookmarked_businesses WHERE id = $1 AND user_id = $2`
	_, err := b.db.Exec(query, req.Id, req.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBookmarkedBusinessResponse{}, nil
}

func (b *BookmarkedBusinessesStorage) GetBookmarkedBusinessById(req *pb.GetBookmarkedBusinessByIdRequest) (*pb.GetBookmarkedBusinessByIdResponse, error) {
	query := `SELECT user_id, business_id, id FROM bookmarked_businesses WHERE id = $1 AND user_id = $2`
	var userId, businessId, id string
	err := b.db.QueryRow(query, req.Id, req.UserId).Scan(&userId, &businessId, &id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("bookmark not found")
		}
		return nil, err
	}

	response := &pb.GetBookmarkedBusinessByIdResponse{
		UserId:     userId,
		BusinessId: businessId,
		Id:         id,
	}

	return response, nil
}

func (b *BookmarkedBusinessesStorage) GetAllBookmarkedBus(req *pb.GetAllBookmarkedBusinessRequest) (*pb.GetAllBookmarkedBusinessResponse, error) {
	query := `SELECT id, user_id, business_id FROM bookmarked_businesses WHERE user_id = $1`
	rows, err := b.db.Query(query, req.UserId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var businesses []*pb.GetBookmarkedBusinessByIdResponse
	for rows.Next() {
		var id, userId, businessId string
		err := rows.Scan(&id, &userId, &businessId)
		if err != nil {
			return nil, err
		}

		business := &pb.GetBookmarkedBusinessByIdResponse{
			Id:          id,
			UserId:      userId,
			BusinessId:  businessId,
		}
		businesses = append(businesses, business)
	}

	response := &pb.GetAllBookmarkedBusinessResponse{
		Businesses: businesses,
	}
	return response, nil
}
