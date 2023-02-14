package storage

import (
	"context"
	"log"

	pb "beverage-grpc/pkg"
)

const memoryStorageTypeName = "memory"

type memory struct {
	beverageList *pb.BeverageList
}

func NewMemoryStorage() *memory {
	return &memory{
		beverageList: &pb.BeverageList{},
	}
}

func (m *memory) Add(ctx context.Context, beverage *pb.Beverage) (*pb.Beverage, error) {
	if beverage == nil {
		log.Fatalf("mem add: mil beverage provided")
	}

	m.beverageList.Beverages = append(m.beverageList.Beverages, beverage)

	return beverage, nil
}

func (m *memory) List(ctx context.Context) (*pb.BeverageList, error) {
	return m.beverageList, nil
}
