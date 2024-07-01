package mongoDB

import (
	"context"
	"fmt"

	pb "gitlab.com/acumen/smart_city/transportation/genproto/transport"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/bson"
 
)

type maintanceRepo struct {
	db *mongo.Collection
}

func NewMaintanceRepo(db *mongo.Collection) *maintanceRepo {
	return &maintanceRepo{
        db: db,
    }
}

func (r *maintanceRepo) CreateMaintranceSchedule(req *pb.MaintanceScheduleCreate) (*pb.MaintanceScheduleRes, error) {
	res := &pb.MaintanceScheduleRes{}
    fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaa")

    _, err := r.db.InsertOne(context.Background(), req)
    if err!= nil {
        return nil, err
    }
    fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaa")

    err = r.db.FindOne(context.Background(), req).Decode(&res)
    if err!= nil {
        return nil, err
    }

    return res, nil
}

func (r *maintanceRepo) GetAllMaintanceSchedules(req *pb.MaintanceScheduleFilter) (*pb.MaintanceScheduleList, error) {
    res := &pb.MaintanceScheduleList{}

    cursor, err := r.db.Find(context.Background(), req)
    if err != nil {
        return nil, err
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var schedule pb.MaintanceScheduleRes
        if err := cursor.Decode(&schedule); err != nil {
            return nil, err
        }
        res.Schedules = append(res.Schedules, &schedule)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return res, nil
}

func (r *maintanceRepo) UpdateMaintranceSchedule(req *pb.MaintanceScheduleUpdate) (*pb.MaintanceScheduleRes, error) {
    res := &pb.MaintanceScheduleRes{}

    ctx := context.Background()

    filter := bson.M{"35": req.ScheduleId}

    update := bson.M{
        "$set": bson.M{
            "1": req.EndDate, 
            "2": req.EndPoint,
        },
    }

    _, err := r.db.UpdateOne(ctx, filter, update)
    if err != nil {
        return nil, err
    }

    err = r.db.FindOne(ctx, filter).Decode(&res)
    if err != nil {
        return nil, err
    }

    return res, nil
}

func (r *maintanceRepo) DeleteMaintranceSchedule(req *pb.ById) (*pb.Void, error) {
    filter := bson.M{"_id": req.Id}

    _, err := r.db.DeleteOne(context.Background(), filter)
    if err != nil {
        return nil, err
    }

    return &pb.Void{}, nil
}

