package postgres

import (
	pb "business/genprotos"
	"database/sql"
	"errors"
	"fmt"
	"strings"
)

type BusinessStorage struct {
	db *sql.DB
}

func NewBusinessStorage(db *sql.DB) *BusinessStorage {
	return &BusinessStorage{db: db}
}

func (b *BusinessStorage) CreateBusiness(req *pb.CreateBusinessRequest) (*pb.CreateBusinessResponse, error) {
	query := `
	INSERT INTO businesses(
		name,
		owner_id,
		description,
		category,
		contact_info, 
		hours_of_operation,
		location_id)
	VALUES($1, $2, $3, $4, $5, $6, $7)`
	_, err := b.db.Exec(query, req.Name, req.OwnerId, req.Description, req.Category, req.ContactInfo, req.HoursOfOperation, req.LocationId)
	if err != nil {
		return nil, err
	}
	return &pb.CreateBusinessResponse{}, nil
}

func (b *BusinessStorage) UpdateBusiness(req *pb.UpdateBusinessRequest) (*pb.UpdateBusinessResponse, error) {
	query := `UPDATE businesses SET `
	var condition []string
	var args []interface{}

	if req.Name != "" && req.Name != "string" {
		condition = append(condition, fmt.Sprintf("name = $%d", len(args)+1))
		args = append(args, req.Name)
	}

	if req.Description != "" && req.Description != "string" {
		condition = append(condition, fmt.Sprintf("description = $%d", len(args)+1))
		args = append(args, req.Description)
	}
	if req.Category != "" && req.Category != "string" {
		condition = append(condition, fmt.Sprintf("category = $%d", len(args)+1))
		args = append(args, req.Category)
	}
	if req.ContactInfo != "" && req.ContactInfo != "string" {
		condition = append(condition, fmt.Sprintf("contact_info = $%d", len(args)+1))
		args = append(args, req.ContactInfo)
	}
	if req.HoursOfOperation != "" && req.HoursOfOperation != "string" {
		condition = append(condition, fmt.Sprintf("hours_of_operation = $%d", len(args)+1))
		args = append(args, req.HoursOfOperation)
	}
	if req.LocationId != "" && req.LocationId != "string" {
		condition = append(condition, fmt.Sprintf("location_id = $%d", len(args)+1))
		args = append(args, req.LocationId)
	}
	if len(condition) == 0 {
		return nil, errors.New("nothing to update")
	}

	query += strings.Join(condition, ", ")
	query += fmt.Sprintf(" WHERE id = $%d AND owner_id = $%d RETURNING owner_id,name,description,category,contact_info,hours_of_operation,location_id", len(args)+1,len(args) + 2)
	args = append(args, req.Id,req.OwnerId)

	res := pb.UpdateBusinessResponse{}
	row := b.db.QueryRow(query, args...)

	err := row.Scan(&res.OwnerId, &res.Name, &res.Description, &res.Category,&res.ContactInfo,&res.HoursOfOperation,&res.LocationId)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

func (b *BusinessStorage) DeleteBusiness(req *pb.DeleteBusinessRequest) (*pb.DeleteBusinessResponse, error) {
	query := `DELETE FROM businesses WHERE id = $1 AND owner_id = $2`
	_, err := b.db.Exec(query, req.Id, req.OwnerId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBusinessResponse{}, nil
}

func (b *BusinessStorage) GetByIdBusiness(req *pb.GetByIdBusinessRequest) (*pb.GetByIdBusinessResponse, error) {
	query := `SELECT id, name, description, category, contact_info, hours_of_operation, owner_id, average_ratings, location_id 
	          FROM businesses WHERE id = $1`
	var id, name, description, category, contactInfo, hoursOfOperation, ownerId, locationId string
	var averageRatings float32

	err := b.db.QueryRow(query, req.Id).Scan(&id, &name, &description, &category, &contactInfo, &hoursOfOperation, &ownerId, &averageRatings, &locationId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("business not found")
		}
		return nil, err
	}

	response := &pb.GetByIdBusinessResponse{
		Id:               id,
		Name:             name,
		Description:      description,
		Category:         category,
		ContactInfo:      contactInfo,
		HoursOfOperation: hoursOfOperation,
		OwnerId:          ownerId,
		AverageRatings:   averageRatings,
		LocationId:       locationId,
	}
	return response, nil
}

func (b *BusinessStorage) GetAllBusinesses(req *pb.GetAllBusinessesRequest) (*pb.GetAllBusinessesResponse, error) {
	query := `SELECT id, name, description, category, contact_info, hours_of_operation, owner_id, average_ratings, location_id 
	          FROM businesses LIMIT 10 OFFSET $1`
	offset := (req.Page - 1) * 10

	rows, err := b.db.Query(query, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var businesses []*pb.GetByIdBusinessResponse
	for rows.Next() {
		var id, name, description, category, contactInfo, hoursOfOperation, ownerId, locationId string
		var averageRatings float32
		err := rows.Scan(&id, &name, &description, &category, &contactInfo, &hoursOfOperation, &ownerId, &averageRatings, &locationId)
		if err != nil {
			continue
		}

		business := &pb.GetByIdBusinessResponse{
			Id:               id,
			Name:             name,
			Description:      description,
			Category:         category,
			ContactInfo:      contactInfo,
			HoursOfOperation: hoursOfOperation,
			OwnerId:          ownerId,
			AverageRatings:   averageRatings,
			LocationId:       locationId,
		}
		businesses = append(businesses, business)
	}

	response := &pb.GetAllBusinessesResponse{
		Businesses: businesses,
	}
	return response, nil
}
