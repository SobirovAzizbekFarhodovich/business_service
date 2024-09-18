package postgres

import (
	"database/sql"
	"fmt"

	"business/config"
	u "business/storage"

	_ "github.com/lib/pq"
)

type Storage struct {
	Db                   *sql.DB
	BusinessC            u.BusinessI
	LocationC            u.LocationI
	BusinessPhotosC      u.BusinessPhotosI
	BookmarkedBusinessC  u.BookmarkedBusinessI
	ReviewC              u.ReviewI
}

func NewPostgresStorage() (u.StorageI, error) {
	config := config.Load()
	con := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.PostgresUser, config.PostgresPassword,
		config.PostgresHost, config.PostgresPort,
		config.PostgresDatabase)
	db, err := sql.Open("postgres", con)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	BusinessS := NewBusinessStorage(db)
	LocationS := NewLocationStorage(db)
	BusinessPhotosS := NewBusinessPhotosStorage(db)
	ReviewS := NewReviewsStorage(db)
	BookmarkedBusinessS := NewBookmarkedBusinessesStorage(db)

	return &Storage{
		BusinessC:           BusinessS,
		LocationC:           LocationS,
		BusinessPhotosC:     BusinessPhotosS,
		ReviewC:             ReviewS,
		BookmarkedBusinessC: BookmarkedBusinessS,
	}, nil
}

func (s *Storage) Business() u.BusinessI {
	if s.BusinessC == nil {
		s.BusinessC = NewBusinessStorage(s.Db)
	}
	return s.BusinessC
}

func (s *Storage) Location() u.LocationI {
	if s.LocationC == nil {
		s.LocationC = NewLocationStorage(s.Db)
	}
	return s.LocationC
}

func (s *Storage) BusinessPhotos() u.BusinessPhotosI {
	if s.BusinessPhotosC == nil {
		s.BusinessPhotosC = NewBusinessPhotosStorage(s.Db)
	}
	return s.BusinessPhotosC
}

func (s *Storage) BookmarkedBusiness() u.BookmarkedBusinessI {
	if s.BookmarkedBusinessC == nil {
		s.BookmarkedBusinessC = NewBookmarkedBusinessesStorage(s.Db)
	}
	return s.BookmarkedBusinessC
}

func (s *Storage) Review() u.ReviewI {
	if s.ReviewC == nil {
		s.ReviewC = NewReviewsStorage(s.Db)
	}
	return s.ReviewC
}
