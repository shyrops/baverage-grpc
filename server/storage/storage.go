package storage

import (
	"context"
	"log"

	pb "beverage-grpc/pkg"
)

type Storage interface {
	StoreBeverage(ctx context.Context, beverage *pb.Beverage) (*pb.Beverage, error)
	GetBeverages(ctx context.Context) (*pb.BeverageList, error)
}

func NewStorage(storageType string) Storage {
	switch storageType {
	case memoryStorageTypeName:
		return NewMemoryStorage()
	case fileStorageTypeName:
		return NewFileStorage()
	default:
		log.Fatalf("incorrect storage type provided: %s", storageType)
	}

	return nil
}
