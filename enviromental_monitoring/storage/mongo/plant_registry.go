package mongo

import (
	"context"
	"errors"
	"time"

	pb "gitlab.com/acumen/smart_city/enviromental_monitoring/genproto/enviromental_monitoring"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type PlantRegistryService struct {
	Collection *mongo.Collection
}

func NewPlantRegistryService(db *mongo.Database) *PlantRegistryService {
	return &PlantRegistryService{
		Collection: db.Collection("plant_registries"),
	}
}

func (s *PlantRegistryService) RegisterPlants(req *pb.PlantRegistry) error {
	if req.RegistryId == "" {
		req.RegistryId = generateUUID() // Generate UUID for new registry
	}

	_, err := s.Collection.InsertOne(context.TODO(), req)
	if err != nil {
		return err
	}
	return nil
}

func (s *PlantRegistryService) GetPlantRegistries(filter *pb.PlantRegistryFilter) (*pb.PlantRegistryList, error) {
	bsonFilter := bson.M{}

	if filter.RegistryId != "" {
		bsonFilter["RegistryId"] = filter.RegistryId
	}
	if filter.SpaceId != "" {
		bsonFilter["SpaceId"] = filter.SpaceId
	}
	if filter.SpeciesName != "" {
		bsonFilter["SpeciesName"] = filter.SpeciesName
	}
	if filter.Quantity != 0 {
		bsonFilter["Quantity"] = filter.Quantity
	}
	if filter.PlantDate != "" {
		bsonFilter["PlantDate"] = filter.PlantDate
	}

	cur, err := s.Collection.Find(context.TODO(), bsonFilter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var plantRegistries []*pb.PlantRegistry
	for cur.Next(context.TODO()) {
		var registry pb.PlantRegistry
		err := cur.Decode(&registry)
		if err != nil {
			return nil, err
		}
		plantRegistries = append(plantRegistries, &registry)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &pb.PlantRegistryList{PlantRegistries: plantRegistries}, nil
}

func (s *PlantRegistryService) UpdatePlantRegistry(req *pb.PlantRegistry) error {
	if req.RegistryId == "" {
		return errors.New("RegistryId cannot be empty")
	}

	updateDoc := bson.M{}

	if req.SpaceId != "" {
		updateDoc["SpaceId"] = req.SpaceId
	}
	if req.SpeciesName != "" {
		updateDoc["SpeciesName"] = req.SpeciesName
	}
	if req.Quantity != 0 {
		updateDoc["Quantity"] = req.Quantity
	}
	if req.PlantDate != "" {
		updateDoc["PlantDate"] = req.PlantDate
	}
	updateDoc["UpdatedAt"] = time.Now().Format(time.RFC3339)

	_, err := s.Collection.UpdateOne(
		context.TODO(),
		bson.M{"RegistryId": req.RegistryId},
		bson.M{"$set": updateDoc},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PlantRegistryService) DeletePlantRegistry(id *pb.ById) error {
	if id.Id == "" {
		return errors.New("RegistryId cannot be empty")
	}

	_, err := s.Collection.DeleteOne(context.TODO(), bson.M{"RegistryId": id.Id})
	if err != nil {
		return err
	}

	return nil
}


