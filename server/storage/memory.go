package storage

import (
	"context"

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

func (m *memory) StoreBeverage(ctx context.Context, beverage *pb.Beverage) (*pb.Beverage, error) {
	m.beverageList.Beverages = append(m.beverageList.Beverages, beverage)

	return beverage, nil
}

func (m *memory) GetBeverages(ctx context.Context) (*pb.BeverageList, error) {
	return m.beverageList, nil
}
