package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"

	pb "beverage-grpc/pkg"
)

const (
	addr = "localhost:50051"
)

type beverage struct {
	beverageType pb.BeverageType
	volume       int32
	name         string
}

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("cannot connect to sever: %v", err)
	}
	defer conn.Close()

	c := pb.NewBeveragesManagementClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for _, bev := range prepareBeverages() {
		r, err := c.CreateBeverage(
			ctx,
			&pb.CreateBeverageRequest{
				Attr: &pb.BeverageMainAttributes{
					Type:   bev.beverageType,
					Volume: bev.volume,
					Name:   bev.name},
			},
		)
		if err != nil {
			log.Fatalf("could not create bevrage: %v", err)
		}
		attr := r.GetAttr()
		log.Printf(`New beverage details:
ID: %d
Name: %s
Type: %d
Volume: %d
Price: %d

`, r.GetId(), attr.GetName(), attr.GetType(), attr.GetVolume(), r.GetPrice())
	}

}

func prepareBeverages() []beverage {
	beverages := make([]beverage, 3)

	beverages[0] = beverage{
		beverageType: pb.BeverageType_BEVERAGE_TYPE_BEER,
		volume:       300,
		name:         "Домашняя кляча",
	}
	beverages[1] = beverage{
		beverageType: pb.BeverageType_BEVERAGE_TYPE_VODKA,
		volume:       550,
		name:         "Вызыватель",
	}
	beverages[2] = beverage{
		beverageType: pb.BeverageType_BEVERAGE_TYPE_GIN,
		volume:       550,
		name:         "Прорастанец",
	}

	return beverages
}
