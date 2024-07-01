package mongo

import (
	"context"
	"errors"
	"time"

	pb "gitlab.com/acumen/smart_city/enviromental_monitoring/genproto/enviromental_monitoring"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type RecyclingCenterService struct {
	Collection *mongo.Collection
}

func NewRecyclingCenterService(db *mongo.Database) *RecyclingCenterService {
	return &RecyclingCenterService{
		Collection: db.Collection("recycling_centers"),
	}
}

func (s *RecyclingCenterService) CreateRecyclingCenter(req *pb.RecyclingCenterCreate) error {
	if req.CenterId == "" {
		req.CenterId = generateUUID() // Yangi registratsiya uchun UUID generatsiya qiling
	}

	center := &pb.RecyclingCenter{
		CenterId:          req.CenterId,
		CenterName:        req.CenterName,
		Address:           req.Address,
		AcceptedMaterials: req.AcceptedMaterials,
		CreatedAt:         time.Now().Format(time.RFC3339),
		UpdatedAt:         time.Now().Format(time.RFC3339),
	}

	_, err := s.Collection.InsertOne(context.TODO(), center)
	if err != nil {
		return err
	}
	return nil
}

func (s *RecyclingCenterService) UpdateRecyclingCenter(req *pb.RecyclingCenterCreate) error {
	if req.CenterId == "" {
		return errors.New("CenterId cannot be empty")
	}

	updateDoc := bson.M{
		"CenterName":        req.CenterName,
		"Address":           req.Address,
		"AcceptedMaterials": req.AcceptedMaterials,
		"UpdatedAt":         time.Now().Format(time.RFC3339),
	}

	_, err := s.Collection.UpdateOne(
		context.TODO(),
		bson.M{"CenterId": req.CenterId},
		bson.M{"$set": updateDoc},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *RecyclingCenterService) DeleteRecyclingCenter(id *pb.ById) error {
	if id.Id == "" {
		return errors.New("CenterId cannot be empty")
	}

	_, err := s.Collection.DeleteOne(context.TODO(), bson.M{"CenterId": id.Id})
	if err != nil {
		return err
	}

	return nil
}

func (s *RecyclingCenterService) GetRecyclingCenters(filter *pb.RecyclingCenterFilter) (*pb.RecyclingCenterList, error) {
	bsonFilter := bson.M{}

	if filter.CenterId != "" {
		bsonFilter["CenterId"] = filter.CenterId
	}
	if filter.CenterName != "" {
		bsonFilter["CenterName"] = filter.CenterName
	}
	if filter.Address != "" {
		bsonFilter["Address"] = filter.Address
	}
	if filter.AcceptedMaterials != "" {
		bsonFilter["AcceptedMaterials"] = filter.AcceptedMaterials
	}

	cur, err := s.Collection.Find(context.TODO(), bsonFilter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var recyclingCenters []*pb.RecyclingCenter
	for cur.Next(context.TODO()) {
		var center pb.RecyclingCenter
		err := cur.Decode(&center)
		if err != nil {
			return nil, err
		}
		recyclingCenters = append(recyclingCenters, &center)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &pb.RecyclingCenterList{RecyclingCenters: recyclingCenters}, nil
}

func generateUUID() string {
	return "mock-uuid" // Bu mock-up, haqiqiy UUID generatsiyasi kerak bo'lsa uni almashtiring.
}
