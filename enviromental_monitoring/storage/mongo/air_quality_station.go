package mongo

import (
	"context"
	"errors"

	pb "gitlab.com/acumen/smart_city/enviromental_monitoring/genproto/enviromental_monitoring"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AirStation struct {
	Station *mongo.Collection
}

func NewAirStation(db *mongo.Database) *AirStation {
	return &AirStation{
		Station: db.Collection("air_quality_stations"),
	}
}

func (s *AirStation) CreateAirStation(req *pb.Station) error {
	_, err := s.Station.InsertOne(context.TODO(), req)
	if err != nil {
		return err
	}
	return nil
}

func (s *AirStation) UpdateAirStation(req *pb.Station) error {
	if req.StationId == "" {
		return errors.New("StationId cannot be empty")
	}

	updateDoc := bson.M{}

	if req.Location != nil {
		if req.Location.Latitude != 0 || req.Location.Longitude != 0 {
			updateDoc["location"] = req.Location
		}
	}
	if req.InstallationDate != "" {
		updateDoc["installationDate"] = req.InstallationDate
	}

	if len(updateDoc) == 0 {
		return errors.New("no fields to update")
	}

	_, err := s.Station.UpdateOne(
		context.TODO(),
		bson.M{"_id": req.StationId},
		bson.M{"$set": updateDoc},
	)
	if err != nil {
		return err
	}

	return nil
}
func (s *AirStation) DeleteStation(id *pb.ById) error {
	if id.Id == "" {
		return errors.New("id cannot be empty")
	}

	_, err := s.Station.DeleteOne(context.TODO(), bson.M{"_id": id.Id})
	if err != nil {
		return err
	}

	return nil
}
func (s *AirStation) GetStation(id *pb.ById) (*pb.Station, error) {
	if id.Id == "" {
        return nil, errors.New("id cannot be empty")
    }

    var station pb.Station
    err := s.Station.FindOne(context.TODO(), bson.M{"_id": id.Id}).Decode(&station)
    if err == mongo.ErrNoDocuments {
        return nil, errors.New("station not found")
    } else if err!= nil {
        return nil, err
    }

    return &station, nil
}