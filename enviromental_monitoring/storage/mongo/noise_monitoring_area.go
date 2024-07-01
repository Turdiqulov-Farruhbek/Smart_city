package mongo

import (
	"context"
	"errors"
	"time"

	pb "gitlab.com/acumen/smart_city/enviromental_monitoring/genproto/enviromental_monitoring"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type NoiseMonitoringAreaService struct {
	Collection *mongo.Collection
}

func NewNoiseMonitoringAreaService(db *mongo.Database) *NoiseMonitoringAreaService {
	return &NoiseMonitoringAreaService{
		Collection: db.Collection("noise_monitoring_areas"),
	}
}

func (s *NoiseMonitoringAreaService) CreateNoiseMonitoringArea(req *pb.NoiseAreaCreate) (*pb.NoiseArea, error) {
	now := time.Now().Format(time.RFC3339)
	area := &pb.NoiseArea{
		ZoneId:      req.ZoneId,
		ZoneName:    req.ZoneName,
		AreaCovered: req.AreaCovered,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	_, err := s.Collection.InsertOne(context.TODO(), area)
	if err != nil {
		return nil, err
	}
	return area, nil
}

func (s *NoiseMonitoringAreaService) UpdateNoiseMonitoringArea(req *pb.NoiseAreaCreate) (*pb.NoiseArea, error) {
	if req.ZoneId == "" {
		return nil, errors.New("ZoneId cannot be empty")
	}

	updateDoc := bson.M{
		"ZoneName":    req.ZoneName,
		"AreaCovered": req.AreaCovered,
		"UpdatedAt":   time.Now().Format(time.RFC3339),
	}

	options := options.FindOneAndUpdate().SetReturnDocument(options.After)
	filter := bson.M{"ZoneId": req.ZoneId}

	var updatedArea pb.NoiseArea
	err := s.Collection.FindOneAndUpdate(context.TODO(), filter, bson.M{"$set": updateDoc}, options).Decode(&updatedArea)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("NoiseMonitoringArea not found")
	} else if err != nil {
		return nil, err
	}

	return &updatedArea, nil
}

func (s *NoiseMonitoringAreaService) DeleteNoiseMonitoringArea(id *pb.ById) error {
	if id.Id == "" {
		return errors.New("ZoneId cannot be empty")
	}

	_, err := s.Collection.DeleteOne(context.TODO(), bson.M{"ZoneId": id.Id})
	if err != nil {
		return err
	}

	return nil
}

func (s *NoiseMonitoringAreaService) GetNoiseMonitoringArea(id *pb.ById) (*pb.NoiseArea, error) {
	if id.Id == "" {
		return nil, errors.New("ZoneId cannot be empty")
	}

	var area pb.NoiseArea
	err := s.Collection.FindOne(context.TODO(), bson.M{"ZoneId": id.Id}).Decode(&area)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("NoiseMonitoringArea not found")
	} else if err != nil {
		return nil, err
	}

	return &area, nil
}
