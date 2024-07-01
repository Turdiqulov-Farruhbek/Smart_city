package mongo

import (
	"context"

	"gitlab.com/acumen/smart_city/enviromental_monitoring/storage"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Storage struct {
	AirStationS storage.AirStationI
}

func ConnectMongo() (*Storage,error) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err!= nil {
        return nil,err
    }
	err = client.Ping(context.TODO(), nil)
	if err!= nil {
        return nil,err
    }
	air_station := NewAirStation(client.Database("enviromental_monitoring"))
	
	return &Storage{
        AirStationS: air_station,
    }, nil
}
