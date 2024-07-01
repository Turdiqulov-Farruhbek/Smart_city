package mongo

import (
	"context"

	pb "gitlab.com/acumen/smart_city/enviromental_monitoring/genproto/enviromental_monitoring"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type NoiseLevelReadingService struct {
	Collection *mongo.Collection
}

func NewNoiseLevelReadingService(db *mongo.Database) *NoiseLevelReadingService {
	return &NoiseLevelReadingService{
		Collection: db.Collection("noise_level_readings"),
	}
}

func (s *NoiseLevelReadingService) CreateNoiseLevelReading(req *pb.NoiseLevelReading) error {
	_, err := s.Collection.InsertOne(context.TODO(), req)
	if err != nil {
		return err
	}
	return nil
}

func (s *NoiseLevelReadingService) GetNoiseLevelReading(filter *pb.ZoneFilter) (*pb.NoiseLevelReadingList, error) {
	findOptions := bson.M{}

	if filter.ZoneId != "" {
		findOptions["ZoneId"] = filter.ZoneId
	}
	if filter.Time != "" {
		findOptions["Time"] = filter.Time
	}

	cursor, err := s.Collection.Find(context.TODO(), findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var readings []*pb.NoiseLevelReading
	for cursor.Next(context.TODO()) {
		var reading pb.NoiseLevelReading
		err := cursor.Decode(&reading)
		if err != nil {
			return nil, err
		}
		readings = append(readings, &reading)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.NoiseLevelReadingList{NoiseLevelReadings: readings}, nil
}
