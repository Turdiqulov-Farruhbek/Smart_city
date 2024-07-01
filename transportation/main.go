package main

import (
	"fmt"

	pb "gitlab.com/acumen/smart_city/transportation/genproto/transport"
	mongo "gitlab.com/acumen/smart_city/transportation/storage/mongoDb"
)

func main() {


    m, err := mongo.NewMongoDBStorage()
    if err!= nil {
        panic(err)
    }


	db := mongo.NewMaintanceRepo(m.Db)


	res, err := db.CreateMaintranceSchedule(
		&pb.MaintanceScheduleCreate{
            RoadId: "1",
            StartPoint: "2",
            EndPoint: "3",
            StartDate: "4",
            EndDate: "5",
            Organizer: "d",
            Status: "6",
            TotalArea: "7",
        },
	)

    if err!= nil {
        fmt.Println(err)
    }
    fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>", res)


    res, err = db.CreateMaintranceSchedule(
        &pb.MaintanceScheduleCreate{
            StartDate: "4",
            EndDate: "5",
            Organizer: "d",
            Status: "6",
            RoadId: "1",
        },
    )

    if err!= nil {
        fmt.Println(err)
    }

    fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>", res)

    res, err = db.UpdateMaintranceSchedule(
        &pb.MaintanceScheduleUpdate{
            StartDate: "4",
            EndDate: "5",
            Organizer: "d",
            Status: "6",
            RoadId: "1",
        },
    )
    if err!= nil {
        fmt.Println(err)
    }

    fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>", res)

    _, err = db.DeleteMaintranceSchedule(
        &pb.ById{
            Id: "667e5c7e8f73c0ae97aa92c2",
        },
    )
    if err!= nil {
        fmt.Println(err)
    }
    fmt.Println("res", res)
}	