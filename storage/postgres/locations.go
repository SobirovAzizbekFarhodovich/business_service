package postgres

import (
    pb "business/genprotos"
    "database/sql"
    "encoding/json"
    "fmt"
    "io"
    "log"
    "net/http"
)

type LocationStorage struct {
    db *sql.DB
}

func NewLocationStorage(db *sql.DB) *LocationStorage {
    return &LocationStorage{db: db}
}

type GeocodingResponse struct {
    PlaceID     int    `json:"place_id"`
    Licence     string `json:"licence"`
    OsmType     string `json:"osm_type"`
    OsmID       int    `json:"osm_id"`
    Lat         string `json:"lat"`
    Lon         string `json:"lon"`
    DisplayName string `json:"display_name"`
    Address     struct {
        HouseNumber   string `json:"house_number"`
        Road          string `json:"road"`
        Neighbourhood string `json:"neighbourhood"`
        County        string `json:"county"`
        City          string `json:"city"`
        Postcode      string `json:"postcode"`
        Country       string `json:"country"`
        CountryCode   string `json:"country_code"`
    } `json:"address"`
    BoundingBox []string `json:"boundingbox"`
}

func getAddressFromCoordinates(lat, lon, apiKey string) (string, error) {
    url := fmt.Sprintf("https://geocode.maps.co/reverse?lat=%s&lon=%s&api_key=%s", lat, lon, apiKey)
    resp, err := http.Get(url)
    if err != nil {
        return "", fmt.Errorf("failed to send request: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", fmt.Errorf("failed to read response body: %v", err)
    }


    var geocodeResponse GeocodingResponse
    err = json.Unmarshal(body, &geocodeResponse)
    if err != nil {
        return "", fmt.Errorf("failed to parse geocode response: %v", err)
    }


    address := fmt.Sprintf("%s %s, %s, %s, %s, %s, %s",
        geocodeResponse.Address.HouseNumber,
        geocodeResponse.Address.Road,
        geocodeResponse.Address.Neighbourhood,
        geocodeResponse.Address.County,
        geocodeResponse.Address.City,
        geocodeResponse.Address.Postcode,
        geocodeResponse.Address.Country)

    return address, nil
}

func (l *LocationStorage) CreateLocation(req *pb.CreateLocationRequest) (*pb.CreateLocationResponse, error) {
    apiKey := "66e9901ee798b172186971fvy42a182"

    lat := fmt.Sprintf("%f", req.Latitude)
    lon := fmt.Sprintf("%f", req.Longitude)

    var address string
    if req.Address == "" || req.Address == "string" {
        var err error
        address, err = getAddressFromCoordinates(lat, lon, apiKey)
        if err != nil {
            return nil, err
        }
    } else {
        address = req.Address
    }

    query := `INSERT INTO locations (latitude, longitude, address) VALUES ($1, $2, $3) RETURNING id`
    var locationID string
    err := l.db.QueryRow(query, req.Latitude, req.Longitude, address).Scan(&locationID)
    if err != nil {
        return nil, err
    }

    response := &pb.CreateLocationResponse{
        Id: locationID,
    }

    return response, nil
}



func (l *LocationStorage) DeleteLocation(req *pb.DeleteLocationRequest)(*pb.DeleteLocationResponse, error){
	query := `DELETE FROM locations WHERE id = $1`
	_, err := l.db.Exec(query, req.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (l *LocationStorage) GetLocationById(req *pb.GetLocationByIdRequest) (*pb.GetLocationByIdResponse, error) {
	query := `SELECT latitude, longitude, address FROM locations WHERE id = $1`
	var latitude, longitude float32
	var address string
	err := l.db.QueryRow(query, req.Id).Scan(&latitude, &longitude, &address)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("location not found")
		}
		log.Printf("Failed to retrieve location: %v", err)
		return nil, err
	}

	response := &pb.GetLocationByIdResponse{
		Latitude:  latitude,
		Longitude: longitude,
		Address:   address,
		Id:        req.Id,
	}
	return response, nil
}

func (l *LocationStorage) GetAllLocations(req *pb.GetAllLocationRequest) (*pb.GetAllLocationResponse, error) {
	query := `SELECT id, latitude, longitude, address FROM locations LIMIT 10 OFFSET $1`
	offset := (req.Page - 1) * 10

	rows, err := l.db.Query(query, offset)
	if err != nil {
		log.Printf("Failed to retrieve locations: %v", err)
		return nil, err
	}
	defer rows.Close()

	var locations []*pb.GetLocationByIdResponse
	for rows.Next() {
		var id string
		var latitude, longitude float32
		var address string
		err := rows.Scan(&id, &latitude, &longitude, &address)
		if err != nil {
			log.Printf("Failed to scan row: %v", err)
			continue
		}

		locations = append(locations, &pb.GetLocationByIdResponse{
			Id:        id,
			Latitude:  latitude,
			Longitude: longitude,
			Address:   address,
		})
	}

	response := &pb.GetAllLocationResponse{
		Locations: locations,
	}
	return response, nil
}
