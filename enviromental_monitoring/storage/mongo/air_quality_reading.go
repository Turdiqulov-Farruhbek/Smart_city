package mongo

import (
	"context"

	pb "gitlab.com/acumen/smart_city/enviromental_monitoring/genproto/enviromental_monitoring"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type AirQualityReadingRepo struct {
	Reading *mongo.Collection
	Station *mongo.Collection
}

func NewAirQualityReadingRepo(db *mongo.Database) *AirQualityReadingRepo {
	return &AirQualityReadingRepo{
		Reading: db.Collection("air_quality_readings"),
		Station: db.Collection("air_quality_stations"),
	}
}
func (ar *AirQualityReadingRepo) CreateAirQualityReading(req *pb.AirQualityReadingCreate) error {
	_, err := ar.Reading.InsertOne(context.TODO(), req)
	if err != nil {
		return err
	}
	return nil
}
func (ar *AirQualityReadingRepo) GetCityWideAirQuality(filter *pb.AirQualityFilter) (*pb.AirQualityList, error) {

	mongoFilter := bson.M{}
	if filter.StationId != "" {
		mongoFilter["StationId"] = filter.StationId
	}
	if filter.Time != "" {
		mongoFilter["Time"] = filter.Time
	}

	cursor, err := ar.Reading.Find(context.TODO(), mongoFilter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	var airQualityReadings []*pb.AirQualityReading

	for cursor.Next(context.TODO()) {
		var reading pb.AirQualityReadingCreate
		err := cursor.Decode(&reading)
		if err != nil {
			return nil, err
		}

		stationFilter := bson.M{"StationId": reading.StationId}
		var station pb.Station
		err = ar.Station.FindOne(context.TODO(), stationFilter).Decode(&station)
		if err != nil {
			return nil, err
		}

		airQualityReading := &pb.AirQualityReading{
			ReadingId:   reading.ReadingId,
			Station:     &station,
			Pm25Level:  reading.Pm25Level,
			Pm10Level:  reading.Pm10Level,
			OzoneLevel: reading.OzoneLevel,
			Time:        reading.Time,
		}

		airQualityReadings = append(airQualityReadings, airQualityReading)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.AirQualityList{AirQualityReadings: airQualityReadings}, nil
}
