package postgres

import (
	pb "business/genprotos"
	"database/sql"
)

type BusinessPhotosStorage struct {
	db *sql.DB
}

func NewBusinessPhotosStorage(db *sql.DB) *BusinessPhotosStorage{
	return &BusinessPhotosStorage{db: db}
}

func (b *BusinessPhotosStorage) CreateBusinessPhotos(req *pb.CreateBusinessPhotosRequest) (*pb.CreateBusinessPhotosResponse, error) {
	query := `
	INSERT INTO business_photos(
		business_id,
		photo_url,
		owner_id)
	VALUES($1, $2, $3)`
	_, err := b.db.Exec(query, req.BusinessId, req.PhotoUrl, req.OwnerId)
	if err != nil {
		return nil, err
	}
	return &pb.CreateBusinessPhotosResponse{}, nil
}

func (b *BusinessPhotosStorage) UpdateBusinessPhotos(req *pb.UpdateBusinessPhotosRequest) (*pb.UpdateBusinessPhotosResponse, error) {
	query := `
	UPDATE business_photos
	SET photo_url = $1
	WHERE business_id = $2 AND owner_id = $3`
	_, err := b.db.Exec(query, req.PhotoUrl, req.BusinessId, req.OwnerId)
	if err != nil {
		return nil, err
	}

	response := &pb.UpdateBusinessPhotosResponse{
		BusinessId: req.BusinessId,
		PhotoUrl:   req.PhotoUrl,
		OwnerId:    req.OwnerId,
	}

	return response, nil
}

func (b *BusinessPhotosStorage) DeleteBusinessPhotos(req *pb.DeleteBusinessPhotosRequest) (*pb.DeleteBusinessPhotosResponse, error) {
	query := `DELETE FROM business_photos WHERE id = $1 AND owner_id = $2`
	_, err := b.db.Exec(query, req.Id, req.OwnerId)
	if err != nil {
		return nil, err
	}
	return &pb.DeleteBusinessPhotosResponse{}, nil
}

func (b *BusinessPhotosStorage) GetByBusinessId(req *pb.GetBusinessIdRequest) (*pb.GetBusinessIdResponse, error) {
	query := `SELECT photo_url, owner_id, business_id FROM business_photos WHERE business_id = $1`
	rows, err := b.db.Query(query, req.BusinessId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var photos []*pb.Business
	var ownerId, businessId string
	for rows.Next() {
		var photoUrl string
		err := rows.Scan(&photoUrl, &ownerId, &businessId)
		if err != nil {
			return nil, err
		}
		photos = append(photos, &pb.Business{
			PhotoUrl: photoUrl,
		})
	}

	response := &pb.GetBusinessIdResponse{
		Photos:     photos,
		OwnerId:    ownerId,
		BusinessId: businessId,
	}

	return response, nil
}

func (b *BusinessPhotosStorage) GetBusinessPhotosByOwner(req *pb.GetBusinessPhotosByOwnerRequest) (*pb.GetBusinessPhotosByOwnerResponse, error) {
	query := `SELECT business_id, owner_id, photo_url FROM business_photos WHERE owner_id = $1`
	rows, err := b.db.Query(query, req.OwnerId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var responses []*pb.GetBusinessIdResponse
	for rows.Next() {
		var businessId, ownerId, photoUrl string
		err := rows.Scan(&businessId, &ownerId, &photoUrl)
		if err != nil {
			return nil, err
		}

		photos := []*pb.Business{
			{PhotoUrl: photoUrl},
		}

		response := &pb.GetBusinessIdResponse{
			Photos:     photos,
			OwnerId:    ownerId,
			BusinessId: businessId,
		}
		responses = append(responses, response)
	}

	finalResponse := &pb.GetBusinessPhotosByOwnerResponse{
		Photos: responses,
	}

	return finalResponse, nil
}

