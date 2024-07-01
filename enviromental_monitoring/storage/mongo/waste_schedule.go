package mongo

import (
	"context"
	"errors"
	"time"

	pb "gitlab.com/acumen/smart_city/enviromental_monitoring/genproto/enviromental_monitoring"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type WasteCollectionService struct {
	Collection *mongo.Collection
}

func NewWasteCollectionService(db *mongo.Database) *WasteCollectionService {
	return &WasteCollectionService{
		Collection: db.Collection("waste_collection_schedules"),
	}
}

func (s *WasteCollectionService) CreateWasteCollectionSchedule(req *pb.WasteCollectionScheduleCreate) error {
	if req.ScheduleId == "" {
		req.ScheduleId = generateUUID() // Yangi registratsiya uchun UUID generatsiya qiling
	}

	schedule := &pb.WasteCollectionSchedule{
		ScheduleId:    req.ScheduleId,
		Address:       req.Address,
		CollectionDay: req.CollectionDay,
		WasteType:     req.WasteType,
		CreatedAt:     time.Now().Format(time.RFC3339),
		UpdatedAt:     time.Now().Format(time.RFC3339),
	}

	_, err := s.Collection.InsertOne(context.TODO(), schedule)
	if err != nil {
		return err
	}
	return nil
}

func (s *WasteCollectionService) UpdateWasteCollectionSchedule(req *pb.WasteCollectionScheduleCreate) error {
	if req.ScheduleId == "" {
		return errors.New("ScheduleId cannot be empty")
	}

	updateDoc := bson.M{
		"Address":       req.Address,
		"CollectionDay": req.CollectionDay,
		"WasteType":     req.WasteType,
		"UpdatedAt":     time.Now().Format(time.RFC3339),
	}

	_, err := s.Collection.UpdateOne(
		context.TODO(),
		bson.M{"ScheduleId": req.ScheduleId},
		bson.M{"$set": updateDoc},
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *WasteCollectionService) DeleteWasteCollectionSchedule(id *pb.ById) error {
	if id.Id == "" {
		return errors.New("ScheduleId cannot be empty")
	}

	_, err := s.Collection.DeleteOne(context.TODO(), bson.M{"ScheduleId": id.Id})
	if err != nil {
		return err
	}

	return nil
}

func (s *WasteCollectionService) GetWasteCollectionSchedules(filter *pb.WasteScheduleFilter) (*pb.WasteScheduleList, error) {
	bsonFilter := bson.M{}

	if filter.Address != "" {
		bsonFilter["Address"] = filter.Address
	}
	if filter.WasteType != "" {
		bsonFilter["WasteType"] = filter.WasteType
	}

	cur, err := s.Collection.Find(context.TODO(), bsonFilter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	var schedules []*pb.WasteCollectionSchedule
	for cur.Next(context.TODO()) {
		var schedule pb.WasteCollectionSchedule
		err := cur.Decode(&schedule)
		if err != nil {
			return nil, err
		}
		schedules = append(schedules, &schedule)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return &pb.WasteScheduleList{Schedules: schedules}, nil
}

